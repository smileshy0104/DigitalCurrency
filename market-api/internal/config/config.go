package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"market-api/internal/database"
)

// Config 是服务的配置结构体，它包含了 REST API 和用户中心 RPC 的配置信息。
type Config struct {
	Prefix string
	// RestConf 嵌入了 rest 包中的 RestConf 结构体，用于处理 REST API 的配置。
	rest.RestConf
	// MarketRpc 是Market市场的RPC客户端配置
	MarketRpc zrpc.RpcClientConf
	Kafka     database.KafkaConfig
}
