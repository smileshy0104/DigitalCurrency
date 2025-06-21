package domain

import (
	"context"
	"job-center/internal/dao"
	"job-center/internal/database"
	"job-center/internal/model"
	"job-center/internal/repo"
	"log"
)

// KlineDomain kline模块
// 该结构体负责处理与kline相关的业务逻辑，包括kline数据的保存和查询等操作。
type KlineDomain struct {
	klineRepo repo.KlineRepo
}

// SaveBatch 批量保存kline数据
// 该方法首先将原始kline数据转换为Kline对象，然后删除数据库中已存在的较旧的kline数据，
// 最后将转换后的kline数据批量保存到数据库中。
// 参数:
//
//	data - 原始kline数据数组，每个元素代表一个时间点的kline数据。
//	symbol - 交易对标识，例如BTCUSDT。
//	period - kline周期，例如1m, 5m等。
func (d *KlineDomain) SaveBatch(data [][]string, symbol string, period string) {
	// 将原始kline数据转换为model.Kline结构体数组。
	klines := make([]*model.Kline, len(data))
	// 遍历原始数据，将每个元素转换为model.Kline结构体。
	for i, v := range data {
		klines[i] = model.NewKline(v, period)
	}
	// 删除数据库中已存在的较旧的kline数据。
	err := d.klineRepo.DeleteGtTime(context.Background(), klines[len(data)-1].Time, symbol, period)
	if err != nil {
		log.Println(err)
		return
	}
	// 保存新的kline数据。
	err = d.klineRepo.SaveBatch(context.Background(), klines, symbol, period)
	if err != nil {
		log.Println(err)
	}
}

// NewKlineDomain 实例化Kline模块
// 该函数创建并返回一个KlineDomain实例，用于处理kline相关的业务逻辑。
// 参数:
//
//	cli - MongoDB客户端，用于连接到数据库。
//
// 返回值:
//
//	*KlineDomain - Kline模块的实例。
func NewKlineDomain(cli *database.MongoClient) *KlineDomain {
	return &KlineDomain{
		klineRepo: dao.NewKlineDao(cli.Db),
	}
}
