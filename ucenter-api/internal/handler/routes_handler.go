// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.3

package handler

import (
	"ucenter-api/internal/svc"
)

// RegisterHandlers 注册处理函数到指定的路由组中。
// 该函数主要用于初始化与用户注册相关的HTTP处理函数。
// 参数:
//
//	r *Routers - 路由器指针，用于注册处理函数。
//	serverCtx *svc.ServiceContext - 服务上下文指针，包含处理函数所需的上下文信息。
func RegisterHandlers(r *Routers, serverCtx *svc.ServiceContext) {
	// 创建一个新的注册处理实例。
	register := NewRegisterHandler(serverCtx)

	// 创建一个路由组用于用户注册相关的处理函数。
	registerGroup := r.Group()

	// 在路由组中注册用户注册处理函数，处理用户通过电话号码注册的请求。
	registerGroup.Post("/uc/register/phone", register.Register)

	// 在路由组中注册发送验证码处理函数，处理发送移动电话验证码的请求。
	registerGroup.Post("/uc/mobile/code", register.SendCode)

	// 创建一个新的登录处理实例。
	login := NewLoginHandler(serverCtx)
	loginGroup := r.Group()
	// 在路由组中注册用户登录处理函数，处理用户登录的请求。
	loginGroup.Post("/uc/login", login.Login)
	loginGroup.Post("/uc/check/login", login.CheckLogin)
}
