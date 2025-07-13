package repo

import (
	"context"
	"exchange/internal/model"
)

type OrderRepo interface {
	FindOrderHistory(ctx context.Context, symbol string, page int64, size int64, memberId int64) (list []*model.ExchangeOrder, total int64, err error)
}
