package model

import (
	"common/op"
	"grpc-common/market/types/market"
)

// Kline K线结构体
type Kline struct {
	Period       string  `bson:"period,omitempty" json:"period"`             // K线周期
	OpenPrice    float64 `bson:"openPrice,omitempty" json:"openPrice"`       // 开盘价
	HighestPrice float64 `bson:"highestPrice,omitempty" json:"highestPrice"` // 最高价
	LowestPrice  float64 `bson:"lowestPrice,omitempty" json:"lowestPrice"`   // 最低价
	ClosePrice   float64 `bson:"closePrice,omitempty" json:"closePrice"`     // 收盘价
	Time         int64   `bson:"time,omitempty" json:"time"`                 // 时间戳
	Count        float64 `bson:"count,omitempty" json:"count"`               // 成交笔数
	Volume       float64 `bson:"volume,omitempty" json:"volume"`             // 成交量
	Turnover     float64 `bson:"turnover,omitempty" json:"turnover"`         // 成交额
}

// ToCoinThumb 将K线数据转换成CoinThumb对象
// symbol: 币种符号
// ct: 原始的CoinThumb对象
// 返回转换后的CoinThumb对象
func (k *Kline) ToCoinThumb(symbol string, ct *market.CoinThumb) *market.CoinThumb {
	isSame := false
	if ct.Symbol == symbol && ct.DateTime == k.Time {
		// 认为这是同一个数据
		isSame = true
	}
	if !isSame {
		// 更新CoinThumb对象的各个字段
		ct.Close = k.ClosePrice
		ct.Open = k.OpenPrice
		if ct.High < k.HighestPrice {
			ct.High = k.HighestPrice
		}
		if ct.Low > k.LowestPrice {
			ct.Low = k.LowestPrice
		}
		ct.Zone = 0
		// 使用精简加法运算
		ct.Volume = op.AddN(k.Volume, ct.Volume, 8)
		ct.Turnover = op.AddN(k.Turnover, ct.Turnover, 8)
		// 计算涨跌额和涨跌幅
		ct.Change = k.LowestPrice - ct.Close
		ct.Chg = op.MulN(op.DivN(ct.Change, ct.Close, 5), 100, 3)
		// 设置USD汇率和基础USD汇率
		ct.UsdRate = k.ClosePrice
		ct.BaseUsdRate = 1
		// 更新趋势列表
		ct.Trend = append(ct.Trend, k.ClosePrice)
		// 更新时间戳
		ct.DateTime = k.Time
	}
	return ct
}

// InitCoinThumb 初始化CoinThumb对象
// symbol: 币种符号
// 返回初始化后的CoinThumb对象
func (k *Kline) InitCoinThumb(symbol string) *market.CoinThumb {
	ct := &market.CoinThumb{}
	// 初始化CoinThumb对象的各个字段
	ct.Symbol = symbol
	ct.Close = k.ClosePrice
	ct.Open = k.OpenPrice
	ct.High = k.HighestPrice
	ct.Volume = k.Volume
	ct.Turnover = k.Turnover
	ct.Low = k.LowestPrice
	ct.Zone = 0
	ct.Change = 0
	ct.Chg = 0.0
	ct.UsdRate = k.ClosePrice
	ct.BaseUsdRate = 1
	ct.Trend = make([]float64, 0)
	ct.DateTime = k.Time
	return ct
}

// CoinThumb 币种摘要信息结构体
type CoinThumb struct {
	Symbol       string    `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol"`                // 币种符号
	Open         float64   `protobuf:"fixed64,2,opt,name=open,proto3" json:"open"`                  // 开盘价
	High         float64   `protobuf:"fixed64,3,opt,name=high,proto3" json:"high"`                  // 最高价
	Low          float64   `protobuf:"fixed64,4,opt,name=low,proto3" json:"low"`                    // 最低价
	Close        float64   `protobuf:"fixed64,5,opt,name=close,proto3" json:"close"`                // 收盘价
	Chg          float64   `protobuf:"fixed64,6,opt,name=chg,proto3" json:"chg"`                    // 涨跌幅
	Change       float64   `protobuf:"fixed64,7,opt,name=change,proto3" json:"change"`              // 涨跌额
	Volume       float64   `protobuf:"fixed64,8,opt,name=volume,proto3" json:"volume"`              // 成交量
	Turnover     float64   `protobuf:"fixed64,9,opt,name=turnover,proto3" json:"turnover"`          // 成交额
	LastDayClose float64   `protobuf:"fixed64,10,opt,name=lastDayClose,proto3" json:"lastDayClose"` // 前收盘价
	UsdRate      float64   `protobuf:"fixed64,11,opt,name=usdRate,proto3" json:"usdRate"`           // USD汇率
	BaseUsdRate  float64   `protobuf:"fixed64,12,opt,name=baseUsdRate,proto3" json:"baseUsdRate"`   // 基础USD汇率
	Zone         float64   `protobuf:"fixed64,13,opt,name=zone,proto3" json:"zone"`                 // 时区
	DateTime     int64     `protobuf:"varint,14,opt,name=dateTime,proto3" json:"dateTime"`          // 时间戳
	Trend        []float64 `protobuf:"fixed64,15,rep,packed,name=trend,proto3" json:"trend"`        // 趋势列表
}
