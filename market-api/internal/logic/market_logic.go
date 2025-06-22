package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/market/types/market"
	"market-api/internal/svc"
	"market-api/internal/types"
	"time"
)

type MarketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMarketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarketLogic {
	return &MarketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MarketLogic) SymbolThumbTrend(req *types.MarketReq) (list []*types.CoinThumbResp, err error) {
	// 创建一个带有超时的上下文，以确保请求不会无限期地等待。
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // 确保在函数退出时取消创建的上下文。

	// 调用ExchangeRateRpc服务的UsdRate方法获取汇率信息。
	symbolThumbRes, err := l.svcCtx.MarketRpc.FindSymbolThumbTrend(ctx, &market.MarketReq{Ip: req.Ip})
	if err != nil {
		return nil, err // 如果发生错误，返回nil和错误信息。
	}
	if err := copier.Copy(&list, symbolThumbRes.List); err != nil {
		return nil, err
	}
	return
}
