package svc

import (
	"common/db"
	"exchange/internal/config"
	"exchange/internal/database"
	"github.com/zeromicro/go-zero/core/stores/cache"
)

// ServiceContext 服务上下文结构体
type ServiceContext struct {
	Config      config.Config         // 配置文件对象
	Cache       cache.Cache           // 缓存组件
	MongoClient *database.MongoClient // mongo
	Db          *db.DB                // 数据库连接
}

// NewServiceContext 创建并初始化一个新的服务上下文。
// 该函数接收一个配置对象作为参数，并基于该配置对象初始化服务上下文中的各个组件。
// 主要负责初始化缓存组件，使用redis作为缓存存储。
func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化Redis缓存组件。
	// 这里使用了cache.New函数来创建一个新的缓存对象。
	// 参数分别为：配置对象中的缓存相关配置、nil（表示没有使用额外的中间件）、
	// 一个统计对象用于监控缓存的性能、nil（表示没有使用额外的插件），
	// 以及一个初始化缓存选项的函数，这里为空函数，表示使用默认配置。
	redisCache := cache.New(
		c.CacheRedis,
		nil,
		cache.NewStat("mscoin"),
		nil,
		func(o *cache.Options) {})
	// 初始化MySQL数据库连接。
	mysql := database.ConnMysql(c.Mysql.DataSource)
	// 返回新的服务上下文对象，包含配置对象和初始化后的缓存组件。
	return &ServiceContext{
		Config:      c,
		Cache:       redisCache,
		MongoClient: database.ConnectMongo(c.Mongo),
		Db:          mysql,
	}
}
