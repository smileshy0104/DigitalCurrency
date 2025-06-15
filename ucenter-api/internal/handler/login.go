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

// LoginHandler 处理登录相关的HTTP请求。
type LoginHandler struct {
	svcCtx *svc.ServiceContext
}

// NewLoginHandler 创建一个新的LoginHandler实例。
func NewLoginHandler(svcCtx *svc.ServiceContext) *LoginHandler {
	return &LoginHandler{
		svcCtx: svcCtx,
	}
}

// Login 处理用户登录请求。
// 该方法解析请求体中的登录信息，验证用户身份，并返回登录结果。
func (h *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req types.LoginReq
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	newResult := common.NewResult()
	if req.Captcha == nil {
		httpx.OkJsonCtx(r.Context(), w, newResult.Deal(nil, errors.New("人机校验不通过")))
		return
	}
	//获取一下ip
	req.Ip = tools.GetRemoteClientIp(r)
	l := logic.NewLoginLogic(r.Context(), h.svcCtx)
	resp, err := l.Login(&req)
	result := newResult.Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

// CheckLogin 检查用户是否已登录。
// 该方法通过HTTP头中的令牌验证用户登录状态，并返回验证结果。
func (h *LoginHandler) CheckLogin(w http.ResponseWriter, r *http.Request) {
	newResult := common.NewResult()
	token := r.Header.Get("x-auth-token")
	l := logic.NewLoginLogic(r.Context(), h.svcCtx)
	//data : true or false
	resp, err := l.CheckLogin(token)
	result := newResult.Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}
