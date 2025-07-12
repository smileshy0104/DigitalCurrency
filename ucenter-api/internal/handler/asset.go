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

// AssetHandler 处理与资产相关的操作。
// 它依赖于服务上下文 svcCtx 来进行操作。
type AssetHandler struct {
	// svcCtx 是服务的上下文，包含了处理资产时需要的服务信息和配置。
	svcCtx *svc.ServiceContext
}

// NewAssetHandler 创建并返回一个新的 AssetHandler 实例。
// 参数 svcCtx 是服务的上下文，对于处理资产是必需的。
// 返回值是新创建的 AssetHandler 实例，通过它来执行与资产相关的操作。
func NewAssetHandler(svcCtx *svc.ServiceContext) *AssetHandler {
	// 使用给定的服务上下文初始化 AssetHandler，并返回该实例。
	return &AssetHandler{
		svcCtx: svcCtx,
	}
}

// FindWalletBySymbol 根据符号查找钱包信息。
// 该方法首先解析请求路径以获取资产请求信息，然后通过逻辑层查询对应符号的钱包信息，
// 并将查询结果以JSON格式返回给客户端。
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

	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

// FindWallet 查询用户钱包信息
func (h *AssetHandler) FindWallet(w http.ResponseWriter, r *http.Request) {
	var req = types.AssetReq{}
	ip := tools.GetRemoteClientIp(r)
	req.Ip = ip
	l := logic.NewAssetLogic(r.Context(), h.svcCtx)
	// 查询用户钱包信息
	resp, err := l.FindWallet(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

// ResetAddress 重置用户钱包地址
func (h *AssetHandler) ResetAddress(w http.ResponseWriter, r *http.Request) {
	var req types.AssetReq
	if err := httpx.ParseForm(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	ip := tools.GetRemoteClientIp(r)
	req.Ip = ip
	l := logic.NewAssetLogic(r.Context(), h.svcCtx)
	// 重置用户钱包地址
	resp, err := l.ResetAddress(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

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
