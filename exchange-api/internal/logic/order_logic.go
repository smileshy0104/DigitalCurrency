package logic

import (
	"context"
	"exchange-api/internal/svc"
	"exchange-api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

// OrderLogic 定义了货币汇率相关的逻辑操作。
// 它嵌入了logx.Logger以支持日志记录，并依赖于svc.ServiceContext提供的服务上下文。
type OrderLogic struct {
	logx.Logger                     // 嵌入logx.Logger以支持日志记录
	ctx         context.Context     // 当前的上下文
	svcCtx      *svc.ServiceContext // 服务上下文，提供了访问其他服务或资源的上下文
}

func (l *OrderLogic) UsdRate() (*types.RateResponse, error) {
	return &types.RateResponse{}, nil
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
