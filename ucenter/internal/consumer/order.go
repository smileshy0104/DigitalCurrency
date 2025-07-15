package consumer

import (
	mainDb "common/db"
	"common/db/tran"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"grpc-common/exchange/ec_client"
	"grpc-common/exchange/types/order"
	"time"
	"ucenter/internal/database"
	"ucenter/internal/domain"
)

// OrderAdd 结构体用于表示订单信息。
type OrderAdd struct {
	UserId     int64   `json:"userId"`     // 用户 ID
	OrderId    string  `json:"orderId"`    // 订单 ID
	Money      float64 `json:"money"`      // 订单金额
	Symbol     string  `json:"symbol"`     // 交易对符号
	Direction  int     `json:"direction"`  // 订单方向（买入/卖出）
	BaseSymbol string  `json:"baseSymbol"` // 基础货币符号
	CoinSymbol string  `json:"coinSymbol"` // 交易货币符号
}

// ExchangeOrderAdd 从 Kafka 中获取创建订单的消息并处理
func ExchangeOrderAdd(redisCli *redis.Redis, cli *database.KafkaClient, orderRpc ec_client.Order, db *mainDb.DB) {
	// 循环读取 Kafka 消息
	for {
		kafkaData := cli.Read()                     // 从 Kafka 客户端读取消息
		orderId := string(kafkaData.Key)            // 获取订单 ID
		logx.Info("接收到创建订单的消息, orderId=" + orderId) // 记录接收到的消息

		var addData OrderAdd
		// 反序列化 Kafka 消息数据到 OrderAdd 结构体
		json.Unmarshal(kafkaData.Data, &addData)

		// 验证消息中的订单 ID 是否与 Kafka 消息的 Key 匹配
		if orderId != addData.OrderId {
			logx.Error("消息数据有误") // 记录错误信息
			continue             // 继续下一次循环
		}

		ctx := context.Background() // 创建上下文

		// 调用 RPC 方法冻结钱
		exchangeOrderOrigin, err := orderRpc.FindByOrderId(ctx, &order.OrderReq{
			OrderId: orderId, // 通过订单 ID 查找订单
		})
		if err != nil {
			logx.Error(err)                                     // 记录错误信息
			cancelOrder(ctx, kafkaData, orderId, orderRpc, cli) // 取消订单
			continue                                            // 继续下一次循环
		}

		if exchangeOrderOrigin == nil {
			logx.Error("orderId :" + orderId + " 不存在") // 记录错误信息
			continue                                   // 继续下一次循环
		}

		// 检查订单状态是否为 4（初始化状态）
		if exchangeOrderOrigin.Status != 4 {
			logx.Error("orderId :" + orderId + " 已经被操作过了") // 记录错误信息
			continue                                       // 继续下一次循环
		}

		// 创建 Redis 锁，锁的名称由用户 ID 和订单 ID 组成，以确保唯一性
		lock := redis.NewRedisLock(redisCli, "exchange_order::"+fmt.Sprintf("%d::%s", addData.UserId, orderId))

		// 尝试获取锁
		acquire, err := lock.Acquire()
		if err != nil {
			logx.Error(err)              // 记录获取锁时的错误
			logx.Info("已经有别的进程在处理了....") // 记录锁已被占用的信息
			continue                     // 跳过本次循环，继续下一个订单处理
		}

		// 如果成功获取到锁
		if acquire {
			transaction := tran.NewTransaction(db.Conn)                // 创建新的数据库事务
			walletDomain := domain.NewMemberWalletDomain(db, nil, nil) // 创建钱包域实例

			// 执行事务操作
			err := transaction.Action(func(conn mainDb.DbConn) error {
				if addData.Direction == 0 {
					// 如果是买入方向
					err := walletDomain.Freeze(ctx, conn, addData.UserId, addData.Money, addData.BaseSymbol) // 冻结基础货币
					return err
				} else {
					// 如果是卖出方向
					err := walletDomain.Freeze(ctx, conn, addData.UserId, addData.Money, addData.CoinSymbol) // 冻结交易货币
					return err
				}
			})

			// 如果事务执行出错，取消订单并继续下一个处理
			if err != nil {
				cancelOrder(ctx, kafkaData, orderId, orderRpc, cli)
				continue
			}

			// 需要将状态改为 trading
			// 完成后通知订单进行状态变更，确保消息发送成功
			for {
				m := make(map[string]any)     // 创建消息体
				m["userId"] = addData.UserId  // 设置用户 ID
				m["orderId"] = orderId        // 设置订单 ID
				marshal, _ := json.Marshal(m) // 将消息体序列化为 JSON

				// 创建 Kafka 数据结构
				data := database.KafkaData{
					Topic: "exchange_order_init_complete_trading", // 主题
					Key:   []byte(orderId),                        // 键
					Data:  marshal,                                // 序列化后的数据
				}

				// 尝试同步发送消息
				err := cli.SendSync(data)
				if err != nil {
					logx.Error(err)                    // 记录发送消息时的错误
					time.Sleep(250 * time.Millisecond) // 等待 250 毫秒后重试
					continue                           // 继续尝试发送
				}

				logx.Info("发送exchange_order_init_complete_trading 消息成功:" + orderId) // 记录成功发送消息的信息
				break                                                               // 发送成功后跳出循环
			}

			lock.Release() // 释放锁
		}

	}
}

// cancelOrder 取消订单并将消息放回 Kafka
func cancelOrder(ctx context.Context, data database.KafkaData, orderId string, orderRpc ec_client.Order, cli *database.KafkaClient) {
	// 调用 RPC 方法取消订单
	_, err := orderRpc.CancelOrder(ctx, &order.OrderReq{
		OrderId: orderId, // 通过订单 ID 取消订单
	})
	if err != nil {
		cli.RPut(data) // 如果取消失败，将消息放回 Kafka
	}
}
