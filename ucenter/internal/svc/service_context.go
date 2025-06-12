package svc

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"ucenter/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Cache  cache.Cache
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

	// 返回新的服务上下文对象，包含配置对象和初始化后的缓存组件。
	return &ServiceContext{
		Config: c,
		Cache:  redisCache,
	}
}
