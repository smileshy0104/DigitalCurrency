package consumer

import (
	"common/db"
	"context"
	"encoding/json"
	"exchange/internal/database"
	"exchange/internal/domain"
	"exchange/internal/model"
	"exchange/internal/processor"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

// KafkaConsumer 消费者，处理 Kafka 中的订单消息
type KafkaConsumer struct {
	cli     *database.KafkaClient       // Kafka 客户端
	factory *processor.CoinTradeFactory // 交易引擎工厂
	db      *db.DB                      // 数据库连接
}

// NewKafkaConsumer 创建新的 KafkaConsumer 实例
func NewKafkaConsumer(cli *database.KafkaClient, factory *processor.CoinTradeFactory, db *db.DB) *KafkaConsumer {
	return &KafkaConsumer{
		cli:     cli,
		factory: factory,
		db:      db,
	}
}

// Run 启动 Kafka 消费者，处理订单交易和完成
func (k *KafkaConsumer) Run() {
	orderDomain := domain.NewOrderDomain(k.db) // 创建订单域
	k.orderTrading()                           // 启动订单交易处理
	k.orderComplete(orderDomain)               // 启动订单完成处理
}

// orderTrading 启动读取交易订单的 goroutine
func (k *KafkaConsumer) orderTrading() {
	cli := k.cli.StartReadNew("exchange_order_trading") // 启动读取交易订单的 Kafka 客户端
	go k.readOrderTrading(cli)                          // 启动 goroutine 读取订单
}

// readOrderTrading 读取交易订单并处理
func (k *KafkaConsumer) readOrderTrading(cli *database.KafkaClient) {
	for {
		kafkaData := cli.Read() // 从 Kafka 读取数据
		var order *model.ExchangeOrder
		json.Unmarshal(kafkaData.Data, &order) // 解析 JSON 数据为 ExchangeOrder 对象
		// 将订单交给撮合交易引擎进行处理
		coinTrade := k.factory.GetCoinTrade(order.Symbol) // 获取对应的交易引擎
		coinTrade.Trade(order)                            // 进行交易
	}
}

// orderComplete 启动读取完成订单的 goroutine
func (k *KafkaConsumer) orderComplete(orderDomain *domain.OrderDomain) {
	cli := k.cli.StartReadNew("exchange_order_complete") // 启动读取完成订单的 Kafka 客户端
	go k.readOrderComplete(cli, orderDomain)             // 启动 goroutine 读取完成订单
}

// readOrderComplete 读取完成订单并更新状态
func (k *KafkaConsumer) readOrderComplete(cli *database.KafkaClient, orderDomain *domain.OrderDomain) {
	for {
		kafkaData := cli.Read() // 从 Kafka 读取数据
		var order *model.ExchangeOrder
		json.Unmarshal(kafkaData.Data, &order) // 解析 JSON 数据为 ExchangeOrder 对象
		// 更新订单状态为完成
		err := orderDomain.UpdateOrderComplete(context.Background(), order)
		if err != nil {
			logx.Error(err)                    // 记录错误信息
			cli.RPut(kafkaData)                // 重新放回 Kafka
			time.Sleep(200 * time.Millisecond) // 等待后重试
			continue
		}
		// 通知钱包更新
		for {
			kafkaData.Topic = "exchange_order_complete_update_success" // 设置成功更新的主题
			err2 := cli.SendSync(kafkaData)                            // 发送同步消息
			if err2 != nil {
				logx.Error(err2)                   // 记录错误信息
				time.Sleep(250 * time.Millisecond) // 等待后重试
				continue
			}
			break // 成功发送后退出循环
		}
	}
}
