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

// ExchangeRateHandler 货币汇率Handler
type ExchangeRateHandler struct {
	svcCtx *svc.ServiceContext
}

// NewExchangeRateHandler 创建ExchangeRateHandler
func NewExchangeRateHandler(svcCtx *svc.ServiceContext) *ExchangeRateHandler {
	return &ExchangeRateHandler{
		svcCtx: svcCtx, // 添加ServiceContext
	}
}

// UsdRate 获取美元汇率
// 该方法处理获取美元汇率的HTTP请求，并返回相应的汇率信息。
// 参数:
//
//	w http.ResponseWriter: HTTP响应写入器，用于向客户端发送响应。
//	r *http.Request: HTTP请求对象，包含客户端发送的请求信息。
func (h *ExchangeRateHandler) UsdRate(w http.ResponseWriter, r *http.Request) {
	var req types.RateRequest
	// 解析请求路径参数，如果解析失败则返回错误响应。
	if err := httpx.ParsePath(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	// 创建一个新的响应结果对象。
	newResult := common.NewResult()
	// 获取客户端的IP地址。
	req.Ip = tools.GetRemoteClientIp(r)
	// 创建新的ExchangeRateLogic对象以处理业务逻辑。
	l := logic.NewExchangeRateLogic(r.Context(), h.svcCtx)
	// 调用业务逻辑方法获取美元汇率信息，并处理结果。
	resp, err := l.UsdRate(&req)
	result := newResult.Deal(resp.Rate, err)
	// 返回处理结果的HTTP响应。
	httpx.OkJsonCtx(r.Context(), w, result)
}
