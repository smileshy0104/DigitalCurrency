package model

import (
	"common/op"
	"grpc-common/market/types/market"
)

// kline 结构体
type Kline struct {
	Period       string  `bson:"period,omitempty"`
	OpenPrice    float64 `bson:"openPrice,omitempty"`
	HighestPrice float64 `bson:"highestPrice,omitempty"`
	LowestPrice  float64 `bson:"lowestPrice,omitempty"`
	ClosePrice   float64 `bson:"closePrice,omitempty"`
	Time         int64   `bson:"time,omitempty"`
	Count        float64 `bson:"count,omitempty"`    //成交笔数
	Volume       float64 `bson:"volume,omitempty"`   //成交量
	Turnover     float64 `bson:"turnover,omitempty"` //成交额
}

// Table 交易行情表表名
func (*Kline) Table(symbol, period string) string {
	return "exchange_kline_" + symbol + "_" + period
}

// ToCoinThumb 将K线数据转换为CoinThumb对象。
// 这个函数用于从K线数据中提取并计算所需字段，以填充CoinThumb结构体。
// 参数:
//
//	symbol - 交易对符号，例如"BTCUSDT"。
//	end - 指向用作参考的结束K线数据，通常用于计算变化量和百分比。
//
// 返回值:
//
//	*market.CoinThumb - 返回一个指向CoinThumb对象的指针，包含了计算出的字段。
func (k *Kline) ToCoinThumb(symbol string, end *Kline) *market.CoinThumb {
	// 初始化CoinThumb对象。
	ct := &market.CoinThumb{}

	// 设置交易对符号。
	ct.Symbol = symbol

	// 设置收盘价和开盘价。
	ct.Close = k.ClosePrice
	ct.Open = k.OpenPrice

	// 设置时区，默认为0。
	ct.Zone = 0

	// 计算并设置变化量（当前K线的收盘价与参考K线的收盘价之差）。
	ct.Change = k.ClosePrice - end.ClosePrice

	// 计算并设置变化百分比，保留5位小数。
	ct.Chg = op.MulN(op.DivN(ct.Change, end.ClosePrice, 5), 100, 5)

	// 设置美元汇率和基础美元汇率。
	ct.UsdRate = k.ClosePrice
	ct.BaseUsdRate = 1

	// 设置日期时间。
	ct.DateTime = k.Time

	// 返回填充完毕的CoinThumb对象。
	return ct
}

// DefaultCoinThumb 默认值空值
func DefaultCoinThumb(symbol string) *market.CoinThumb {
	// 创建一个默认的CoinThumb对象
	ct := &market.CoinThumb{}
	// 设置符号
	ct.Symbol = symbol
	// 设置为空float64切片
	ct.Trend = []float64{}
	return ct
}
