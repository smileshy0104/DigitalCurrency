package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"market/internal/model"
)

// KlineDao 关于kline的数据库操作
type KlineDao struct {
	db *mongo.Database
}

// SaveBatch 保存kline
// ctx 上下文信息
// data 要保存的kline数据
// symbol 货币类型
// period kline周期
func (k *KlineDao) SaveBatch(ctx context.Context, data []*model.Kline, symbol, period string) error {
	// mongobd数据库操作
	mk := &model.Kline{}
	collection := k.db.Collection(mk.Table(symbol, period))
	ds := make([]interface{}, len(data))
	for i, v := range data {
		ds[i] = v
	}
	// 批量插入
	_, err := collection.InsertMany(ctx, ds)
	return err
}

// DeleteGtTime 删除大于指定时间
// ctx 上下文信息
// time 指定的时间
// symbol 货币类型
// period kline周期
func (k *KlineDao) DeleteGtTime(ctx context.Context, time int64, symbol string, period string) error {
	mk := &model.Kline{}
	collection := k.db.Collection(mk.Table(symbol, period))
	// 批量删除指定时间之前的数据
	deleteResult, err := collection.DeleteMany(ctx, bson.D{{"time", bson.D{{"$gte", time}}}})
	if err != nil {
		return err
	}
	log.Printf("%s %s 删除了%d条数据 \n", symbol, period, deleteResult.DeletedCount)
	return nil
}

// FindBySymbol 查询指定货币类型symbol的kline
// ctx 上下文信息
// symbol 货币类型
// period kline周期
// count 查询的数量
func (k *KlineDao) FindBySymbol(ctx context.Context, symbol, period string, count int64) (list []*model.Kline, err error) {
	//按照时间 降序排列
	mk := &model.Kline{}
	collection := k.db.Collection(mk.Table(symbol, period))
	// 查询指定数量的数据
	cur, err := collection.Find(ctx, bson.D{{}}, &options.FindOptions{
		Limit: &count,
		Sort:  bson.D{{"time", -1}},
	})
	if err != nil {
		return nil, err
	}
	// 获取数据
	err = cur.All(ctx, &list)
	if err != nil {
		return nil, err
	}
	return
}

// FindBySymbolTime 根据符号和时间范围查找K线数据。
// 该方法从数据库中查询指定符号、周期和时间范围内的K线数据，并根据指定的排序方式返回结果。
// 参数:
//
//	ctx: 上下文，用于传递请求范围的 deadline、取消信号、身份验证信息等。
//	symbol: 交易对符号，例如 "BTCUSDT"。
//	period: K线周期，例如 "1m"、"4h"。
//	from: 查询时间范围的起始时间戳。
//	end: 查询时间范围的结束时间戳。
//	sort: 排序方式，"asc" 表示升序，其他值表示降序。
//
// 返回值:
//
//	list: 包含K线数据的切片。
//	err: 错误信息，如果执行过程中遇到错误则返回。
func (k *KlineDao) FindBySymbolTime(ctx context.Context, symbol, period string, from, end int64, sort string) (list []*model.Kline, err error) {
	// 安装时间范围 查询
	mk := &model.Kline{}
	sortInt := -1
	if "asc" == sort {
		sortInt = 1
	}
	collection := k.db.Collection(mk.Table(symbol, period))
	// 查询指定时间范围的数据
	cur, err := collection.Find(ctx,
		bson.D{{"time", bson.D{{"$gte", from}, {"$lte", end}}}},
		&options.FindOptions{
			Sort: bson.D{{"time", sortInt}},
		})
	if err != nil {
		return nil, err
	}
	// 获取数据
	err = cur.All(ctx, &list)
	if err != nil {
		return nil, err
	}
	return
}

// NewKlineDao 创建KlineDao实例
// db MongoDB的数据库实例
func NewKlineDao(db *mongo.Database) *KlineDao {
	return &KlineDao{
		db: db,
	}
}
