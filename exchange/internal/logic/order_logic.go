package logic

import (
	"context"
	"errors"
	"exchange/internal/domain"
	"exchange/internal/model"
	"exchange/internal/svc"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/exchange/types/order"
	"grpc-common/ucenter/types/member"
)

// OrderLogic 用于处理汇率转换的逻辑
type OrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	orderDomain *domain.OrderDomain
}

// NewOrderLogic 创建一个新的 OrderLogic 实例
// 参数 ctx 是上下文环境信息
// 参数 svcCtx 是服务的上下文信息，包含了服务所需的各种配置和初始化信息
func NewOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderLogic {
	return &OrderLogic{
		ctx:         ctx,
		svcCtx:      svcCtx,
		Logger:      logx.WithContext(ctx),
		orderDomain: domain.NewOrderDomain(svcCtx.Db),
	}
}

// FindOrderHistory 查询指定用户的历史订单
func (l *OrderLogic) FindOrderHistory(req *order.OrderReq) (*order.OrderRes, error) {
	// 查询对应用户的历史订单
	orderList, total, err := l.orderDomain.FindOrderHistory(l.ctx, req.Symbol, req.Page, req.PageSize, req.UserId)
	if err != nil {
		return nil, err
	}
	var list []*order.ExchangeOrder
	copier.Copy(&list, orderList)
	return &order.OrderRes{
		List:  list,
		Total: total,
	}, nil
}

// FindOrderCurrent 查询当前用户订单（交易中）
func (l *OrderLogic) FindOrderCurrent(req *order.OrderReq) (*order.OrderRes, error) {
	// 查询对应用户的历史订单
	orderList, total, err := l.orderDomain.FindOrderCurrent(l.ctx, req.Symbol, req.Page, req.PageSize, req.UserId)
	if err != nil {
		return nil, err
	}
	var list []*order.ExchangeOrder
	copier.Copy(&list, orderList)
	return &order.OrderRes{
		List:  list,
		Total: total,
	}, nil
}

// Add 添加订单
func (l *OrderLogic) Add(req *order.OrderReq) (*order.AddOrderRes, error) {
	// 通过用户id查询用户信息，判断用户是否存在
	memberInfo, err := l.svcCtx.MemberRpc.FindMemberById(l.ctx, &member.MemberReq{
		MemberId: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	if memberInfo.TransactionStatus == 0 {
		return nil, errors.New("此用户已经被禁止交易")
	}
	if req.Type == model.TypeMap[model.LimitPrice] && req.Price <= 0 {
		return nil, errors.New("限价模式下价格不能小于等于0")
	}
	if req.Amount <= 0 {
		return nil, errors.New("数量不能小于等于0")
	}
	return &order.AddOrderRes{}, nil
}

func (l *OrderLogic) FindByOrderId(req *order.OrderReq) (*order.ExchangeOrderOrigin, error) {
	return nil, nil
}

func (l *OrderLogic) CancelOrder(req *order.OrderReq) (*order.CancelOrderRes, error) {
	return nil, nil
}
