package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/market/types/market"
	"market/internal/domain"
	"market/internal/svc"
)

type MarketLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	exchangeCoinDomain *domain.ExchangeCoinDomain
	marketDomain       *domain.MarketDomain
	coinDomain         *domain.CoinDomain
}

func NewMarketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarketLogic {
	return &MarketLogic{
		ctx:                ctx,
		svcCtx:             svcCtx,
		Logger:             logx.WithContext(ctx),
		exchangeCoinDomain: domain.NewExchangeCoinDomain(svcCtx.Db),
		marketDomain:       domain.NewMarketDomain(svcCtx.MongoClient),
		coinDomain:         domain.NewCoinDomain(svcCtx.Db),
	}
}
func (l *MarketLogic) FindSymbolThumbTrend(req *market.MarketReq) (*market.SymbolThumbRes, error) {
	coins := l.exchangeCoinDomain.FindVisible(l.ctx)
	//查询mongo中相应的数据
	//查询1H间隔的 可以根据时间来进行查询 当天的价格变化趋势
	coinThumbs := l.marketDomain.SymbolThumbTrend(coins)
	//coinThumbs := make([]*market.CoinThumb, len(coins))
	//for i, v := range coins {
	//	ct := &market.CoinThumb{}
	//	ct.Symbol = v.Symbol
	//	trend := make([]float64, 0)
	//	for p := 0; p <= 24; p++ {
	//		trend = append(trend, rand.Float64())
	//	}
	//	ct.Trend = trend
	//	coinThumbs[i] = ct
	//}
	return &market.SymbolThumbRes{
		List: coinThumbs,
	}, nil
}

func (l *MarketLogic) FindSymbolInfo(req *market.MarketReq) (*market.ExchangeCoin, error) {
	exchangeCoin, err := l.exchangeCoinDomain.FindBySymbol(l.ctx, req.Symbol)
	if err != nil {
		return nil, err
	}
	ec := &market.ExchangeCoin{}
	copier.Copy(ec, exchangeCoin)
	return ec, nil
}
