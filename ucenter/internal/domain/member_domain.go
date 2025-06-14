package domain

import (
	"common/db"
	"common/tools"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"ucenter/internal/dao"
	"ucenter/internal/model"
	"ucenter/internal/repo"
)

// MemberDomain 定义了成员领域逻辑的结构体
type MemberDomain struct {
	memberRepo repo.MemberRepo
}

// FindByPhone 根据电话号码查找成员
// ctx: 上下文，用于传递请求范围的数据、取消信号等
// phone: 要查找的成员的电话号码
// 返回找到的成员信息和错误信息（如果有的话）
func (d *MemberDomain) FindByPhone(ctx context.Context, phone string) (*model.Member, error) {
	//涉及到数据库的查询
	mem, err := d.memberRepo.FindByPhone(ctx, phone)
	if err != nil {
		logx.Error(err)
		return nil, errors.New("数据库异常")
	}
	return mem, nil
}

// Register 函数用于注册新会员。
// 该函数接收注册信息，包括电话、密码、用户名、国家、合作伙伴和推广代码，并执行注册流程。
// 参数:
//
//	ctx - 上下文，用于传递请求范围的信息。
//	phone - 用户的电话号码，作为会员的唯一标识。
//	password - 用户设置的密码，将进行加密处理后存储。
//	username - 用户名。
//	country - 用户所在的国家。
//	partner - 合作伙伴代码，用于跟踪用户来源。
//	promotion - 推广代码，用于营销活动跟踪。
//
// 返回值:
//
//	如果注册过程中发生错误，返回相应的错误。
func (d *MemberDomain) Register(ctx context.Context, phone string, password string,
	username string, country string, partner string, promotion string) error {
	// 创建一个新的会员实例。
	mem := model.NewMember()

	// 对密码进行MD5加密并加盐，以增强安全性。
	// 注意：MD5加密不安全（可通过彩虹表进行破解）。
	salt, pwd := tools.Encode(password, nil)

	// 使用反射为会员实例的默认字段填充默认值。
	// 这是因为会员表中有许多字段不能为空，需要在注册时赋予默认值。
	_ = tools.Default(mem)

	// 设置会员的基本信息。
	mem.Username = username
	mem.Country = country
	mem.Password = pwd
	mem.MobilePhone = phone
	mem.FillSuperPartner(partner)
	mem.PromotionCode = promotion
	mem.MemberLevel = model.GENERAL
	mem.Salt = salt

	// 设置默认头像。
	mem.Avatar = "https://mszlu.oss-cn-beijing.aliyuncs.com/mscoin/defaultavatar.png"

	// 将会员信息保存到数据库中。
	err := d.memberRepo.Save(ctx, mem)
	if err != nil {
		// 记录错误日志并返回一个表示数据库异常的错误。
		logx.Error(err)
		return errors.New("数据库异常")
	}

	// 注册成功，返回nil表示没有发生错误。
	return nil
}

// UpdateLoginCount 更新成员的登录次数
// ctx: 上下文，用于传递请求范围的数据、取消信号等
// id: 成员的ID
// step: 登录次数的增加量
func (d *MemberDomain) UpdateLoginCount(ctx context.Context, id int64, step int) {
	err := d.memberRepo.UpdateLoginCount(ctx, id, step)
	if err != nil {
		logx.Error(err)
	}
}

// FindMemberById 根据ID查找成员
// ctx: 上下文，用于传递请求范围的数据、取消信号等
// memberId: 要查找的成员的ID
// 返回找到的成员信息和错误信息（如果有的话）
func (d *MemberDomain) FindMemberById(ctx context.Context, memberId int64) (*model.Member, error) {
	id, err := d.memberRepo.FindMemberById(ctx, memberId)
	if err == nil && id == nil {
		return nil, errors.New("用户不存在")
	}
	return id, err
}

// NewMemberDomain 创建新的MemberDomain实例
// db: 数据库连接实例
// 返回新的MemberDomain实例
func NewMemberDomain(db *db.DB) *MemberDomain {
	return &MemberDomain{
		memberRepo: dao.NewMemberDao(db),
	}
}
