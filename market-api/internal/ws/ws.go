package ws

import (
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"strings"
)

// ROOM 定义了市场数据广播的房间名称
const ROOM = "market"

// WebsocketServer 表示WebSocket服务器的结构体
type WebsocketServer struct {
	path   string           // 过滤路径进行websocket
	server *socketio.Server // socketio服务
}

// Start 启动WebSocket服务器
func (ws *WebsocketServer) Start() {
	ws.server.Serve() // 启动服务器
}

// Stop 停止WebSocket服务器
func (ws *WebsocketServer) Stop() {
	ws.server.Close() // 停止服务器
}

// allowOriginFunc 是一个用于跨域请求验证的函数，这里允许所有跨域请求
var allowOriginFunc = func(r *http.Request) bool {
	return true
}

// NewWebsocketServer 创建一个新的WebsocketServer实例。
// 参数 path: 服务的路径。
// 返回值: 返回一个指向WebsocketServer的指针。
func NewWebsocketServer(path string) *WebsocketServer {
	// 解决跨域
	server := socketio.NewServer(&engineio.Options{
		// 解决跨域
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: allowOriginFunc,
			},
			&websocket.Transport{
				CheckOrigin: allowOriginFunc,
			},
		},
	})
	// 当客户端连接到服务器时的处理逻辑。
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		logx.Info("connected:", s.ID())
		s.Join(ROOM)
		return nil
	})
	// 返回WebsocketServer实例。
	return &WebsocketServer{
		path:   path,   // 路径
		server: server, // 服务器实例
	}
}

// BroadcastToNamespace 广播消息到指定的命名空间。
//
// path: 指定的命名空间路径。
// event: 事件名称，用于标识消息的类型。
// data: 任意类型的数据，将被广播到命名空间内的所有客户端。
//
// 该函数使用goroutine异步执行广播操作，以提高响应性能和避免阻塞当前协程。
// 通过委托给server的BroadcastToRoom方法来实现实际的广播逻辑。
func (w *WebsocketServer) BroadcastToNamespace(path string, event string, data any) {
	go func() {
		// 创建一个goroutine，用于异步执行广播操作
		w.server.BroadcastToRoom(path, ROOM, event, data)
	}()
}

// ServerHandler 是一个中间件处理函数，用于拦截和处理传入的HTTP请求。
// 它检查请求的路径是否以特定前缀开头，以便决定是交由WebSocket服务器处理，还是继续执行链中的下一个处理程序。
// 参数 next 是链中的下一个 http.Handler，允许在不处理当前请求时将其传递到下一个处理程序。
// 返回值是一个 http.Handler，可以透明地处理所有传入的请求，根据请求路径决定处理方式。
func (ws *WebsocketServer) ServerHandler(next http.Handler) http.Handler {

	// 返回一个匿名的 http.HandlerFunc，它将处理每个传入的请求。
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 所有的请求都要通过这里
		// 获取请求的路径，用于后续的判断。
		path := r.URL.Path
		// 记录日志，用于调试和监控。
		logx.Info("============Web Socket==============", path)
		// 检查路径是否以特定前缀开头，如果是，则进行WebSocket处理。
		if strings.HasPrefix(path, ws.path) {
			// 进行我们的处理
			// 交由WebSocket服务器处理该请求。
			ws.server.ServeHTTP(w, r)
		} else {
			// 如果路径不匹配，继续执行链中的下一个处理程序。
			next.ServeHTTP(w, r)
		}
	})
}
