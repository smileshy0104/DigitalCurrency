package svc

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"job-center/internal/config"
	"job-center/internal/database"
)

// ServiceContext 服务上下文
type ServiceContext struct {
	Config      config.Config         // 配置
	MongoClient *database.MongoClient // mongo
	KafkaClient *database.KafkaClient // kafka
	Cache       cache.Cache           // 缓存
}

// NewServiceContext 创建服务上下文
func NewServiceContext(c config.Config) *ServiceContext {
	client := database.NewKafkaClient(c.Kafka)
	client.StartWrite()
	redisCache := cache.New(
		c.CacheRedis,
		nil,
		cache.NewStat("mscoin"),
		nil,
		func(o *cache.Options) {})
	return &ServiceContext{
		Config:      c,
		MongoClient: database.ConnectMongo(c.Mongo),
		KafkaClient: client,
		Cache:       redisCache,
	}
}
