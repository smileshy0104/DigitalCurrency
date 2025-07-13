package model

import (
	"github.com/jinzhu/copier"
)

// ExchangeOrder 交易订单
type ExchangeOrder struct {
	Id            int64   `gorm:"column:id"`             // 订单ID
	OrderId       string  `gorm:"column:order_id"`       // 订单编号
	Amount        float64 `gorm:"column:amount"`         // 订单数量
	BaseSymbol    string  `gorm:"column:base_symbol"`    // 基础币种符号
	CanceledTime  int64   `gorm:"column:canceled_time"`  // 取消时间戳
	CoinSymbol    string  `gorm:"column:coin_symbol"`    // 数字币符号
	CompletedTime int64   `gorm:"column:completed_time"` // 完成时间戳
	Direction     int     `gorm:"column:direction"`      // 交易方向，0-买入，1-卖出
	MemberId      int64   `gorm:"column:member_id"`      // 会员ID
	Price         float64 `gorm:"column:price"`          // 订单价格
	Status        int     `gorm:"column:status"`         // 订单状态
	Symbol        string  `gorm:"column:symbol"`         // 交易对符号
	Time          int64   `gorm:"column:time"`           // 订单创建时间戳
	TradedAmount  float64 `gorm:"column:traded_amount"`  // 成交数量
	Turnover      float64 `gorm:"column:turnover"`       // 成交额
	Type          int     `gorm:"column:type"`           // 订单类型，0-市价单，1-限价单
	UseDiscount   string  `gorm:"column:use_discount"`   // 是否使用优惠
}

// TableName 返回ExchangeOrder对应的数据库表名
func (*ExchangeOrder) TableName() string {
	return "exchange_order"
}

// status 订单状态常量
const (
	Trading = iota // 交易中
	Completed
	Canceled
	OverTimed
)

// statusMap 订单状态代码到文本的映射
var statusMap = map[int]string{
	Trading:   "TRADING",
	Completed: "COMPLETED",
	Canceled:  "CANCELED",
	OverTimed: "OVERTIMED",
}

// direction 交易方向常量
const (
	BUY = iota // 买入
	SELL
)

// directionMap 交易方向代码到文本的映射
var directionMap = map[int]string{
	BUY:  "BUY",
	SELL: "SELL",
}

// type 订单类型常量
const (
	MarketPrice = iota // 市价单
	LimitPrice
)

// typeMap 订单类型代码到文本的映射
var typeMap = map[int]string{
	MarketPrice: "MARKET_PRICE",
	LimitPrice:  "LIMIT_PRICE",
}

// ExchangeOrderVo 交易订单的视图对象
type ExchangeOrderVo struct {
	OrderId       string  `gorm:"column:order_id"`       // 订单编号
	Amount        float64 `gorm:"column:amount"`         // 订单数量
	BaseSymbol    string  `gorm:"column:base_symbol"`    // 基础币种符号
	CanceledTime  int64   `gorm:"column:canceled_time"`  // 取消时间戳
	CoinSymbol    string  `gorm:"column:coin_symbol"`    // 数字币符号
	CompletedTime int64   `gorm:"column:completed_time"` // 完成时间戳
	Direction     string  `gorm:"column:direction"`      // 交易方向文本
	MemberId      int64   `gorm:"column:member_id"`      // 会员ID
	Price         float64 `gorm:"column:price"`          // 订单价格
	Status        string  `gorm:"column:status"`         // 订单状态文本
	Symbol        string  `gorm:"column:symbol"`         // 交易对符号
	Time          int64   `gorm:"column:time"`           // 订单创建时间戳
	TradedAmount  float64 `gorm:"column:traded_amount"`  // 成交数量
	Turnover      float64 `gorm:"column:turnover"`       // 成交额
	Type          string  `gorm:"column:type"`           // 订单类型文本
	UseDiscount   string  `gorm:"column:use_discount"`   // 是否使用优惠
}

// ToVo 将ExchangeOrder转换为ExchangeOrderVo
func (old *ExchangeOrder) ToVo() *ExchangeOrderVo {
	eo := &ExchangeOrderVo{}
	copier.Copy(eo, old)
	eo.Status = statusMap[old.Status]          // 将状态代码转换为文本
	eo.Direction = directionMap[old.Direction] // 将方向代码转换为文本
	eo.Type = typeMap[old.Type]                // 将类型代码转换为文本
	return eo
}
