package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"grpc-common/market/mk_client"
	"market-api/internal/config"
	"market-api/internal/database"
	"market-api/internal/processor"
	"market-api/internal/ws"
)

// ServiceContext 是服务的上下文结构体，包含了服务运行所需的各种客户端和配置信息。
// 它是服务运行环境的集合，通过它可以在服务的各个部分之间共享配置和资源。
type ServiceContext struct {
	// Config 是服务的配置信息，包含了服务运行所需的各种参数和设置。
	Config config.Config
	// ExchangeRateRpc 是一个 RPC 客户端，用于调用汇率服务。
	ExchangeRateRpc mk_client.ExchangeRate
	MarketRpc       mk_client.Market
	Processor       processor.Processor // 处理器
}

// NewServiceContext 创建并返回一个新的 ServiceContext 实例。
// 参数 c 是服务的配置信息，用于初始化 ServiceContext 的 Config 字段。
// 这个函数负责初始化 ServiceContext 结构体，并创建用户注册的RPC客户端。
// 返回值是初始化后的 *ServiceContext 实例，即服务上下文的指针。
func NewServiceContext(c config.Config, server *ws.WebsocketServer) *ServiceContext {
	// 初始化 kafka
	kafaCli := database.NewKafkaClient(c.Kafka)
	// 初始化市场模块的 RPC 客户端
	market := mk_client.NewMarket(zrpc.MustNewClient(c.MarketRpc))
	// 创建一个新的DefaultProcessor实例
	defaultProcessor := processor.NewDefaultProcessor(kafaCli)
	defaultProcessor.Init(market)
	// 添加web socket到handler中
	defaultProcessor.AddHandler(processor.NewWebsocketHandler(server))
	// 创建并返回一个新的 ServiceContext 实例，其中 Config 字段设置为传入的配置信息 c，
	// 并根据配置信息中的用户中心RPC地址，创建用户注册的RPC客户端。
	return &ServiceContext{
		Config:          c,
		ExchangeRateRpc: mk_client.NewExchangeRate(zrpc.MustNewClient(c.MarketRpc)),
		MarketRpc:       market,
		Processor:       defaultProcessor,
	}
}
