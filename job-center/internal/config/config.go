package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
	"job-center/internal/database"
	"job-center/internal/logic"
)

type Config struct {
	Okx        logic.OkxConfig
	Mongo      database.MongoConfig
	CacheRedis cache.CacheConf
	UCenterRpc zrpc.RpcClientConf
}
