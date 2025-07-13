package logic

import (
	"common/pages"
	"context"
	"errors"
	"exchange-api/internal/svc"
	"exchange-api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/exchange/types/order"
	"time"
)

// OrderLogic 定义了货币汇率相关的逻辑操作。
// 它嵌入了logx.Logger以支持日志记录，并依赖于svc.ServiceContext提供的服务上下文。
type OrderLogic struct {
	logx.Logger                     // 嵌入logx.Logger以支持日志记录
	ctx         context.Context     // 当前的上下文
	svcCtx      *svc.ServiceContext // 服务上下文，提供了访问其他服务或资源的上下文
}

// NewOrderLogic 创建并返回一个新的OrderLogic实例。
// 它需要一个context和一个ServiceContext作为输入参数。
func NewOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderLogic {
	return &OrderLogic{
		Logger: logx.WithContext(ctx), // 使用给定的context配置Logger
		ctx:    ctx,                   // 设置当前上下文
		svcCtx: svcCtx,                // 设置服务上下文
	}
}

// History 方法用于获取订单历史记录。
func (l *OrderLogic) History(req *types.ExchangeReq) (*pages.PageResult, error) {
	// 创建一个带有超时的上下文，以确保请求不会无限期地等待。
	// 这里设置的超时时间是5秒，旨在防止在服务调用响应缓慢时导致资源浪费或潜在的死锁情况。
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// 在函数返回前取消创建的上下文，以释放相关资源。
	defer cancel()

	// 从当前上下文中提取用户ID。
	// 注意：这里假设了上下文中已经设置了"userId"，且其能成功转换为int64类型。
	userId := l.ctx.Value("userId").(int64)

	// 通过RPC调用资产服务，查找用户指定货币的钱包信息。
	// 这里将用户ID和请求的货币名称作为参数传递给服务。
	orderRes, err := l.svcCtx.OrderRpc.FindOrderHistory(ctx, &order.OrderReq{
		Symbol:   req.Symbol,
		Page:     req.PageNo,
		PageSize: req.PageSize,
		UserId:   userId,
	})
	if err != nil {
		return nil, err
	}
	// 创建一个切片，用于存储返回的订单信息
	list := orderRes.List
	b := make([]any, len(list))
	for i := range list {
		b[i] = list[i]
	}
	return pages.New(b, req.PageNo, req.PageSize, orderRes.Total), nil
}

// Current 获取当前订单（交易中）
func (l *OrderLogic) Current(req *types.ExchangeReq) (*pages.PageResult, error) {
	// 创建一个带有超时的上下文，以确保请求不会无限期地等待。
	// 这里设置的超时时间是5秒，旨在防止在服务调用响应缓慢时导致资源浪费或潜在的死锁情况。
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// 在函数返回前取消创建的上下文，以释放相关资源。
	defer cancel()

	// 从当前上下文中提取用户ID。
	// 注意：这里假设了上下文中已经设置了"userId"，且其能成功转换为int64类型。
	userId := l.ctx.Value("userId").(int64)

	// 通过RPC调用资产服务，查找用户指定货币的钱包信息。
	// 这里将用户ID和请求的货币名称作为参数传递给服务。
	orderRes, err := l.svcCtx.OrderRpc.FindOrderCurrent(ctx, &order.OrderReq{
		Symbol:   req.Symbol,
		Page:     req.PageNo,
		PageSize: req.PageSize,
		UserId:   userId,
	})
	if err != nil {
		return nil, err
	}
	// 创建一个切片，用于存储返回的订单信息
	list := orderRes.List
	b := make([]any, len(list))
	for i := range list {
		b[i] = list[i]
	}
	return pages.New(b, req.PageNo, req.PageSize, orderRes.Total), nil
}

// AddOrder 添加订单
// 该方法从上下文中获取用户ID，并验证订单请求的合法性
// 如果请求合法，它将调用订单RPC服务来添加订单
func (l *OrderLogic) AddOrder(req *types.ExchangeReq) (string, error) {
	// 获取用户ID
	value := l.ctx.Value("userId").(int64)

	// 校验请求参数
	if !req.OrderValid() {
		return "", errors.New("参数传递错误")
	}

	// Add 添加订单
	orderRes, err := l.svcCtx.OrderRpc.Add(l.ctx, &order.OrderReq{
		Symbol:    req.Symbol,    // 交易对
		UserId:    value,         // 用户ID
		Direction: req.Direction, // 订单方向 0 买 1 卖
		Type:      req.Type,      // 挂单类型 0 市场价 1 最低价
		Price:     req.Price,     // 交易价格
		Amount:    req.Amount,    // 交易数量
	})
	if err != nil {
		return "", err
	}

	return orderRes.OrderId, nil
}
