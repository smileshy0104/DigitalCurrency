package handler

import (
	"common"
	"common/tools"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ucenter-api/internal/logic"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"
)

type AssetHandler struct {
	svcCtx *svc.ServiceContext
}

func NewAssetHandler(svcCtx *svc.ServiceContext) *AssetHandler {
	return &AssetHandler{
		svcCtx: svcCtx,
	}
}

// FindWalletBySymbol 根据符号查找钱包信息。
// 该方法首先解析请求路径以获取资产请求信息，然后通过逻辑层查询对应符号的钱包信息，
// 并将查询结果以JSON格式返回给客户端。
//
// 参数:
//
//	w http.ResponseWriter: HTTP响应写入接口，用于向客户端发送响应。
//	r *http.Request: HTTP请求对象，包含客户端请求信息。
func (h *AssetHandler) FindWalletBySymbol(w http.ResponseWriter, r *http.Request) {
	// 解析请求路径以获取资产请求对象。
	var req types.AssetReq
	if err := httpx.ParsePath(r, &req); err != nil {
		// 如果解析出错，记录错误并返回错误响应。
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}

	// 获取客户端IP地址，并将其添加到请求对象中。
	ip := tools.GetRemoteClientIp(r)
	req.Ip = ip

	// 创建资产逻辑处理对象。
	l := logic.NewAssetLogic(r.Context(), h.svcCtx)

	// 调用逻辑层方法，根据符号查找钱包信息。
	resp, err := l.FindWalletBySymbol(&req)

	// 根据查询结果创建响应对象。
	result := common.NewResult().Deal(resp, err)

	// 将响应结果以JSON格式返回给客户端。
	httpx.OkJsonCtx(r.Context(), w, result)
}

//
//func (h *AssetHandler) FindWallet(w http.ResponseWriter, r *http.Request) {
//	var req = types.AssetReq{}
//	ip := tools.GetRemoteClientIp(r)
//	req.Ip = ip
//	l := logic.NewAssetLogic(r.Context(), h.svcCtx)
//	resp, err := l.FindWallet(&req)
//	result := common.NewResult().Deal(resp, err)
//	httpx.OkJsonCtx(r.Context(), w, result)
//}
//
//func (h *AssetHandler) ResetAddress(w http.ResponseWriter, r *http.Request) {
//	var req types.AssetReq
//	if err := httpx.ParseForm(r, &req); err != nil {
//		httpx.ErrorCtx(r.Context(), w, err)
//		return
//	}
//	ip := tools.GetRemoteClientIp(r)
//	req.Ip = ip
//	l := logic.NewAssetLogic(r.Context(), h.svcCtx)
//	resp, err := l.ResetAddress(&req)
//	result := common.NewResult().Deal(resp, err)
//	httpx.OkJsonCtx(r.Context(), w, result)
//}
//
//func (h *AssetHandler) FindTransaction(w http.ResponseWriter, r *http.Request) {
//	var req types.AssetReq
//	if err := httpx.ParseForm(r, &req); err != nil {
//		httpx.ErrorCtx(r.Context(), w, err)
//		return
//	}
//	ip := tools.GetRemoteClientIp(r)
//	req.Ip = ip
//	l := logic.NewAssetLogic(r.Context(), h.svcCtx)
//	resp, err := l.FindTransaction(&req)
//	result := common.NewResult().Deal(resp, err)
//	httpx.OkJsonCtx(r.Context(), w, result)
//}
