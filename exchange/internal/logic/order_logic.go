package logic

import (
	"context"
	"exchange/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

// ExchangeRateLogic 用于处理汇率转换的逻辑
type ExchangeRateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

// NewExchangeRateLogic 创建一个新的 ExchangeRateLogic 实例
// 参数 ctx 是上下文环境信息
// 参数 svcCtx 是服务的上下文信息，包含了服务所需的各种配置和初始化信息
// 返回值 ExchangeRateLogic 是一个新的汇率转换逻辑处理器实例
func NewExchangeRateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExchangeRateLogic {
	return &ExchangeRateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
