package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc-common/market/types/market"
	"grpc-common/market/types/rate"
	"market/internal/config"
	"market/internal/server"
	"market/internal/svc"
)

var configFile = flag.String("f", "etc/conf.yaml", "the config file")

func main() {
	flag.Parse()
	// 日志的打印格式替换一下
	logx.MustSetup(logx.LogConf{Stat: false, Encoding: "plain"})
	var c config.Config
	conf.MustLoad(*configFile, &c)
	// 初始化服务上下文
	ctx := svc.NewServiceContext(c)
	// 创建rpc服务
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		// 注册货币汇率服务
		rate.RegisterExchangeRateServer(grpcServer, server.NewExchangeRateServer(ctx))
		// 注册市场行情服务
		market.RegisterMarketServer(grpcServer, server.NewMarketServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
