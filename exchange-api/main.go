package main

import (
	"exchange-api/internal/config"
	"exchange-api/internal/handler"
	"exchange-api/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
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
	server := rest.MustNewServer(
		c.RestConf,
		rest.WithCustomCors(func(header http.Header) {
			header.Set("Access-Control-Allow-Headers", "DNT,X-Mx-ReqToken,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Authorization,token,x-auth-token")
		}, nil, "http://localhost:8080"))
	defer server.Stop()

	// 创建并初始化一个新的服务上下文（初始化各个组件）
	ctx := svc.NewServiceContext(c)
	router := handler.NewRouters(server, c.Prefix)
	handler.RegisterHandlers(router, ctx)

	// 启动服务
	group := service.NewServiceGroup()
	group.Add(server) // 添加服务
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
