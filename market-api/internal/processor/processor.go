package processor

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/market/mk_client"
	"grpc-common/market/types/market"
	"market-api/internal/database"
	"market-api/internal/model"
)

// kafka中key名称
const KLINE1M = "kline_1m"
const KLINE = "kline"
const TRADE = "trade"
const TradePlateTopic = "exchange_order_trade_plate"
const TradePlate = "tradePlate"

// MarketHandler定义了处理市场数据的接口
type MarketHandler interface {
	HandleTrade(symbol string, data []byte)
	HandleKLine(symbol string, kline *model.Kline, thumbMap map[string]*market.CoinThumb)
	HandleTradePlate(symbol string, tp *model.TradePlateResult)
}

// ProcessData表示从kafka中读取的数据结构
type ProcessData struct {
	Type string //trade 交易 kline k线
	Key  []byte
	Data []byte
}

// Processor定义了处理市场数据的处理器接口
type Processor interface {
	GetThumb() any
	Process(data ProcessData)
	AddHandler(h MarketHandler)
}

// DefaultProcessor是Processor接口的默认实现
type DefaultProcessor struct {
	kafkaCli *database.KafkaClient
	handlers []MarketHandler
	thumbMap map[string]*market.CoinThumb
}

// NewDefaultProcessor 创建一个新的 DefaultProcessor 实例。
// 该函数接收一个 *database.KafkaClient 参数 kafkaCli，用于处理 Kafka 相关的操作。
// 返回值是 *DefaultProcessor，即初始化后的 DefaultProcessor 实例。
// 此函数主要用于设置 DefaultProcessor 的初始状态，包括初始化 Kafka 客户端、市场处理器切片和币种缩略图映射。
func NewDefaultProcessor(kafkaCli *database.KafkaClient) *DefaultProcessor {
	return &DefaultProcessor{
		kafkaCli: kafkaCli,                           // 初始化 Kafka 客户端
		handlers: make([]MarketHandler, 0),           // 创建一个空的 MarketHandler 切片
		thumbMap: make(map[string]*market.CoinThumb), // 创建一个空的币种缩略图映射
	}
}

// Process 是 DefaultProcessor 的一个方法，用于处理不同类型的数据。
// 它根据数据的类型，解析数据并调用相应的处理函数。
// 参数:
//
//	data (ProcessData): 包含待处理数据的对象，包括数据类型、键、和数据内容。
func (d *DefaultProcessor) Process(data ProcessData) {
	// 判断数据类型
	if data.Type == KLINE { // 如果是K线数据
		// 从Key中获取交易对符号
		symbol := string(data.Key)
		// 初始化K线数据对象
		kline := &model.Kline{}
		// 解析数据到K线数据对象————kline
		json.Unmarshal(data.Data, kline)
		// 遍历所有处理器并调用它们的HandleKLine方法处理K线数据
		for _, v := range d.handlers {
			v.HandleKLine(symbol, kline, d.thumbMap)
		}
	} else if data.Type == TradePlate { // 如果是盘口数据
		// 从Key中获取交易对符号
		symbol := string(data.Key)
		// 初始化盘口数据对象
		tp := &model.TradePlateResult{}
		// 解析数据到盘口数据对象
		json.Unmarshal(data.Data, tp)
		// 遍历所有处理器并调用它们的HandleTradePlate方法处理盘口数据
		for _, v := range d.handlers {
			v.HandleTradePlate(symbol, tp)
		}
	}
}

// AddHandler添加一个MarketHandler处理市场数据
func (d *DefaultProcessor) AddHandler(h MarketHandler) {
	//发送到websocket的服务，后面统一处理
	d.handlers = append(d.handlers, h)
}

// Init初始化DefaultProcessor
func (p *DefaultProcessor) Init(marketRpc mk_client.Market) {
	// 启动从Kafka中读取K线数据的goroutine
	p.startReadFromKafka(KLINE1M, KLINE)
	// 启动从Kafka中读取盘口数据的goroutine
	p.startReadTradePlate(TradePlateTopic)
	// 初始化缩略图信息————初始化d.thumbMap信息
	p.initThumbMap(marketRpc)
}

// GetThumb获取所有币种的缩略图信息
func (d *DefaultProcessor) GetThumb() any {
	// 创建一个切片，用于存储缩略图信息
	cs := make([]*market.CoinThumb, len(d.thumbMap))
	i := 0
	// 遍历缩略图信息，将信息复制到切片中
	for _, v := range d.thumbMap {
		cs[i] = v
		i++
	}
	return cs
}

// startReadFromKafka 开始从Kafka主题读取数据并处理
// 该函数首先启动对指定Kafka主题的读取，然后在单独的协程中处理读取到的数据
// 参数:
//
//	topic - 需要读取的Kafka主题名称
//	tp - 处理数据的类型或策略
func (p *DefaultProcessor) startReadFromKafka(topic string, tp string) {
	// 一定要先start 后read
	p.kafkaCli.StartRead(topic) // 启动读取kafka数据
	go p.dealQueueData(p.kafkaCli, tp)
}

// dealQueueData处理从kafka中读取的数据
func (p *DefaultProcessor) dealQueueData(cli *database.KafkaClient, tp string) {
	//这就是队列的数据
	for {
		// Read 从读通道k.readChan读取消息
		msg := cli.Read()
		// 创建一个ProcessData对象
		data := ProcessData{
			Type: tp,       // topic
			Key:  msg.Key,  // key
			Data: msg.Data, // value
		}
		// 处理数据
		p.Process(data)
	}
}

// initThumbMap 初始化币种缩略图映射表。
// 该方法通过调用 Market 服务的 FindSymbolThumbTrend 方法获取币种缩略图信息，并将其存储在 thumbMap 中，
// 以便快速查找和访问。这有助于提高处理效率，减少重复的网络请求。
func (d *DefaultProcessor) initThumbMap(marketRpc mk_client.Market) {
	// 调用 Market 服务的 FindSymbolThumbTrend 方法获取币种缩略图信息。
	symbolThumbRes, err := marketRpc.FindSymbolThumbTrend(context.Background(),
		&market.MarketReq{})
	if err != nil {
		// 如果发生错误，记录错误信息。
		logx.Info(err)
	} else {
		// 获取币种缩略图列表。
		coinThumbs := symbolThumbRes.List
		// 遍历币种缩略图列表，将币种符号和对应的缩略图信息存储在 thumbMap 中。
		for _, v := range coinThumbs {
			d.thumbMap[v.Symbol] = v
		}
	}
}

// startReadTradePlate开始读取交易盘数据
func (p *DefaultProcessor) startReadTradePlate(topic string) {
	// 创建一个读取交易盘数据的 goroutine
	cli := p.kafkaCli.StartReadNew(topic)
	// 处理队列数据
	go p.dealQueueData(cli, TradePlate)
}
