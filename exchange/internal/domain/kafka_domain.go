package domain

import (
	"context"
	"encoding/json"
	"exchange/internal/database"
	"exchange/internal/model"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

// KafkaDomain 实现kafka消息队列
type KafkaDomain struct {
	cli         *database.KafkaClient
	orderDomain *OrderDomain
}

// SendOrderAdd 发送订单添加消息到Kafka
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
//	error: 错误信息，如果发送消息成功则为nil
func (d *KafkaDomain) SendOrderAdd(topic string, userId int64, orderId string, money float64, symbol string,
	direction int, baseSymbol string, coinSymbol string) error {
	m := make(map[string]any)
	m["userId"] = userId
	m["orderId"] = orderId
	m["money"] = money
	m["symbol"] = symbol
	m["direction"] = direction
	m["baseSymbol"] = baseSymbol
	m["coinSymbol"] = coinSymbol
	marshal, _ := json.Marshal(m)
	// 使用Kafka客户端发送消息
	data := database.KafkaData{
		Topic: topic,
		Key:   []byte(orderId),
		Data:  marshal,
	}
	// 同步发送消息
	err := d.cli.SendSync(data)
	logx.Info("创建订单，发消息成功,orderId=" + orderId)
	return err
}

// OrderResult 订单结果结构体
type OrderResult struct {
	UserId  int64  `json:"userId"`
	OrderId string `json:"orderId"`
}

// WaitAddOrderResult 监听并处理订单初始化完成的消息
func (d *KafkaDomain) WaitAddOrderResult() {
	cli := d.cli.StartReadNew("exchange_order_init_complete_trading")
	for {
		kafkaData := cli.Read()
		logx.Info("读取exchange_order_init_complete_trading 消息成功:" + string(kafkaData.Key))
		var orderResult OrderResult
		json.Unmarshal(kafkaData.Data, &orderResult)
		exchangeOrder, err := d.orderDomain.FindOrderByOrderId(context.Background(), orderResult.OrderId)
		if err != nil {
			logx.Error(err)
			err := d.orderDomain.UpdateStatusCancel(context.Background(), orderResult.OrderId)
			if err != nil {
				d.cli.RPut(kafkaData)
			}
			continue
		}
		if exchangeOrder == nil {
			logx.Error("订单id不存在")
			continue
		}
		if exchangeOrder.Status != model.Init {
			logx.Error("订单已经被处理过")
			continue
		}
		err = d.orderDomain.UpdateOrderStatusTrading(context.Background(), orderResult.OrderId)
		if err != nil {
			logx.Error(err)
			d.cli.RPut(kafkaData)
			continue
		}
		//需要发送消息到kafka 订单需要加入到撮合交易当中
		//如果没有撮合交易成功 加入撮合交易的队列 继续等待完成撮合
		exchangeOrder.Status = model.Trading
		for {
			bytes, _ := json.Marshal(exchangeOrder)
			orderData := database.KafkaData{
				Topic: "exchange_order_trading",
				Key:   []byte(exchangeOrder.OrderId),
				Data:  bytes,
			}
			err := d.cli.SendSync(orderData)
			if err != nil {
				logx.Error(err)
				time.Sleep(250 * time.Millisecond)
				continue
			}
			break
		}
	}
}

// NewKafkaDomain 创建新的KafkaDomain实例
func NewKafkaDomain(cli *database.KafkaClient, orderDomain *OrderDomain) *KafkaDomain {
	// 创建一个新的KafkaDomain实例
	k := &KafkaDomain{
		cli:         cli,
		orderDomain: orderDomain,
	}
	// 启动一个协程，监听并处理订单初始化完成消息
	go k.WaitAddOrderResult()
	return k
}
