package handler

import (
	"common"
	"common/tools"
	"errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ucenter-api/internal/logic"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"
)

// RegisterHandler 处理用户注册相关的请求.
type RegisterHandler struct {
	svcCtx *svc.ServiceContext // 服务上下文，用于访问服务层逻辑
}

// NewRegisterHandler 创建一个新的 RegisterHandler 实例.
func NewRegisterHandler(svcCtx *svc.ServiceContext) *RegisterHandler {
	return &RegisterHandler{
		svcCtx: svcCtx,
	}
}

// Register 处理用户注册请求.
// 该方法解析请求体中的用户信息，并进行注册逻辑处理.
func (h *RegisterHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req types.Request
	// 解析请求体，如果解析失败则返回错误信息.
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}

	newResult := common.NewResult()
	// 验证验证码，如果验证码为空则返回错误信息.
	if req.Captcha == nil {
		httpx.OkJsonCtx(r.Context(), w, newResult.Deal(nil, errors.New("人机校验不通过")))
		return
	}
	//获取一下ip
	req.Ip = tools.GetRemoteClientIp(r)
	// 创建注册逻辑实例，并调用注册逻辑.
	l := logic.NewRegisterLogic(r.Context(), h.svcCtx)
	resp, err := l.Register(&req)
	result := newResult.Deal(resp, err)
	// 返回注册结果.
	httpx.OkJsonCtx(r.Context(), w, result)
}

// SendCode 处理发送验证码请求.
// 该方法解析请求体中的验证码发送信息，并调用发送验证码逻辑.
func (h *RegisterHandler) SendCode(w http.ResponseWriter, r *http.Request) {
	var req types.CodeRequest
	// 解析请求体，如果解析失败则返回错误信息.
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}

	// 创建注册逻辑实例，并调用发送验证码逻辑.
	l := logic.NewRegisterLogic(r.Context(), h.svcCtx)
	resp, err := l.SendCode(&req)
	result := common.NewResult().Deal(resp, err)
	// 返回发送验证码结果.
	httpx.OkJsonCtx(r.Context(), w, result)
}
