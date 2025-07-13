package logic

import (
	"context"
	"exchange/internal/domain"
	"exchange/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/exchange/types/order"
)

// ExchangeOrderLogic 用于处理汇率转换的逻辑
type ExchangeOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	orderDomain *domain.OrderDomain
}

// NewExchangeOrderLogic 创建一个新的 ExchangeOrderLogic 实例
// 参数 ctx 是上下文环境信息
// 参数 svcCtx 是服务的上下文信息，包含了服务所需的各种配置和初始化信息
func NewExchangeOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExchangeOrderLogic {
	return &ExchangeOrderLogic{
		ctx:         ctx,
		svcCtx:      svcCtx,
		Logger:      logx.WithContext(ctx),
		orderDomain: domain.NewOrderDomain(svcCtx.Db),
	}
}

// FindOrderHistory 查询指定用户的历史订单
func (l *ExchangeOrderLogic) FindOrderHistory(req *order.OrderReq) (*order.OrderRes, error) {
	return &order.OrderRes{}, nil
}

func (l *ExchangeOrderLogic) FindOrderCurrent(req *order.OrderReq) (*order.OrderRes, error) {
	return &order.OrderRes{}, nil
}

func (l *ExchangeOrderLogic) Add(req *order.OrderReq) (*order.AddOrderRes, error) {
	return &order.AddOrderRes{}, nil
}

func (l *ExchangeOrderLogic) FindByOrderId(req *order.OrderReq) (*order.ExchangeOrderOrigin, error) {
	return nil, nil
}

func (l *ExchangeOrderLogic) CancelOrder(req *order.OrderReq) (*order.CancelOrderRes, error) {
	return nil, nil
}
