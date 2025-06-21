package model

import (
	"common/tools"
)

// Kline k线 (使用bson)
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
	TimeStr      string  `bson:"timeStr,omitempty"`
}

// Table 生成k线数据表名
// 参数:
//
//	symbol - 交易对符号，如BTC/USDT
//	period - k线周期，如1m, 5m
//
// 返回值:
//
//	表名字符串
func (*Kline) Table(symbol, period string) string {
	// 返回自定义表名字符串
	return "exchange_kline_" + symbol + "_" + period
}

// NewKline 创建新的Kline实例
// 参数:
//
//	data - 包含k线数据的字符串数组
//	period - k线周期
//
// 返回值:
//
//	*Kline - 指向新创建的Kline实例的指针
func NewKline(data []string, period string) *Kline {
	// 将字符串数组的第一个元素转换为int64
	toInt64 := tools.ToInt64(data[0])
	// 创建新的Kline实例
	return &Kline{
		Time:         toInt64,
		Period:       period,
		OpenPrice:    tools.ToFloat64(data[1]),
		HighestPrice: tools.ToFloat64(data[2]),
		LowestPrice:  tools.ToFloat64(data[3]),
		ClosePrice:   tools.ToFloat64(data[4]),
		Count:        tools.ToFloat64(data[5]),
		Volume:       tools.ToFloat64(data[6]),
		Turnover:     tools.ToFloat64(data[7]),
		TimeStr:      tools.ToTimeString(toInt64),
	}
}

// OkxKlineRes OKX交易所k线响应数据结构
type OkxKlineRes struct {
	Code string     `json:"code"`
	Msg  string     `json:"msg"`
	Data [][]string `json:"data"`
}
