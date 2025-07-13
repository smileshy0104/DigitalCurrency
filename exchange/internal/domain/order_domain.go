package domain

import (
	"common/db"
	"context"
	"exchange/internal/dao"
	"exchange/internal/model"
	"exchange/internal/repo"
)

// OrderDomain 交易货币模块
type OrderDomain struct {
	OrderRepo repo.OrderRepo
}

// FindOrderHistory 获取对应用户的订单记录
func (d *OrderDomain) FindOrderHistory(ctx context.Context, symbol string, page int64, size int64, memberId int64) (list []*model.ExchangeOrderVo, total int64, err error) {
	// 获取订单历史
	orderHistory, total, err := d.OrderRepo.FindOrderHistory(ctx, symbol, page, size, memberId)
	if err != nil {
		return nil, 0, err
	}
	// 进行数据处理，然后返回指定的数据结构
	voList := make([]*model.ExchangeOrderVo, len(list))
	for i, v := range orderHistory {
		voList[i] = v.ToVo()
	}

	return voList, total, nil
}

// FindOrderCurrent 查询当前订单（交易中）
func (d *OrderDomain) FindOrderCurrent(ctx context.Context, symbol string, page int64, size int64, memberId int64) (list []*model.ExchangeOrderVo, total int64, err error) {
	// 获取订单历史
	orderHistory, total, err := d.OrderRepo.FindOrderCurrent(ctx, symbol, page, size, memberId)
	if err != nil {
		return nil, 0, err
	}
	// 进行数据处理，然后返回指定的数据结构
	voList := make([]*model.ExchangeOrderVo, len(list))
	for i, v := range orderHistory {
		voList[i] = v.ToVo()
	}

	return voList, total, nil
}

// NewOrderDomain 创建交易货币模块
func NewOrderDomain(db *db.DB) *OrderDomain {
	return &OrderDomain{
		OrderRepo: dao.NewOrderDao(db),
	}
}
