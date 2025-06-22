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

func (h *MarketHandler) SymbolThumbTrend(w http.ResponseWriter, r *http.Request) {
	var req = &types.MarketReq{}
	newResult := common.NewResult()
	//获取一下ip
	req.Ip = tools.GetRemoteClientIp(r)
	l := logic.NewMarketLogic(r.Context(), h.svcCtx)
	resp, err := l.SymbolThumbTrend(req)
	result := newResult.Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

//func (h *MarketHandler) SymbolThumb(w http.ResponseWriter, r *http.Request) {
//	var req = &types.MarketReq{}
//	newResult := common.NewResult()
//	//获取一下ip
//	req.Ip = tools.GetRemoteClientIp(r)
//	l := logic.NewMarketLogic(r.Context(), h.svcCtx)
//	resp, err := l.SymbolThumb(req)
//	result := newResult.Deal(resp, err)
//	httpx.OkJsonCtx(r.Context(), w, result)
//}
