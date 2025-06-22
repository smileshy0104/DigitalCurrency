package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
	"job-center/internal/database"
	"job-center/internal/logic"
)

// 配置文件
type Config struct {
	Okx        logic.OkxConfig      // okx配置
	Mongo      database.MongoConfig // mongo配置
	Kafka      database.KafkaConfig // kafka配置
	CacheRedis cache.CacheConf      // redis配置
	UCenterRpc zrpc.RpcClientConf   // ucenter rpc
}
