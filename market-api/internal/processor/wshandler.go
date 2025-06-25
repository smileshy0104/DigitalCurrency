package processor

import (
	"common/tools"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/market/types/market"
	"market-api/internal/model"
	"market-api/internal/ws"
)

// WebsocketHandler 处理websocket消息
type WebsocketHandler struct {
	wsServer *ws.WebsocketServer
}

// HandleTrade 处理交易买盘消息
func (w *WebsocketHandler) HandleTradePlate(symbol string, plate *model.TradePlateResult) {
	bytes, _ := json.Marshal(plate)
	logx.Info("====买卖盘通知:", symbol, plate.Direction, ":", fmt.Sprintf("%d", len(plate.Items)))
	w.wsServer.BroadcastToNamespace("/", "/topic/market/trade-plate/"+symbol, string(bytes))
}

// HandleTrade 处理交易消息
func (w *WebsocketHandler) HandleTrade(symbol string, data []byte) {
	//订单交易完成后 进入这里进行处理 订单就称为K线的一部分 数据量小 无法维持K线 K线来源 okx平台来
	//TODO implement me
	panic("implement me")
}

// HandleKLine 处理K线数据并广播到订阅者。
// 该函数接收一个符号字符串，一个K线数据对象，以及一个货币概览的映射表。
// 它更新或初始化指定符号的货币概览，并将其广播到所有订阅者。
func (w *WebsocketHandler) HandleKLine(symbol string, kline *model.Kline, thumbMap map[string]*market.CoinThumb) {
	// 开始处理K线数据时的日志记录。
	logx.Info("================WebsocketHandler Start=======================")
	logx.Info("================成功调用Kafka中的数据：：", tools.ToTimeString(kline.Time))
	logx.Info("symbol:", symbol)

	// 从映射表中获取指定符号的货币概览。
	thumb := thumbMap[symbol]
	// 如果货币概览不存在，则初始化。
	if thumb == nil {
		thumb = kline.InitCoinThumb(symbol)
	}
	// 将K线数据转换为货币概览对象。
	coinThumb := kline.ToCoinThumb(symbol, thumb)
	// 创建一个空的货币概览对象以存储转换后的数据。
	result := &model.CoinThumb{}
	// 将转换后的数据复制到结果对象中。
	copier.Copy(result, coinThumb)
	// 将结果对象序列化为JSON格式。
	marshal, _ := json.Marshal(result)
	// 广播货币概览数据到所有订阅者。
	w.wsServer.BroadcastToNamespace("/", "/topic/market/thumb", string(marshal))

	// 将K线数据序列化为JSON格式。
	bytes, _ := json.Marshal(kline)
	// 广播K线数据到所有订阅者。
	w.wsServer.BroadcastToNamespace("/", "/topic/market/kline/"+symbol, string(bytes))

	// 结束处理K线数据时的日志记录。
	logx.Info("================WebsocketHandler End=======================")
}

// NewWebsocketHandler 初始化websocket处理器
func NewWebsocketHandler(wsServer *ws.WebsocketServer) *WebsocketHandler {
	return &WebsocketHandler{
		wsServer: wsServer,
	}
}
