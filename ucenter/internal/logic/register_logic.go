package logic

import (
	"common/tools"
	"context"
	"errors"
	"grpc-common/ucenter/types/register"
	"time"
	"ucenter/internal/domain"
	"ucenter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// RegisterCacheKey 是验证码缓存的键前缀
const RegisterCacheKey = "REGISTER::"

// RegisterLogic 是处理注册逻辑的结构体
type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	CaptchaDomain *domain.CaptchaDomain
	MemberDomain  *domain.MemberDomain
}

// NewRegisterLogic 创建一个新的 RegisterLogic 实例
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	// 创建一个新的 CaptchaDomain 实例
	return &RegisterLogic{
		ctx:           ctx,
		svcCtx:        svcCtx,
		Logger:        logx.WithContext(ctx),
		CaptchaDomain: domain.NewCaptchaDomain(),
		MemberDomain:  domain.NewMemberDomain(svcCtx.Db),
	}
}

// RegisterByPhone 通过手机号码注册用户
// 参数 in 包含注册所需的信息，如验证码、手机号码、密码等
// 返回注册结果和错误信息（如果有）
func (l *RegisterLogic) RegisterByPhone(in *register.RegReq) (*register.RegRes, error) {
	//1. 先校验人机是否通过
	//isVerify := l.CaptchaDomain.Verify(
	//	in.Captcha.Server,
	//	l.svcCtx.Config.Captcha.Vid,
	//	l.svcCtx.Config.Captcha.Key,
	//	in.Captcha.Token,
	//	2,
	//	in.Ip)
	//if !isVerify {
	//	return nil, errors.New("人机校验不通过")
	//}
	//2.校验验证码
	redisValue := ""
	err := l.svcCtx.Cache.GetCtx(context.Background(), RegisterCacheKey+in.Phone, &redisValue)
	if err != nil {
		return nil, errors.New("验证码获取错误")
	}
	if in.Code != redisValue {
		return nil, errors.New("验证码输入错误")
	}
	//3.验证码通过 进行注册即可 手机号首先验证此手机号是否注册过
	mem, err := l.MemberDomain.FindByPhone(context.Background(), in.Phone)
	if err != nil {
		return nil, errors.New("服务异常，请联系管理员")
	}
	if mem != nil {
		return nil, errors.New("此手机号已经被注册")
	}
	//4. 生成member模型，存入数据库
	err = l.MemberDomain.Register(
		context.Background(),
		in.Phone,
		in.Password,
		in.Username,
		in.Country,
		in.SuperPartner,
		in.Promotion)
	if err != nil {
		return nil, errors.New("注册失败")
	}
	return &register.RegRes{}, nil
}

// SendCode 发送验证码到指定手机号码
// 参数 req 包含手机号码和国家标识
// 返回发送结果和错误信息（如果有）
func (l *RegisterLogic) SendCode(req *register.CodeReq) (*register.NoRes, error) {
	//* 收到手机号和国家标识
	//* 生成验证码
	//* 根据对应的国家和手机号调用对应的短信平台发送验证码
	//* 将验证码存入redis，过期时间5分钟
	//* 返回成功
	code := tools.Rand4Num()
	//假设调用短信平台发送验证码成功
	go func() {
		logx.Info("调用短信平台发送验证码成功")
	}()
	logx.Infof("验证码为：%s \n", code)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := l.svcCtx.Cache.SetWithExpireCtx(ctx, RegisterCacheKey+req.Phone, code, 15*time.Minute)
	if err != nil {
		return nil, errors.New("验证码发送失败")
	}
	return &register.NoRes{}, nil
}
