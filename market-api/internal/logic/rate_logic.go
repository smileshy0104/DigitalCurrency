package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/market/types/rate"
	"market-api/internal/svc"
	"market-api/internal/types"
	"time"
)

// ExchangeRateLogic 定义了货币汇率相关的逻辑操作。
// 它嵌入了logx.Logger以支持日志记录，并依赖于svc.ServiceContext提供的服务上下文。
type ExchangeRateLogic struct {
	logx.Logger                     // 嵌入logx.Logger以支持日志记录
	ctx         context.Context     // 当前的上下文
	svcCtx      *svc.ServiceContext // 服务上下文，提供了访问其他服务或资源的上下文
}

// UsdRate 处理USD汇率请求。
// 它接收一个RateRequest作为输入，并返回一个RateResponse作为输出，或者一个错误。
// 此函数通过ExchangeRateRpc服务获取USD汇率，并将结果封装在RateResponse中返回。
func (l *ExchangeRateLogic) UsdRate(req *types.RateRequest) (*types.RateResponse, error) {
	// 创建一个带有超时的上下文，以确保请求不会无限期地等待。
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // 确保在函数退出时取消创建的上下文。

	// 调用ExchangeRateRpc服务的UsdRate方法获取汇率信息。
	rateRes, err := l.svcCtx.ExchangeRateRpc.UsdRate(ctx, &rate.RateReq{
		Unit: req.Unit,
		Ip:   req.Ip,
	})
	if err != nil {
		return nil, err // 如果发生错误，返回nil和错误信息。
	}

	// 成功时，将获取的汇率信息封装在RateResponse中并返回。
	return &types.RateResponse{
		Rate: rateRes.Rate,
	}, nil
}

// NewExchangeRateLogic 创建并返回一个新的ExchangeRateLogic实例。
// 它需要一个context和一个ServiceContext作为输入参数。
func NewExchangeRateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExchangeRateLogic {
	return &ExchangeRateLogic{
		Logger: logx.WithContext(ctx), // 使用给定的context配置Logger
		ctx:    ctx,                   // 设置当前上下文
		svcCtx: svcCtx,                // 设置服务上下文
	}
}
