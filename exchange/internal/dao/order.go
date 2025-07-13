package dao

import (
	"common/db"
	"common/db/gorms"
	"context"
	"exchange/internal/model"
)

type OrderDao struct {
	conn *gorms.GormConn
}

func (o *OrderDao) FindOrderHistory(ctx context.Context, symbol string, page int64, size int64, memberId int64) (list []*model.ExchangeOrder, total int64, err error) {
	//TODO implement me
	panic("implement me")
}

func NewOrderDao(db *db.DB) *OrderDao {
	return &OrderDao{
		conn: gorms.New(db.Conn),
	}
}
