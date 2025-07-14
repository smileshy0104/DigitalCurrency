package domain

import (
	"common/db"
	"common/op"
	"common/tools"
	"context"
	"errors"
	"exchange/internal/dao"
	"exchange/internal/model"
	"exchange/internal/repo"
	"grpc-common/market/mk_client"
	"grpc-common/ucenter/uc_client"
	"time"
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

// AddOrder 添加订单信息到数据库。
// 该函数接收一个上下文、数据库连接、订单信息、币种信息以及两个钱包对象作为参数，
// 并根据订单类型和方向计算所需冻结的资金或币量，最后返回计算结果或错误信息。
// 参数:
//
//	ctx - 上下文，用于传递请求范围的数据、取消信号等。
//	conn - 数据库连接对象，用于执行数据库操作。
//	order - 订单信息对象，包含订单的相关数据。
//	coin - 币种信息对象，包含币种的相关数据。
//	baseWallet - 主钱包对象，用于买单时扣除资金。
//	coinWallet - 币种钱包对象，用于卖单时扣除币量。
//
// 返回值:
//
//	float64 - 计算得到的需冻结的资金或币量。
//	error - 错误信息，如果执行过程中遇到错误则返回。
func (d *OrderDomain) AddOrder(ctx context.Context, conn db.DbConn, order *model.ExchangeOrder,
	coin *mk_client.ExchangeCoin, baseWallet *uc_client.MemberWallet, coinWallet *uc_client.MemberWallet) (float64, error) {
	// 初始化订单状态、成交量和下单时间
	order.Status = model.Init
	order.TradedAmount = 0
	order.Time = time.Now().UnixMilli()
	// 生成唯一的订单ID
	order.OrderId = tools.Unq("E")

	// 交易时暂时不考虑手续费
	// 定义变量money来存储根据订单类型和方向计算出的需冻结的资金或币量
	var money float64

	// 根据订单方向判断是买入还是卖出
	if order.Direction == model.BUY {
		// 如果是市价单，直接使用订单的amount作为冻结资金
		if order.Type == model.MarketPrice {
			money = order.Amount
		} else {
			// 如果是限价单，计算冻结资金，这里使用op.MulFloor函数来避免精度损失
			money = op.MulFloor(order.Price, order.Amount, 8)
		}
		// 检查主钱包余额是否足够冻结
		if baseWallet.Balance < money {
			return 0, errors.New("余额不足")
		}
	} else {
		// 如果是卖单，直接使用订单的amount作为冻结币量
		money = order.Amount
		// 检查币种钱包余额是否足够冻结
		if coinWallet.Balance < money {
			return 0, errors.New("余额不足")
		}
	}

	// 保存订单
	err := d.OrderRepo.Save(ctx, conn, order)

	// 函数执行完毕，返回计算出的冻结资金或币量以及nil作为错误信息
	return money, err
}

// NewOrderDomain 创建交易货币模块
func NewOrderDomain(db *db.DB) *OrderDomain {
	return &OrderDomain{
		OrderRepo: dao.NewOrderDao(db),
	}
}
