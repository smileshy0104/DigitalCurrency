package logic

import (
	"context"
	"exchange/internal/domain"
	"exchange/internal/svc"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/exchange/types/order"
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

func (l *OrderLogic) FindOrderCurrent(req *order.OrderReq) (*order.OrderRes, error) {
	return &order.OrderRes{}, nil
}

func (l *OrderLogic) Add(req *order.OrderReq) (*order.AddOrderRes, error) {
	return &order.AddOrderRes{}, nil
}

func (l *OrderLogic) FindByOrderId(req *order.OrderReq) (*order.ExchangeOrderOrigin, error) {
	return nil, nil
}

func (l *OrderLogic) CancelOrder(req *order.OrderReq) (*order.CancelOrderRes, error) {
	return nil, nil
}
