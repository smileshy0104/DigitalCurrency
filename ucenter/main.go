package main

import (
	"flag"
	"fmt"
	"grpc-common/ucenter/types/asset"
	"grpc-common/ucenter/types/login"
	"grpc-common/ucenter/types/register"
	"ucenter/internal/config"
	"ucenter/internal/server"
	"ucenter/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/conf.yaml", "the config file")

func main() {
	flag.Parse()
	// 加载配置文件
	var c config.Config
	conf.MustLoad(*configFile, &c)
	// 创建并初始化一个新的服务上下文（初始化各个组件）
	ctx := svc.NewServiceContext(c)

	// 创建一个新的 gRPC 服务器，并注册 RegisterServer 和反射服务
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		register.RegisterRegisterServer(grpcServer, server.NewRegisterServer(ctx))
		login.RegisterLoginServer(grpcServer, server.NewLoginServer(ctx))
		asset.RegisterAssetServer(grpcServer, server.NewAssetServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
