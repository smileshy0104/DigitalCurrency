package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/market/types/market"
	"market-api/internal/svc"
	"market-api/internal/types"
	"time"
)

// MarketLogic 市场模块逻辑
type MarketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewMarketLogic 初始化市场模块逻辑
func NewMarketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarketLogic {
	return &MarketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// SymbolThumbTrend 获取市场行情的缩略趋势信息。
// 该方法首先尝试从缓存中获取数据，如果缓存未命中，则调用远程服务获取数据。
// 参数:
//
//	req *types.MarketReq - 包含请求信息的结构体，如IP地址。
//
// 返回值:
//
//	list []*types.CoinThumbResp - 一组币种的缩略信息。
//	err error - 错误信息，如果执行过程中遇到错误。
func (l *MarketLogic) SymbolThumbTrend(req *types.MarketReq) (list []*types.CoinThumbResp, err error) {
	// 定义一个 CoinThumb 类型的切片来存储缩略信息。
	var thumbs []*market.CoinThumb

	// 尝试从缓存中获取缩略信息。
	thumb := l.svcCtx.Processor.GetThumb()
	isCache := false

	// 检查缓存的数据是否为 CoinThumb 类型的切片。
	if thumb != nil {
		switch thumb.(type) {
		case []*market.CoinThumb:
			thumbs = thumb.([]*market.CoinThumb)
			isCache = true
		}
	}

	// 如果缓存中没有数据，调用远程服务获取数据。
	if !isCache {
		// 设置一个带有超时的上下文。
		ctx, cancelFunc := context.WithTimeout(l.ctx, 10*time.Second)
		defer cancelFunc()

		// 调用市场服务获取缩略信息。
		symbolThumbRes, err := l.svcCtx.MarketRpc.FindSymbolThumbTrend(ctx,
			&market.MarketReq{
				Ip: req.Ip,
			})
		if err != nil {
			return nil, err
		}

		// 将服务返回的数据赋值给 thumbs。
		thumbs = symbolThumbRes.List
	}

	// 将 thumbs 的数据复制到返回值 list 中。
	if err := copier.Copy(&list, thumbs); err != nil {
		return nil, err
	}

	// 返回结果。
	return
}

func (l *MarketLogic) SymbolThumb(req *types.MarketReq) (list []*types.CoinThumbResp, err error) {
	var thumbs []*market.CoinThumb
	thumb := l.svcCtx.Processor.GetThumb()
	if thumb != nil {
		switch thumb.(type) {
		case []*market.CoinThumb:
			thumbs = thumb.([]*market.CoinThumb)
		}
	}
	if err := copier.Copy(&list, thumbs); err != nil {
		return nil, err
	}
	return
}

func (l *MarketLogic) SymbolInfo(req types.MarketReq) (resp *types.ExchangeCoinResp, err error) {
	ctx, cancelFunc := context.WithTimeout(l.ctx, 10*time.Second)
	defer cancelFunc()
	esRes, err := l.svcCtx.MarketRpc.FindSymbolInfo(ctx,
		&market.MarketReq{
			Ip:     req.Ip,
			Symbol: req.Symbol,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.ExchangeCoinResp{}
	if err := copier.Copy(resp, esRes); err != nil {
		return nil, err
	}
	return
}

func (l *MarketLogic) CoinInfo(req *types.MarketReq) (*types.Coin, error) {
	ctx, cancel := context.WithTimeout(l.ctx, 5*time.Second)
	defer cancel()
	coin, err := l.svcCtx.MarketRpc.FindCoinInfo(ctx, &market.MarketReq{
		Unit: req.Unit,
	})
	if err != nil {
		return nil, err
	}
	ec := &types.Coin{}
	if err := copier.Copy(&ec, coin); err != nil {
		return nil, errors.New("数据格式有误")
	}
	return ec, nil
}

func (l *MarketLogic) History(req *types.MarketReq) (*types.HistoryKline, error) {
	ctx, cancel := context.WithTimeout(l.ctx, 10*time.Second)
	defer cancel()
	historyKline, err := l.svcCtx.MarketRpc.HistoryKline(ctx, &market.MarketReq{
		Symbol:     req.Symbol,
		From:       req.From,
		To:         req.To,
		Resolution: req.Resolution,
	})
	if err != nil {
		return nil, err
	}
	histories := historyKline.List
	var list = make([][]any, len(histories))
	for i, v := range histories {
		content := make([]any, 6)
		content[0] = v.Time
		content[1] = v.Open
		content[2] = v.High
		content[3] = v.Low
		content[4] = v.Close
		content[5] = v.Volume
		list[i] = content
	}
	return &types.HistoryKline{
		List: list,
	}, nil
}
