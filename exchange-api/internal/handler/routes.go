package handler

import (
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

// Routers 路由管理器，用于管理HTTP请求的路由和中间件
type Routers struct {
	server      *rest.Server      // Go Zero的REST服务器实例
	middlewares []rest.Middleware // 中间件集合
	prefix      string
}

// NewRouters 创建新的Routers实例
// 参数 server: Go Zero的REST服务器实例
// 返回值: Routers的指针
func NewRouters(server *rest.Server, prefix string) *Routers {
	return &Routers{
		server: server,
		prefix: prefix,
	}
}

// Get 添加GET请求的路由
// 参数 path: 请求路径
// 参数 handlerFunc: 处理函数
func (r *Routers) Get(path string, handlerFunc http.HandlerFunc) {
	r.server.AddRoutes(
		rest.WithMiddlewares(
			r.middlewares,
			rest.Route{
				Method:  http.MethodGet,
				Path:    path,
				Handler: handlerFunc,
			},
		),
		rest.WithPrefix(r.prefix),
	)
}

// Post 添加POST请求的路由
// 参数 path: 请求路径
// 参数 handlerFunc: 处理函数
func (r *Routers) Post(path string, handlerFunc http.HandlerFunc) {
	r.server.AddRoutes(
		rest.WithMiddlewares(
			r.middlewares,
			rest.Route{
				Method:  http.MethodPost,
				Path:    path,
				Handler: handlerFunc,
			},
		),
		rest.WithPrefix(r.prefix),
	)
}

func (r *Routers) GetNoPrefix(path string, handlerFunc http.HandlerFunc) {
	r.server.AddRoutes(
		rest.WithMiddlewares(
			r.middlewares,
			rest.Route{
				Method:  http.MethodGet,
				Path:    path,
				Handler: handlerFunc,
			},
		),
	)
}

func (r *Routers) PostNoPrefix(path string, handlerFunc http.HandlerFunc) {
	r.server.AddRoutes(
		rest.WithMiddlewares(
			r.middlewares,
			rest.Route{
				Method:  http.MethodPost,
				Path:    path,
				Handler: handlerFunc,
			},
		),
	)
}

// Group 创建一个新的路由组，用于路由的分组管理
// 返回值: Routers的指针，用于链式调用
func (r *Routers) Group() *Routers {
	return &Routers{
		server: r.server,
		prefix: r.prefix,
	}
}

// Use 设置中间件
// 参数 middlewares: 中间件列表
func (r *Routers) Use(middlewares ...rest.Middleware) {
	r.middlewares = append(r.middlewares, middlewares...)
}
