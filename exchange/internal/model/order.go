package model

import (
	"common/enum"
	"github.com/jinzhu/copier"
)

// ExchangeOrder 交易订单
type ExchangeOrder struct {
	Id            int64   `gorm:"column:id" json:"id"`
	OrderId       string  `gorm:"column:order_id" json:"orderId"`
	Amount        float64 `gorm:"column:amount" json:"amount"`
	BaseSymbol    string  `gorm:"column:base_symbol" json:"baseSymbol"`
	CanceledTime  int64   `gorm:"column:canceled_time" json:"canceledTime"`
	CoinSymbol    string  `gorm:"column:coin_symbol" json:"coinSymbol"`
	CompletedTime int64   `gorm:"column:completed_time" json:"completedTime"`
	Direction     int     `gorm:"column:direction" json:"direction"`
	MemberId      int64   `gorm:"column:member_id" json:"memberId"`
	Price         float64 `gorm:"column:price" json:"price"`
	Status        int     `gorm:"column:status" json:"status"`
	Symbol        string  `gorm:"column:symbol" json:"symbol"`
	Time          int64   `gorm:"column:time" json:"time"`
	TradedAmount  float64 `gorm:"column:traded_amount" json:"tradedAmount"`
	Turnover      float64 `gorm:"column:turnover" json:"turnover"`
	Type          int     `gorm:"column:type" json:"type"`
	UseDiscount   string  `gorm:"column:use_discount" json:"useDiscount"`
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
	Init
)

// StatusMap 订单状态代码到文本的映射
var StatusMap = enum.Enum{
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

// DirectionMap 交易方向代码到文本的映射
var DirectionMap = enum.Enum{
	BUY:  "BUY",
	SELL: "SELL",
}

// type 订单类型常量
const (
	MarketPrice = iota // 市价单
	LimitPrice
)

// TypeMap 订单类型代码到文本的映射
var TypeMap = enum.Enum{
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
	eo.Status = StatusMap.Value(old.Status)    // 将状态代码转换为文本
	eo.Direction = DirectionMap[old.Direction] // 将方向代码转换为文本
	eo.Type = TypeMap.Value(old.Type)          // 将类型代码转换为文本
	return eo
}

// NewOrder 创建一个ExchangeOrder
func NewOrder() *ExchangeOrder {
	return &ExchangeOrder{}
}
