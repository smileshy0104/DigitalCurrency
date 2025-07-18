package processor

import (
	"exchange/internal/model"
	"sync"
)

// TradeTimeQueue 定义一个交易订单的时间队列
type TradeTimeQueue []*model.ExchangeOrder

// Len 返回队列的长度
func (t TradeTimeQueue) Len() int {
	return len(t)
}

// Less 返回两个订单的时间比较结果，升序排列
func (t TradeTimeQueue) Less(i, j int) bool {
	// 升序：较早的时间在前
	return t[i].Time < t[j].Time
}

// Swap 交换队列中两个订单的位置
func (t TradeTimeQueue) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

// LimitPriceQueue 定义一个限价队列，包含读写锁和交易队列
type LimitPriceQueue struct {
	mux  sync.RWMutex // 读写锁，用于并发安全
	list TradeQueue   // 限价队列
}

// LimitPriceMap 定义一个限价映射，包含价格和对应的订单列表
type LimitPriceMap struct {
	price float64                // 限价
	list  []*model.ExchangeOrder // 对应的订单列表
}

// TradeQueue 定义一个交易队列，用于降序排列
type TradeQueue []*LimitPriceMap

// Len 返回交易队列的长度
func (t TradeQueue) Len() int {
	return len(t)
}

// Less 返回两个限价映射的价格比较结果，降序排列
func (t TradeQueue) Less(i, j int) bool {
	// 降序：较高的价格在前
	return t[i].price > t[j].price
}

// Swap 交换交易队列中两个限价映射的位置
func (t TradeQueue) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

// TradePlate 盘口信息
type TradePlate struct {
	Items     []*TradePlateItem `json:"items"`
	Symbol    string
	direction int
	maxDepth  int
	mux       sync.RWMutex
}

type TradePlateItem struct {
	Price  float64 `json:"price"`
	Amount float64 `json:"amount"`
}
