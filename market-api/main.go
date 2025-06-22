package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/chain"
	"market-api/internal/config"
	"market-api/internal/handler"
	"market-api/internal/svc"
	"market-api/internal/ws"
	"net/http"
)

var configFile = flag.String("f", "etc/conf.yaml", "the config file")

func main() {
	flag.Parse()
	// 读取配置文件
	var c config.Config
	conf.MustLoad(*configFile, &c)
	//防止打印过多的日志
	logx.MustSetup(logx.LogConf{Encoding: "plain", Stat: false})
	wsServer := ws.NewWebsocketServer("/socket.io")
	//server := rest.MustNewServer(c.RestConf)
	server := rest.MustNewServer(
		c.RestConf,
		rest.WithChain(chain.New(wsServer.ServerHandler)),
		rest.WithCustomCors(func(header http.Header) {
			header.Set("Access-Control-Allow-Headers", "DNT,X-Mx-ReqToken,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Authorization,token,x-auth-token")
		}, nil, "*"))
	defer server.Stop()
	// 创建并初始化一个新的服务上下文（初始化各个组件）
	ctx := svc.NewServiceContext(c, wsServer)
	// 注册路由
	routers := handler.NewRouters(server, c.Prefix)
	handler.RegisterHandlers(routers, ctx)

	group := service.NewServiceGroup()
	group.Add(server)
	group.Add(wsServer)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
