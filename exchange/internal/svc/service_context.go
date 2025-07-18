package svc

import (
	"common/db"
	"exchange/internal/config"
	"exchange/internal/consumer"
	"exchange/internal/database"
	"exchange/internal/processor"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
	"grpc-common/market/mk_client"
	"grpc-common/ucenter/uc_client"
)

// ServiceContext 服务上下文结构体
type ServiceContext struct {
	Config      config.Config         // 配置文件对象
	Cache       cache.Cache           // 缓存组件
	Db          *db.DB                // 数据库连接
	MongoClient *database.MongoClient // mongo
	MemberRpc   uc_client.Member      // 用户中心服务
	MarketRpc   mk_client.Market      // 市场行情服务
	AssetRpc    uc_client.Asset       // 资产服务
	KafkaClient *database.KafkaClient // kafka
}

// init 初始化服务上下文
func (sc *ServiceContext) init() {
	// 创建交易引擎工厂
	factory := processor.NewCoinTradeFactory()
	// 初始化交易引擎工厂，传入市场 RPC 客户端、Kafka 客户端和数据库连接
	factory.Init(sc.MarketRpc, sc.KafkaClient, sc.Db)
	// 创建 Kafka 消费者，传入 Kafka 客户端、交易引擎工厂和数据库连接
	kafkaConsumer := consumer.NewKafkaConsumer(sc.KafkaClient, factory, sc.Db)
	// 启动 Kafka 消费者，开始处理订单
	kafkaConsumer.Run()
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
		cache.NewStat("market"),
		nil,
		func(o *cache.Options) {})
	// 初始化MySQL数据库连接。
	mysql := database.ConnMysql(c.Mysql.DataSource)
	// 返回新的服务上下文对象，包含配置对象和初始化后的缓存组件。
	s := &ServiceContext{
		Config:      c,
		Cache:       redisCache,
		MongoClient: database.ConnectMongo(c.Mongo),
		Db:          mysql,
		MemberRpc:   uc_client.NewMember(zrpc.MustNewClient(c.UCenterRpc)),
		MarketRpc:   mk_client.NewMarket(zrpc.MustNewClient(c.MarketRpc)),
		AssetRpc:    uc_client.NewAsset(zrpc.MustNewClient(c.UCenterRpc)),
		KafkaClient: database.NewKafkaClient(c.Kafka),
	}
	s.init()
	return s
}
