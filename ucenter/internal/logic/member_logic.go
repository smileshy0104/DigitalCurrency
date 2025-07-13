package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"grpc-common/ucenter/types/member"
	"ucenter/internal/domain"
	"ucenter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberLogic 定义了处理成员相关逻辑的结构体。
// 它包含了处理逻辑所需要的服务上下文、日志记录器和成员领域对象。
type MemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	memberDomain *domain.MemberDomain
}

// FindMemberById 根据成员ID查找成员信息。
// 参数 req 包含了需要查找的成员ID。
// 返回值是找到的成员信息（MemberInfo）和一个错误对象（如果有的话）。
func (l *MemberLogic) FindMemberById(req *member.MemberReq) (*member.MemberInfo, error) {
	// 调用成员领域对象的FindMemberById方法查找成员。
	mem, err := l.memberDomain.FindMemberById(l.ctx, req.MemberId)
	if err != nil {
		// 如果发生错误，返回nil和错误对象。
		return nil, err
	}
	// 创建一个成员信息响应对象，并将找到的成员数据复制到响应对象中。
	resp := &member.MemberInfo{}
	copier.Copy(resp, mem)
	// 返回成员信息响应对象和nil错误。
	return resp, nil
}

// NewMemberLogic 创建并返回一个新的MemberLogic实例。
// 参数 ctx 是上下文对象，用于取消请求和传递请求级值。
// 参数 svcCtx 包含了服务的全局配置和资源。
func NewMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLogic {
	// 返回一个新的MemberLogic实例，初始化了上下文、服务上下文、日志记录器和成员领域对象。
	return &MemberLogic{
		ctx:          ctx,
		svcCtx:       svcCtx,
		Logger:       logx.WithContext(ctx),
		memberDomain: domain.NewMemberDomain(svcCtx.Db),
	}
}
