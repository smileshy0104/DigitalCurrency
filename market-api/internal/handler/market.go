package handler

import (
	"common"
	"common/tools"
	"github.com/zeromicro/go-zero/rest/httpx"
	"market-api/internal/logic"
	"market-api/internal/svc"
	"market-api/internal/types"
	"net/http"
)

// MarketHandler 市场Handler
type MarketHandler struct {
	svcCtx *svc.ServiceContext
}

// NewMarketHandler 创建MarketHandler
func NewMarketHandler(svcCtx *svc.ServiceContext) *MarketHandler {
	return &MarketHandler{
		svcCtx: svcCtx, // 添加ServiceContext
	}
}

// SymbolThumbTrend 获取币种行情趋势
func (h *MarketHandler) SymbolThumbTrend(w http.ResponseWriter, r *http.Request) {
	var req = &types.MarketReq{}
	newResult := common.NewResult()
	//获取一下ip
	req.Ip = tools.GetRemoteClientIp(r)
	// 初始化市场模块逻辑层
	l := logic.NewMarketLogic(r.Context(), h.svcCtx)
	// 调用货币趋势模块逻辑层
	resp, err := l.SymbolThumbTrend(req)
	result := newResult.Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

// SymbolThumb 获取币种行情
func (h *MarketHandler) SymbolThumb(w http.ResponseWriter, r *http.Request) {
	var req = &types.MarketReq{}
	newResult := common.NewResult()
	//获取一下ip
	req.Ip = tools.GetRemoteClientIp(r)
	l := logic.NewMarketLogic(r.Context(), h.svcCtx)
	resp, err := l.SymbolThumb(req)
	result := newResult.Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

// SymbolInfo 获取币种信息
func (h *MarketHandler) SymbolInfo(w http.ResponseWriter, r *http.Request) {
	var req types.MarketReq
	if err := httpx.ParseForm(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	newResult := common.NewResult()
	req.Ip = tools.GetRemoteClientIp(r)
	l := logic.NewMarketLogic(r.Context(), h.svcCtx)
	resp, err := l.SymbolInfo(req)
	result := newResult.Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

// CoinInfo 货币信息
func (h *MarketHandler) CoinInfo(w http.ResponseWriter, r *http.Request) {
	var req types.MarketReq
	if err := httpx.ParseForm(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	ip := tools.GetRemoteClientIp(r)
	req.Ip = ip
	l := logic.NewMarketLogic(r.Context(), h.svcCtx)
	resp, err := l.CoinInfo(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

// History
func (h *MarketHandler) History(w http.ResponseWriter, r *http.Request) {
	var req types.MarketReq
	if err := httpx.ParseForm(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	ip := tools.GetRemoteClientIp(r)
	req.Ip = ip
	l := logic.NewMarketLogic(r.Context(), h.svcCtx)
	resp, err := l.History(&req)
	result := common.NewResult().Deal(resp.List, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}
