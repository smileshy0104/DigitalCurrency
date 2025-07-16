package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
	"ucenter/internal/database"
)

// Config 是服务的配置结构体。
// 它包含了 RPC 服务器配置、MySQL 数据库配置、缓存 Redis 配置和验证码配置。
type Config struct {
	zrpc.RpcServerConf                      // RPC 服务器配置，继承自 zrpc.RpcServerConf
	Mysql              MysqlConfig          // MySQL 数据库配置
	CacheRedis         cache.CacheConf      // 缓存 Redis 配置，使用 go-zero 的缓存配置类型
	Captcha            CaptchaConf          // 验证码配置
	JWT                AuthConfig           // JWT 配置
	MarketRpc          zrpc.RpcClientConf   // Market RPC 配置
	ExchangeRpc        zrpc.RpcClientConf   // Exchange Rpc 配置
	Kafka              database.KafkaConfig // Kafka 配置
}

// MysqlConfig 是 MySQL 数据库的配置结构体。
// 它包含了连接数据库所需的数据源信息。
type MysqlConfig struct {
	DataSource string // 数据源地址，用于连接 MySQL 数据库
}

// CaptchaConf 是验证码的配置结构体。
// 它包含了生成验证码所需的 VID 和 Key。
type CaptchaConf struct {
	Vid string // 验证码的 VID
	Key string // 验证码的 Key
}

// AuthConfig 是 JWT 配置结构体。
type AuthConfig struct {
	AccessSecret string
	AccessExpire int64
}
