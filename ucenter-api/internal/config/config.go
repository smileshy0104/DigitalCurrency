package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

// Config 是服务的配置结构体，它包含了 REST API 和用户中心 RPC 的配置信息。
type Config struct {
	// RestConf 嵌入了 rest 包中的 RestConf 结构体，用于处理 REST API 的配置。
	rest.RestConf
	// UCenterRpc 是用户中心服务的 RPC 配置，通过它来配置与用户中心服务的 RPC 通信。
	UCenterRpc zrpc.RpcClientConf
}
