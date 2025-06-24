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

func (d *MarketDomain) HistoryKline(
	ctx context.Context,
	symbol string,
	from int64,
	to int64,
	period string) ([]*market.History, error) {
	klines, err := d.klineRepo.FindBySymbolTime(ctx, symbol, period, from, to, "asc")
	if err != nil {
		return nil, err
	}
	list := make([]*market.History, len(klines))
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
	return list, nil
}
