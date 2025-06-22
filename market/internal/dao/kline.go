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

// FindBySymbolTime 时间范围查询
// ctx 上下文信息
// symbol 货币类型
// period kline周期
// from 查询的开始时间
// end 查询的结束时间
// sort 排序方式，"asc"为升序，其他为降序
func (k *KlineDao) FindBySymbolTime(ctx context.Context, symbol, period string, from, end int64, sort string) (list []*model.Kline, err error) {
	//安装时间范围 查询
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
