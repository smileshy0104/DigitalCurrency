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

	// TODO 尝试从缓存中获取缩略信息。（还是读取mongodb中的数据，但是这个数据在程序初始化时就调用了FindSymbolThumbTrend方法）
	thumb := l.svcCtx.Processor.GetThumb()
	// 初始缓存标识为false
	isCache := false

	// 检查缓存的数据是否为 CoinThumb 类型的切片。
	if thumb != nil {
		// 尝试将缓存数据转换为 CoinThumb 类型的切片。
		switch thumb.(type) {
		case []*market.CoinThumb:
			// 将缓存数据转换为 CoinThumb 类型的切片。
			thumbs = thumb.([]*market.CoinThumb)
			// 设置缓存标识为true
			isCache = true
		}
	}

	// TODO 如果缓存中没有数据，调用远程服务获取数据。
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

// SymbolThumb 获取市场行情的缩略信息。
func (l *MarketLogic) SymbolThumb(req *types.MarketReq) (list []*types.CoinThumbResp, err error) {
	var thumbs []*market.CoinThumb
	// 尝试从缓存中获取数据。
	thumb := l.svcCtx.Processor.GetThumb()
	// 检查缓存的数据是否为 CoinThumb 类型的切片。
	if thumb != nil {
		// 若缓存的数据为 CoinThumb 类型的切片，则将其赋值给 thumbs。
		switch thumb.(type) {
		case []*market.CoinThumb:
			thumbs = thumb.([]*market.CoinThumb)
		}
	}
	// 缓存中没有数据，则调用远程服务获取数据。
	if err := copier.Copy(&list, thumbs); err != nil {
		return nil, err
	}
	return
}

// SymbolInfo 查询指定市场的符号信息。
// 该方法通过RPC调用获取指定符号（Symbol）的市场信息，并返回相关信息。
// 参数:
//
//	req - 包含请求参数的MarketReq对象，包括IP地址和市场符号。
//
// 返回值:
//
//	*types.ExchangeCoinResp - 包含市场符号信息的响应对象。
//	error - 如果发生错误，则返回错误信息。
func (l *MarketLogic) SymbolInfo(req types.MarketReq) (resp *types.ExchangeCoinResp, err error) {
	// 创建一个带有超时的上下文，以确保请求不会无限期地等待。
	// 超时设置为10秒，以平衡响应速度和等待时间。
	ctx, cancelFunc := context.WithTimeout(l.ctx, 10*time.Second)
	// 在函数返回时取消上下文，以释放相关资源。
	defer cancelFunc()

	// 调用MarketRpc服务的FindSymbolInfo方法，根据IP和Symbol获取市场信息。
	esRes, err := l.svcCtx.MarketRpc.FindSymbolInfo(ctx,
		&market.MarketReq{
			Ip:     req.Ip,
			Symbol: req.Symbol,
		})
	// 如果发生错误，返回nil和错误信息。
	if err != nil {
		return nil, err
	}

	// 初始化一个ExchangeCoinResp对象来存储转换后的响应数据。
	resp = &types.ExchangeCoinResp{}
	// 使用copier.Copy将RPC调用的结果复制到响应对象中。
	// 如果复制过程中发生错误，返回nil和错误信息。
	if err := copier.Copy(resp, esRes); err != nil {
		return nil, err
	}

	// 返回填充好的响应对象和nil错误，表示操作成功。
	return
}

// CoinInfo 获取指定货币的市场信息。
// 该方法通过RPC调用MarketRpc服务来查询货币信息，并将结果转换为types.Coin类型返回。
// 参数req是包含查询条件的请求对象，如货币单位。
// 返回值是一个填充了查询结果的types.Coin对象，以及一个错误对象，如果执行过程中遇到任何问题，则返回相应的错误。
func (l *MarketLogic) CoinInfo(req *types.MarketReq) (*types.Coin, error) {
	// 创建一个带有超时的上下文，以确保请求不会无限期地等待。
	// 超时设置为5秒，这是与外部服务通信的一个常见合理限制。
	ctx, cancel := context.WithTimeout(l.ctx, 5*time.Second)
	// 确保在函数退出时取消上下文，释放相关资源。
	defer cancel()

	// 调用MarketRpc服务的FindCoinInfo方法获取货币信息。
	// 这里传递的是从上下文中创建的带有超时的ctx，以及一个包含了查询条件的MarketReq对象。
	coin, err := l.svcCtx.MarketRpc.FindCoinInfo(ctx, &market.MarketReq{
		Unit: req.Unit,
	})
	// 如果查询过程中出现错误，返回nil和错误信息。
	if err != nil {
		return nil, err
	}

	// 创建一个types.Coin类型的对象来存储查询结果。
	ec := &types.Coin{}
	// 使用copier库将查询结果从RPC调用复制到types.Coin对象中。
	// 这里处理数据转换，如果转换过程中出现错误，返回一个描述性的错误信息。
	if err := copier.Copy(&ec, coin); err != nil {
		return nil, errors.New("数据格式有误")
	}

	// 成功返回填充了查询结果的types.Coin对象。
	return ec, nil
}

func (l *MarketLogic) AllCoinInfo(req *types.MarketReq) ([]*types.Coin, error) {
	// 创建一个带有超时的上下文，以确保请求不会无限期地等待。
	// 超时设置为5秒，这是与外部服务通信的一个常见合理限制。
	ctx, cancel := context.WithTimeout(l.ctx, 5*time.Second)
	// 确保在函数退出时取消上下文，释放相关资源。
	defer cancel()

	// 调用MarketRpc服务的FindCoinInfo方法获取货币信息。
	// 这里传递的是从上下文中创建的带有超时的ctx，以及一个包含了查询条件的MarketReq对象。
	coin, err := l.svcCtx.MarketRpc.FindAllCoin(ctx, &market.MarketReq{})
	// 如果查询过程中出现错误，返回nil和错误信息。
	if err != nil {
		return nil, err
	}

	// 创建一个types.Coin类型的对象来存储查询结果。
	ec := []*types.Coin{}
	// 使用copier库将查询结果从RPC调用复制到types.Coin对象中。
	// 这里处理数据转换，如果转换过程中出现错误，返回一个描述性的错误信息。
	if err := copier.Copy(&ec, coin.List); err != nil {
		return nil, errors.New("数据格式有误")
	}

	// 成功返回填充了查询结果的types.Coin对象。
	return ec, nil
}

func (l *MarketLogic) CoinInfoById(req *types.MarketReq) (*types.Coin, error) {
	// 创建一个带有超时的上下文，以确保请求不会无限期地等待。
	// 超时设置为5秒，这是与外部服务通信的一个常见合理限制。
	ctx, cancel := context.WithTimeout(l.ctx, 5*time.Second)
	// 确保在函数退出时取消上下文，释放相关资源。
	defer cancel()

	// 调用MarketRpc服务的FindCoinInfo方法获取货币信息。
	// 这里传递的是从上下文中创建的带有超时的ctx，以及一个包含了查询条件的MarketReq对象。
	coin, err := l.svcCtx.MarketRpc.FindCoinById(ctx, &market.MarketReq{
		Id: req.Id,
	})
	// 如果查询过程中出现错误，返回nil和错误信息。
	if err != nil {
		return nil, err
	}

	// 创建一个types.Coin类型的对象来存储查询结果。
	ec := &types.Coin{}
	// 使用copier库将查询结果从RPC调用复制到types.Coin对象中。
	// 这里处理数据转换，如果转换过程中出现错误，返回一个描述性的错误信息。
	if err := copier.Copy(&ec, coin); err != nil {
		return nil, errors.New("数据格式有误")
	}

	// 成功返回填充了查询结果的types.Coin对象。
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
