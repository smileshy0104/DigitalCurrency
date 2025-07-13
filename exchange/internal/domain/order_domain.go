package domain

import (
	"common/db"
	"context"
	"exchange/internal/model"
	"exchange/internal/repo"
)

// OrderDomain 交易货币模块
type OrderDomain struct {
	OrderRepo repo.OrderRepo
}

func (d *OrderDomain) FindOrderHistory(ctx context.Context, symbol string, page int64, size int64, memberId int64) (list []*model.ExchangeOrder, total int64, err error) {
	return d.OrderRepo.FindOrderHistory(ctx, symbol, page, size, memberId)
}

// NewOrderDomain 创建交易货币模块
func NewOrderDomain(db *db.DB) *OrderDomain {
	return &OrderDomain{}
}
