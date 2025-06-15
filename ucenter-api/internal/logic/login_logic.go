package logic

import (
	"common/tools"
	"context"
	"github.com/jinzhu/copier"
	"grpc-common/ucenter/types/login"
	"time"

	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// LoginLogic 结构体定义了登录逻辑的上下文和依赖的服务上下文。
// 它继承了 logx.Logger，以便能够使用日志记录功能。
type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewLoginLogic 创建并返回一个新的 LoginLogic 实例。
// ctx: 用于取消请求和传递请求范围的值的上下文。
// svcCtx: 包含服务相关依赖的上下文。
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Login 处理用户登录请求。
// req: 包含用户登录信息的请求对象。
// 返回: 登录响应对象，包含登录结果和可能的错误。
func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRes, err error) {
	// 为登录操作设置一个带有超时的上下文，以防止长时间运行的登录操作阻塞系统。
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	// 初始化登录请求对象，并将传入的登录信息复制到该对象中。
	loginReq := &login.LoginReq{}
	if err := copier.Copy(loginReq, req); err != nil {
		return nil, err
	}

	// 调用登录RPC服务，执行登录操作。
	loginResp, err := l.svcCtx.ULoginRpc.Login(ctx, loginReq)
	if err != nil {
		return nil, err
	}

	// 初始化登录响应对象，并将登录操作的结果复制到该对象中。
	resp = &types.LoginRes{}
	if err := copier.Copy(resp, loginResp); err != nil {
		return nil, err
	}
	return
}

// CheckLogin 检查给定的token是否有效。
// token: 用于验证用户登录状态的令牌字符串。
// 返回: 布尔值表示token是否有效，以及可能的错误。
func (l *LoginLogic) CheckLogin(token string) (bool, error) {
	// 尝试解析token，如果解析失败，则记录错误并返回false。
	_, err := tools.ParseToken(token, l.svcCtx.Config.JWT.AccessSecret)
	if err != nil {
		logx.Error(err)
		return false, nil
	}
	// 如果token解析成功，则返回true，表示登录有效。
	return true, nil
}
