package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
	"job-center/internal/logic"
)

type Config struct {
	Okx        logic.OkxConfig
	CacheRedis cache.CacheConf
	UCenterRpc zrpc.RpcClientConf
}
