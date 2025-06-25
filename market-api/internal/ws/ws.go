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
	path   string
	server *socketio.Server
}

// Start 启动WebSocket服务器
func (ws *WebsocketServer) Start() {
	ws.server.Serve()
}

// Stop 停止WebSocket服务器
func (ws *WebsocketServer) Stop() {
	ws.server.Close()
}

// allowOriginFunc 是一个用于跨域请求验证的函数，这里允许所有跨域请求
var allowOriginFunc = func(r *http.Request) bool {
	return true
}

// NewWebsocketServer 创建并返回一个新的WebSocket服务器实例
func NewWebsocketServer(path string) *WebsocketServer {
	//解决跨域
	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: allowOriginFunc,
			},
			&websocket.Transport{
				CheckOrigin: allowOriginFunc,
			},
		},
	})
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		logx.Info("connected:", s.ID())
		s.Join(ROOM)
		return nil
	})
	return &WebsocketServer{
		path:   path,
		server: server,
	}
}

// BroadcastToNamespace 在指定的命名空间和房间内广播事件
func (w *WebsocketServer) BroadcastToNamespace(path string, event string, data any) {
	go func() {
		w.server.BroadcastToRoom(path, ROOM, event, data)
	}()
}

// ServerHandler 返回一个HTTP处理函数，用于处理WebSocket连接或其他请求
func (ws *WebsocketServer) ServerHandler(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//所有的请求都要通过这里
		path := r.URL.Path
		logx.Info("==========================", path)
		if strings.HasPrefix(path, ws.path) {
			//进行我们的处理
			ws.server.ServeHTTP(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
