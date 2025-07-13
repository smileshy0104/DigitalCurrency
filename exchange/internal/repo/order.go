package repo

import (
	"common/db"
	"context"
	"exchange/internal/model"
)

type OrderRepo interface {
	FindOrderHistory(ctx context.Context, symbol string, page int64, size int64, memberId int64) (list []*model.ExchangeOrder, total int64, err error)
	FindOrderCurrent(ctx context.Context, symbol string, page int64, size int64, memberId int64) (list []*model.ExchangeOrder, total int64, err error)
	FindCurrentTradingCount(ctx context.Context, id int64, symbol string, direction int) (int64, error)
	Save(ctx context.Context, conn db.DbConn, order *model.ExchangeOrder) error
}
