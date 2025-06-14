package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"grpc-common/ucenter/types/register"
	"time"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// RegisterLogic 是处理用户注册逻辑的结构体。
// 它包含了日志记录器、上下文和服务中心的引用。
type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewRegisterLogic 创建一个新的 RegisterLogic 实例。
// 参数 ctx 是上下文，svcCtx 是服务上下文。
// 返回值是 RegisterLogic 的实例。
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Register 处理用户注册请求。
// 参数 req 是用户的注册请求。
// 返回值 resp 是注册响应，err 是错误信息（如果有）。
// 该函数通过 gRPC 调用远程服务完成注册逻辑。
func (l *RegisterLogic) Register(req *types.Request) (resp *types.Response, err error) {
	// 为防止长时间运行的操作阻塞，设置一个5秒的超时上下文。
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 创建一个注册请求对象，并将传入的请求数据复制到该对象中。
	regReq := &register.RegReq{}
	if err := copier.Copy(regReq, req); err != nil {
		return nil, err
	}
	// 调用 gRPC 服务进行注册。
	_, err = l.svcCtx.URegisterRpc.RegisterByPhone(ctx, regReq)
	if err != nil {
		return nil, err
	}
	return
}

// SendCode 发送验证码。
// 参数 req 是发送验证码的请求，包含电话号码和国家代码。
// 返回值 resp 是发送验证码的响应，err 是错误信息（如果有）。
// 该函数通过 gRPC 调用远程服务发送验证码。
func (l *RegisterLogic) SendCode(req *types.CodeRequest) (resp *types.CodeResponse, err error) {
	// 为防止长时间运行的操作阻塞，设置一个5秒的超时上下文。
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 调用 gRPC 服务发送验证码。
	_, err = l.svcCtx.URegisterRpc.SendCode(ctx, &register.CodeReq{
		Phone:   req.Phone,
		Country: req.Country,
	})
	if err != nil {
		return nil, err
	}
	return
}
