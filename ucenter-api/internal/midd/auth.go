package midd

import (
	"common"
	"common/tools"
	"context"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// Auth 是一个中间件生成函数，用于验证请求的认证信息。
// 它接受一个 secret 参数，用于解析和验证 token。
// 返回一个中间件函数，该中间件函数可以包装另一个处理函数，
// 以确保只有通过认证的请求才能被处理。
func Auth(secret string) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		// 返回一个处理函数，它将接收一个 ResponseWriter 和一个指向 Request 的指针。
		return func(w http.ResponseWriter, r *http.Request) {
			// 创建一个新的结果对象，用于处理认证失败的情况。
			result := common.NewResult()
			// 默认情况下设置认证失败的结果。
			result.Fail(4000, "no login")

			// 从请求头中获取 token。
			token := r.Header.Get("x-auth-token")
			// 如果 token 为空，写入认证失败的响应并返回。
			if token == "" {
				httpx.WriteJson(w, 200, result)
				return
			}

			// 解析 token，获取用户 ID。
			userId, err := tools.ParseToken(token, secret)
			// 如果解析失败，写入认证失败的响应并返回。
			if err != nil {
				httpx.WriteJson(w, 200, result)
				return
			}

			// 获取请求的上下文，并将用户 ID 存入上下文中。
			ctx := r.Context()
			ctx = context.WithValue(ctx, "userId", userId)
			// 使用更新后的上下文创建一个新的请求对象。
			r = r.WithContext(ctx)

			// 调用下一个处理函数，传递当前的 ResponseWriter 和更新后的请求对象。
			next(w, r)
		}
	}
}
