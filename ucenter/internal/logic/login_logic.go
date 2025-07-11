package logic

import (
	"common/tools"
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"grpc-common/ucenter/types/login"
	"time"
	"ucenter/internal/domain"
	"ucenter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// LoginLogic 是处理登录逻辑的结构体，包含上下文、服务上下文、日志记录器以及验证码和会员相关的领域逻辑。
type LoginLogic struct {
	ctx           context.Context       // 上下文，用于控制请求生命周期和传递元数据。
	svcCtx        *svc.ServiceContext   // 服务上下文，包含服务运行时所需的依赖和配置。
	logx.Logger                         // 日志记录器，用于记录日志信息。
	CaptchaDomain *domain.CaptchaDomain // 验证码领域逻辑，处理与验证码相关的业务逻辑。
	MemberDomain  *domain.MemberDomain  // 会员领域逻辑，处理与会员相关的业务逻辑。
}

// NewLoginLogic 创建并初始化一个 LoginLogic 实例。
// 参数:
// - ctx: 上下文，用于控制请求生命周期和传递元数据。
// - svcCtx: 服务上下文，包含服务运行时所需的依赖和配置。
// 返回值:
// - *LoginLogic: 返回一个初始化完成的 LoginLogic 实例。
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	// 初始化 LoginLogic 实例，并为其字段赋值。
	return &LoginLogic{
		ctx:           ctx,                               // 设置上下文。
		svcCtx:        svcCtx,                            // 设置服务上下文。
		Logger:        logx.WithContext(ctx),             // 初始化日志记录器，并绑定上下文。
		CaptchaDomain: domain.NewCaptchaDomain(svcCtx),   // 初始化验证码领域逻辑。
		MemberDomain:  domain.NewMemberDomain(svcCtx.Db), // 初始化会员领域逻辑，并绑定数据库连接。
	}
}

// Login 登录操作
func (l *LoginLogic) Login(in *login.LoginReq) (*login.LoginRes, error) {
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
	//2. 校验密码
	member, err := l.MemberDomain.FindByPhone(context.Background(), in.GetUsername())
	if err != nil {
		logx.Error(err)
		return nil, errors.New("登录失败")
	}
	if member == nil {
		return nil, errors.New("此用户未注册")
	}
	password := member.Password
	salt := member.Salt
	// tools.Verify验证原始密码是否与已加密的密码匹配
	verify := tools.Verify(in.Password, salt, password, nil)
	if !verify {
		return nil, errors.New("密码不正确")
	}
	//3. 登录成功，生成token，提供给前端，前端调用传递token，我们进行token认证即可
	key := l.svcCtx.Config.JWT.AccessSecret    // 密钥
	expire := l.svcCtx.Config.JWT.AccessExpire // 过期时间
	// getJwtToken 生成JWT令牌字符串token
	token, err := l.getJwtToken(key, time.Now().Unix(), expire, member.Id)
	if err != nil {
		return nil, errors.New("token生成错误")
	}
	// 登录次数
	loginCount := member.LoginCount + 1
	go func() {
		// 更新登录次数
		l.MemberDomain.UpdateLoginCount(context.Background(), member.Id, 1)
	}()
	// 返回用户登录结果
	return &login.LoginRes{
		Token:         token,
		Id:            member.Id,
		Username:      member.Username,
		MemberLevel:   member.MemberLevelStr(),
		MemberRate:    member.MemberRate(),
		RealName:      member.RealName,
		Country:       member.Country,
		Avatar:        member.Avatar,
		PromotionCode: member.PromotionCode,
		SuperPartner:  member.SuperPartner,
		LoginCount:    int32(loginCount),
	}, nil
}

// getJwtToken 生成JWT令牌字符串
// 参数:
//
//	secretKey: 用于签名令牌的密钥字符串
//	iat: 令牌的签发时间(Unix时间戳)
//	seconds: 令牌的有效期时长(秒)
//	userId: 需要包含在令牌中的用户ID
//
// 返回值:
//
//	string: 签名后的完整JWT令牌字符串
//	error: 签名过程中发生的错误(如密钥无效)
func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	// 构建JWT声明(claims)信息
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds // 设置过期时间
	claims["iat"] = iat           // 设置签发时间
	claims["userId"] = userId     // 嵌入用户标识

	// 创建HS256签名方法的令牌对象
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	// 使用密钥进行签名并返回令牌字符串
	return token.SignedString([]byte(secretKey))
}
