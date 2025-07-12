package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// MongoConfig 包含了连接MongoDB所需的配置信息
type MongoConfig struct {
	Url      string // 数据库的URL地址
	Username string // 用户名
	Password string // 密码
	DataBase string // 数据库名称
}

// MongoClient 代表了与MongoDB的连接客户端
type MongoClient struct {
	cli *mongo.Client   // MongoDB客户端
	Db  *mongo.Database // MongoDB数据库实例
}

// ConnectMongo 用于建立与MongoDB的连接
// 参数c包含了连接所需的配置信息
// 返回MongoClient的实例，其中包含了连接客户端和指定的数据库实例
func ConnectMongo(c MongoConfig) *MongoClient {
	// 创建一个带有超时的上下文，以防止连接尝试无限期地挂起
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 设置MongoDB连接的认证信息
	credential := options.Credential{
		Username: c.Username,
		Password: c.Password,
	}

	// 尝试连接到MongoDB实例
	client, err := mongo.Connect(ctx, options.Client().
		ApplyURI(c.Url).
		SetAuth(credential))
	if err != nil {
		panic(err)
	}

	// 检查连接是否成功
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	// 获取指定名称的数据库实例
	database := client.Database(c.DataBase)

	// 返回包含MongoDB客户端和数据库实例的信息
	return &MongoClient{cli: client, Db: database}
}

// Disconnect 用于断开与MongoDB的连接
func (c *MongoClient) Disconnect() {
	// 创建一个带有超时的上下文，以防止断开连接的操作无限期地挂起
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 尝试断开与MongoDB的连接
	err := c.cli.Disconnect(ctx)
	if err != nil {
		log.Println(err)
	}

	// 记录日志，表示Mongo连接已经关闭
	log.Println("关闭Mongo连接..")
}
