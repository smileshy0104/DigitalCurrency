package logic

import (
	"context"
	"grpc-common/market/types/rate"
	"market/internal/domain"
	"market/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ExchangeRateLogic 用于处理汇率转换的逻辑
type ExchangeRateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	exchangeRateDomain *domain.ExchangeRateDomain
}

// UsdRate 获取美元汇率
// 该方法根据请求中的货币单位返回对应的美元汇率
// 参数 req 包含了请求的货币单位信息
// 返回值 RateRes 包含了对应的美元汇率，如果出现错误，返回错误信息
func (l *ExchangeRateLogic) UsdRate(req *rate.RateReq) (*rate.RateRes, error) {
	usdRate := l.exchangeRateDomain.UsdRate(req.Unit)
	return &rate.RateRes{
		Rate: usdRate,
	}, nil
}

// NewExchangeRateLogic 创建一个新的 ExchangeRateLogic 实例
// 参数 ctx 是上下文环境信息
// 参数 svcCtx 是服务的上下文信息，包含了服务所需的各种配置和初始化信息
// 返回值 ExchangeRateLogic 是一个新的汇率转换逻辑处理器实例
func NewExchangeRateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExchangeRateLogic {
	return &ExchangeRateLogic{
		ctx:                ctx,
		svcCtx:             svcCtx,
		Logger:             logx.WithContext(ctx),
		exchangeRateDomain: domain.NewExchangeRateDomain(),
	}
}
