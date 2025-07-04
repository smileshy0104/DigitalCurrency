package domain

import (
	"common/op"
	"common/tools"
	"context"
	"grpc-common/market/types/market"
	"market/internal/dao"
	"market/internal/database"
	"market/internal/model"
	"market/internal/repo"
	"time"
)

type MarketDomain struct {
	klineRepo repo.KlineRepo
}

func NewMarketDomain(mongoClient *database.MongoClient) *MarketDomain {
	return &MarketDomain{
		klineRepo: dao.NewKlineDao(mongoClient.Db),
	}
}

// SymbolThumbTrend 根据币种信息生成缩略趋势图数据。
// 此函数主要用于为给定的币种数组生成一个包含价格趋势的缩略图数据数组。
// 参数:
//
//	coins - 一个ExchangeCoin对象的切片，代表需要生成缩略图数据的币种。
//
// 返回值:
//
//	一个CoinThumb对象的切片，每个对象包含一个币种的缩略图数据，包括价格趋势、最高价、最低价、成交量和成交额。
func (d *MarketDomain) SymbolThumbTrend(coins []*model.ExchangeCoin) []*market.CoinThumb {
	// 初始化一个CoinThumb切片，长度与输入的币种数量相同。
	coinThumbs := make([]*market.CoinThumb, len(coins))

	// 遍历币种coins 交易货币类型
	for i, v := range coins {
		// 定义时间范围，从零时刻到当前时间。
		from := tools.ZeroTime()
		end := time.Now().UnixMilli()

		// 查询1H间隔的K线数据，以获取当天的价格变化趋势。
		klines, err := d.klineRepo.FindBySymbolTime(context.Background(), v.Symbol, "1h", from, end, "")
		if err != nil {
			// 如果查询出错，使用默认的CoinThumb数据。
			coinThumbs[i] = model.DefaultCoinThumb(v.Symbol)
			continue
		}

		// 检查K线数据是否为空。
		length := len(klines)
		if length <= 0 {
			// 如果没有数据，使用默认的CoinThumb数据。
			coinThumbs[i] = model.DefaultCoinThumb(v.Symbol)
			continue
		}

		// 降序排列K线数据，以便最新数据在前。
		// 构建价格趋势数据。
		trend := make([]float64, length)
		var high float64 = 0
		var low float64 = klines[0].LowestPrice
		var volumes float64 = 0
		var turnover float64 = 0

		// 遍历K线数据，构建趋势数据并计算最高价、最低价、成交量和成交额。
		for i := length - 1; i >= 0; i-- {
			trend[i] = klines[i].ClosePrice
			// 获取当前K线的最高价。
			highestPrice := klines[i].HighestPrice
			// 判断当前K线是否是最高价。
			if highestPrice > high {
				high = highestPrice
			}
			// 获取最低价
			lowPrice := klines[i].LowestPrice
			// 判断当前K线是否是最低价。
			if lowPrice < low {
				low = lowPrice
			}
			// 累加成交量
			volumes = op.AddN(volumes, klines[i].Volume, 8)
			turnover = op.AddN(turnover, klines[i].Turnover, 8)
		}

		// 使用最新和最旧的K线数据来构建CoinThumb对象。
		newKline := klines[0]
		oldKline := klines[length-1]
		// 将K线数据转换为CoinThumb对象。
		thumb := newKline.ToCoinThumb(v.Symbol, oldKline)
		thumb.Trend = trend
		thumb.High = high
		thumb.Low = low
		thumb.Volume = volumes
		thumb.Turnover = turnover

		// 将构建好的CoinThumb对象添加到结果切片中。
		coinThumbs[i] = thumb
	}

	// 返回所有币种的缩略图数据。
	return coinThumbs
}

// HistoryKline 获取指定交易对在指定时间段内的K线历史数据。
// 该方法通过查询K线数据存储库，根据交易对、时间范围和周期来筛选数据，
// 并将查询到的数据转换为市场历史数据结构返回。
//
// 参数:
//
//	ctx - 上下文，用于传递请求范围的上下文信息。
//	symbol - 交易对符号，例如"BTCUSDT"。
//	from - 查询的起始时间戳。
//	to - 查询的结束时间戳。
//	period - K线周期，例如"1m"、"4h"、"1d"。
//
// 返回值:
//
//	[]*market.History - 一个指向市场历史数据切片的指针，包含K线数据。
//	error - 如果查询过程中发生错误，返回错误信息。
func (d *MarketDomain) HistoryKline(ctx context.Context, symbol string, from int64,
	to int64, period string) ([]*market.History, error) {
	// 调用K线数据存储库的FindBySymbolTime方法查询符合条件的K线数据。
	klines, err := d.klineRepo.FindBySymbolTime(ctx, symbol, period, from, to, "asc")
	if err != nil {
		return nil, err
	}
	// 初始化一个市场历史数据切片，长度为查询到的K线数据数量。
	list := make([]*market.History, len(klines))
	// 遍历查询到的K线数据，将每条数据转换为市场历史数据结构。
	for i, v := range klines {
		h := &market.History{}
		h.Time = v.Time
		h.Open = v.OpenPrice
		h.High = v.HighestPrice
		h.Low = v.LowestPrice
		h.Volume = v.Volume
		h.Close = v.ClosePrice
		list[i] = h
	}
	// 返回转换后的市场历史数据切片和nil错误。
	return list, nil
}
