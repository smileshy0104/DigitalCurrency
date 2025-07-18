package handler

import (
	"common"
	"common/tools"
	"exchange-api/internal/logic"
	"exchange-api/internal/svc"
	"exchange-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// OrderHandler 订单Handler
type OrderHandler struct {
	svcCtx *svc.ServiceContext
}

// NewOrderHandler 创建OrderHandler
func NewOrderHandler(svcCtx *svc.ServiceContext) *OrderHandler {
	return &OrderHandler{
		svcCtx: svcCtx, // 添加ServiceContext
	}
}

// History 历史委托订单 所有的订单
func (h *OrderHandler) History(w http.ResponseWriter, r *http.Request) {
	// 解析请求路径以获取交易请求对象。
	var req types.ExchangeReq
	if err := httpx.ParseForm(r, &req); err != nil {
		// 如果解析出错，记录错误并返回错误响应。
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}

	// 获取客户端IP地址，并将其添加到请求对象中。
	ip := tools.GetRemoteClientIp(r)
	req.Ip = ip

	// 创建资产逻辑处理对象。
	l := logic.NewOrderLogic(r.Context(), h.svcCtx)
	// 调用逻辑层方法，获取订单历史。
	resp, err := l.History(&req)

	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

// Current 当前委托订单 状态 正在交易的状态
func (h *OrderHandler) Current(w http.ResponseWriter, r *http.Request) {
	// 解析请求路径以获取交易请求对象。
	var req types.ExchangeReq
	if err := httpx.ParseForm(r, &req); err != nil {
		// 如果解析出错，记录错误并返回错误响应。
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}

	// 获取客户端IP地址，并将其添加到请求对象中。
	ip := tools.GetRemoteClientIp(r)
	req.Ip = ip

	// 创建资产逻辑处理对象。
	l := logic.NewOrderLogic(r.Context(), h.svcCtx)
	// 调用逻辑层方法，获取当前委托订单 状态 正在交易的状态
	resp, err := l.Current(&req)

	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

// Add 添加委托订单(添加订单)
func (h *OrderHandler) Add(w http.ResponseWriter, r *http.Request) {
	// 解析请求路径以获取交易请求对象。
	var req types.ExchangeReq
	if err := httpx.ParseForm(r, &req); err != nil {
		// 如果解析出错，记录错误并返回错误响应。
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}

	// 获取客户端IP地址，并将其添加到请求对象中。
	ip := tools.GetRemoteClientIp(r)
	req.Ip = ip

	// 创建资产逻辑处理对象。
	l := logic.NewOrderLogic(r.Context(), h.svcCtx)
	// 调用逻辑层方法，获取当前委托订单 状态 正在交易的状态
	resp, err := l.AddOrder(&req)

	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}
