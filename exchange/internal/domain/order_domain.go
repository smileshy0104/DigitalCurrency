package domain

import (
	"common/db"
	"exchange/internal/repo"
)

// OrderDomain 交易货币模块
type OrderDomain struct {
	exchangeCoinRepo repo.ExchangeOrderRepo
}

// NewOrderDomain 创建交易货币模块
func NewOrderDomain(db *db.DB) *OrderDomain {
	return &OrderDomain{}
}
