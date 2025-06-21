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
	Cache       cache.Cache           // 缓存
}

// NewServiceContext 创建服务上下文
func NewServiceContext(c config.Config) *ServiceContext {
	redisCache := cache.New(
		c.CacheRedis,
		nil,
		cache.NewStat("mscoin"),
		nil,
		func(o *cache.Options) {})
	return &ServiceContext{
		Config:      c,
		MongoClient: database.ConnectMongo(c.Mongo),
		Cache:       redisCache,
	}
}
