package domain

import (
	"context"
	"encoding/json"
	"exchange/internal/database"
	"exchange/internal/model"
	"github.com/zeromicro/go-zero/core/logx"
)

// KafkaDomain 实现 Kafka 消息队列
type KafkaDomain struct {
	cli         *database.KafkaClient // Kafka 客户端
	orderDomain *OrderDomain          // 订单域，用于与订单相关的操作
}

// SendOrderAdd 发送订单添加消息到 Kafka
// 参数:
//
//	topic: 消息主题
//	userId: 用户ID
//	orderId: 订单ID
//	money: 金额
//	symbol: 交易对符号
//	direction: 交易方向
//	baseSymbol: 基础货币符号
//	coinSymbol: 加密货币符号
//
// 返回值:
//
//	error: 错误信息，如果发送消息成功则为 nil
func (d *KafkaDomain) SendOrderAdd(topic string, userId int64, orderId string, money float64, symbol string,
	direction int, baseSymbol string, coinSymbol string) error {
	// 创建一个包含订单信息的 map
	m := make(map[string]any)
	m["userId"] = userId
	m["orderId"] = orderId
	m["money"] = money
	m["symbol"] = symbol
	m["direction"] = direction
	m["baseSymbol"] = baseSymbol
	m["coinSymbol"] = coinSymbol

	// 将 map 转换为 JSON 格式
	marshal, _ := json.Marshal(m)

	// 使用 Kafka 客户端发送消息
	data := database.KafkaData{
		Topic: topic,
		Key:   []byte(orderId), // 使用订单ID作为消息键
		Data:  marshal,         // 消息内容
	}

	// 同步发送消息
	err := d.cli.SendSync(data)
	logx.Info("创建订单，发消息成功, orderId=" + orderId)
	return err // 返回发送结果的错误信息
}

// OrderResult 订单结果结构体
type OrderResult struct {
	UserId  int64  `json:"userId"`  // 用户ID
	OrderId string `json:"orderId"` // 订单ID
}

// WaitAddOrderResult 监听并处理订单初始化完成的消息
func (d *KafkaDomain) WaitAddOrderResult() {
	// 启动 Kafka 客户端，开始读取消息
	cli := d.cli.StartReadNew("exchange_order_init_complete_trading")
	for {
		kafkaData := cli.Read() // 读取消息
		logx.Info("读取 exchange_order_init_complete_trading 消息成功:" + string(kafkaData.Key))

		// 解析消息数据
		var orderResult OrderResult
		json.Unmarshal(kafkaData.Data, &orderResult)

		// 根据订单ID查找订单
		exchangeOrder, err := d.orderDomain.FindOrderByOrderId(context.Background(), orderResult.OrderId)
		if err != nil {
			logx.Error(err)                                                                    // 记录错误信息
			err := d.orderDomain.UpdateStatusCancel(context.Background(), orderResult.OrderId) // 更新订单状态为取消
			if err != nil {
				d.cli.RPut(kafkaData) // 重新放回消息队列
			}
			continue
		}

		if exchangeOrder == nil {
			logx.Error("订单id不存在") // 订单不存在的情况
			continue
		}

		if exchangeOrder.Status != model.Init {
			logx.Error("订单已经被处理过") // 订单已被处理的情况
			continue
		}

		// 更新订单状态为交易中
		err = d.orderDomain.UpdateOrderStatusTrading(context.Background(), orderResult.OrderId)
		if err != nil {
			logx.Error(err)       // 记录错误信息
			d.cli.RPut(kafkaData) // 重新放回消息队列
			continue
		}

		// 需要发送消息到 Kafka，订单需要加入到撮合交易当中
		//exchangeOrder.Status = model.Trading
		//for {
		//	bytes, _ := json.Marshal(exchangeOrder) // 将订单信息转换为 JSON 格式
		//	orderData := database.KafkaData{
		//		Topic: "exchange_order_trading",      // 目标主题
		//		Key:   []byte(exchangeOrder.OrderId), // 使用订单ID作为消息键
		//		Data:  bytes,                         // 消息内容
		//	}
		//	err := d.cli.SendSync(orderData) // 同步发送消息
		//	if err != nil {
		//		logx.Error(err)                    // 记录错误信息
		//		time.Sleep(250 * time.Millisecond) // 等待后重试
		//		continue
		//	}
		//	break // 发送成功，退出循环
		//}
	}
}

// NewKafkaDomain 创建新的 KafkaDomain 实例
func NewKafkaDomain(cli *database.KafkaClient, orderDomain *OrderDomain) *KafkaDomain {
	// 创建一个新的 KafkaDomain 实例
	k := &KafkaDomain{
		cli:         cli,
		orderDomain: orderDomain,
	}
	// 启动一个协程，监听并处理订单初始化完成消息
	go k.WaitAddOrderResult()
	return k // 返回实例
}
