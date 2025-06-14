package tools

import (
	"net"
	"net/http"
)

// GetRemoteClientIp 获取远程客户端的IP地址。
// 该函数优先使用请求头中的X-Real-IP或X-Forwarded-For头字段，
// 如果都不存在，则使用RemoteAddr字段。
// 参数:
//
//	r *http.Request - HTTP请求对象，用于获取客户端信息。
//
// 返回值:
//
//	string - 客户端的IP地址。
func GetRemoteClientIp(r *http.Request) string {
	// 初始化remoteIp为请求的远程地址。
	remoteIp := r.RemoteAddr

	// 检查并使用请求头中的X-Real-IP。
	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		remoteIp = ip
	} else if ip = r.Header.Get("X-Forwarded-For"); ip != "" {
		// 如果X-Real-IP不存在，检查并使用X-Forwarded-For。
		remoteIp = ip
	} else {
		// 如果上述两个头都不存在，尝试从RemoteAddr中提取IP。
		remoteIp, _, _ = net.SplitHostPort(remoteIp)
	}

	// 本地ip处理，将IPv6的本地地址转换为IPv4的本地地址。
	if remoteIp == "::1" {
		remoteIp = "127.0.0.1"
	}

	// 返回最终确定的客户端IP地址。
	return remoteIp
}
