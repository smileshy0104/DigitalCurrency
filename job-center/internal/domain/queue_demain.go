package domain

import (
	"encoding/json"
	"job-center/internal/database"
	"job-center/internal/model"
	"log"
)

// kafka topic 主题
const KLINE1M = "kline_1m"
const BtcTransactionTopic = "BTC_TRANSACTION"

// QueueDomain 队列模块
type QueueDomain struct {
	kafkaCli *database.KafkaClient
}

// Send1mKline 发送1分钟K线数据到Kafka
// 参数:
//
//	data: K线数据数组
//	symbol: 交易对符号
func (d *QueueDomain) Send1mKline(data []string, symbol string) {
	// 创建Kline实例
	kline := model.NewKline(data, "1m")
	bytes, _ := json.Marshal(kline)
	// 将Kline数据转换为JSON格式
	msg := database.KafkaData{
		Topic: KLINE1M,
		Data:  bytes,
		Key:   []byte(symbol),
	}
	// 发送Kline数据到Kafka
	d.kafkaCli.Send(msg)
	log.Println("=================发送数据成功==============")
}

// SendRecharge 发送充值信息到Kafka
// 参数:
//
//	value: 充值金额
//	address: 充值地址
//	time: 充值时间戳
func (d *QueueDomain) SendRecharge(value float64, address string, time int64) {
	data := make(map[string]any)
	data["value"] = value
	data["address"] = address
	data["time"] = time
	data["type"] = model.RECHARGE
	data["symbol"] = "BTC"
	marshal, _ := json.Marshal(data)
	msg := database.KafkaData{
		Topic: BtcTransactionTopic,
		Data:  marshal,
		Key:   []byte(address),
	}
	d.kafkaCli.Send(msg)
}

// NewQueueDomain 创建新的QueueDomain实例
func NewQueueDomain(kafkaCli *database.KafkaClient) *QueueDomain {
	// 创建一个新的QueueDomain实例
	return &QueueDomain{
		kafkaCli: kafkaCli,
	}
}
