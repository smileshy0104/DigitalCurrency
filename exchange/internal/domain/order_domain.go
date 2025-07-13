package domain

import (
	"common/db"
	"context"
	"exchange/internal/dao"
	"exchange/internal/model"
	"exchange/internal/repo"
	"grpc-common/market/mk_client"
	"grpc-common/ucenter/uc_client"
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

// FindCurrentTradingCount 查询当前用户在指定交易对下的订单数（交易中的订单）。
func (d *OrderDomain) FindCurrentTradingCount(ctx context.Context, userId int64, symbol string, direction string) (int64, error) {
	return d.OrderRepo.FindCurrentTradingCount(ctx, userId, symbol, model.DirectionMap.Code(direction))
}

// AddOrder 添加订单
func (d *OrderDomain) AddOrder(ctx context.Context, conn db.DbConn, order *model.ExchangeOrder,
	coin *mk_client.ExchangeCoin, baseWallet *uc_client.MemberWallet, coinWallet *uc_client.MemberWallet) (float64, error) {

	return 0, nil
}

// NewOrderDomain 创建交易货币模块
func NewOrderDomain(db *db.DB) *OrderDomain {
	return &OrderDomain{
		OrderRepo: dao.NewOrderDao(db),
	}
}
