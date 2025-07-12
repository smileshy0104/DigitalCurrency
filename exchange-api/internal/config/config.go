package config

import (
	"exchange-api/internal/database"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

// Config 是服务的配置结构体，它包含了 REST API 和用户中心 RPC 的配置信息。
type Config struct {
	Prefix string
	// RestConf 嵌入了 rest 包中的 RestConf 结构体，用于处理 REST API 的配置。
	rest.RestConf
	// ExchangeRpc 是Exchange的RPC客户端配置
	ExchangeRpc zrpc.RpcClientConf
	Kafka       database.KafkaConfig
	// JWT 是用于处理 JWT 的配置，包括访问令牌 secret 和过期时间。
	JWT AuthConfig
}

// AuthConfig 是用于处理 JWT 的配置结构体，它包含访问令牌 secret 和过期时间。
type AuthConfig struct {
	AccessSecret string
	AccessExpire int64
}
