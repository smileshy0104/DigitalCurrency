package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"job-center/internal/model"
	"log"
)

// KlineDao 定义了用于操作K线数据的MongoDB数据库访问对象。
type KlineDao struct {
	db *mongo.Database
}

// SaveBatch 批量保存K线数据到MongoDB中。
// 该方法接收一个上下文、K线数据切片、交易对符号和周期作为参数，并返回可能发生的错误。
func (k *KlineDao) SaveBatch(ctx context.Context, data []*model.Kline, symbol, period string) error {
	// 创建Kline实例以获取正确的集合名称。
	mk := &model.Kline{}
	collection := k.db.Collection(mk.Table(symbol, period))

	// 将数据转换为interface{}切片，以便InsertMany方法接受。
	ds := make([]interface{}, len(data))
	for i, v := range data {
		ds[i] = v
	}

	// 执行批量插入并处理结果。
	_, err := collection.InsertMany(ctx, ds)
	return err
}

// DeleteGtTime 删除集合中所有时间大于等于给定时间的数据。
// 该方法接收一个上下文、时间戳、交易对符号和周期作为参数，并返回可能发生的错误。
func (k *KlineDao) DeleteGtTime(ctx context.Context, time int64, symbol string, period string) error {
	// 创建Kline实例以获取正确的集合名称。
	mk := &model.Kline{}
	collection := k.db.Collection(mk.Table(symbol, period))

	// 执行删除操作并处理结果。
	// 使用bson.D结构体创建删除条件。bson.D{{"time", bson.D{{"$gte", time}}}} 表示删除时间大于等于给定时间的数据。
	deleteResult, err := collection.DeleteMany(ctx, bson.D{{"time", bson.D{{"$gte", time}}}})
	if err != nil {
		return err
	}

	// 记录删除操作的日志。
	log.Printf("%s %s 删除了%d条数据 \n", symbol, period, deleteResult.DeletedCount)
	return nil
}

// NewKlineDao 创建并返回一个新的KlineDao实例。
// 该函数接收一个MongoDB数据库实例作为参数，并返回一个新的KlineDao实例。
func NewKlineDao(db *mongo.Database) *KlineDao {
	return &KlineDao{
		db: db,
	}
}
