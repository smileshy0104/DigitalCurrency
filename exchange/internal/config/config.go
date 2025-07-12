package config

import (
	"exchange/internal/database"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

// Config 是服务的配置结构体。
// 它包含了 RPC 服务器配置、MySQL 数据库配置、缓存 Redis 配置和验证码配置。
type Config struct {
	zrpc.RpcServerConf                      // RPC 服务器配置，继承自 zrpc.RpcServerConf
	Mysql              database.MysqlConfig // MySQL 数据库配置
	CacheRedis         cache.CacheConf      // 缓存 Redis 配置，使用 go-zero 的缓存配置类型
	Mongo              database.MongoConfig // mongo配置
}
