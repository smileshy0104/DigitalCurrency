package dao

import (
	"common/db"
	"common/db/gorms"
	"context"
	"gorm.io/gorm"
	"ucenter/internal/model"
)

// MemberDao 定义了与成员相关的数据访问对象接口
type MemberDao struct {
	conn *gorms.GormConn // 使用 GORM 连接
}

// FindMemberById 根据成员 ID 查询并返回成员信息
// 如果未找到成员，则返回 nil, nil
func (m *MemberDao) FindMemberById(ctx context.Context, memberId int64) (mem *model.Member, err error) {
	session := m.conn.Session(ctx)
	err = session.Model(&model.Member{}).Where("id=?", memberId).Take(&mem).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return
}

// UpdateLoginCount 更新成员的登录次数
// 增加指定 ID 成员的登录次数，增加量为 step
func (m *MemberDao) UpdateLoginCount(ctx context.Context, id int64, step int) error {
	session := m.conn.Session(ctx)
	// login_count = login_count + step
	err := session.Exec("update member set login_count = login_count+? where id=?", step, id).Error
	return err
}

// Save 插入或更新成员信息
// 如果成员已经存在于数据库中，则更新该成员；否则插入新成员
func (m *MemberDao) Save(ctx context.Context, mem *model.Member) error {
	session := m.conn.Session(ctx)
	err := session.Save(mem).Error
	return err
}

// FindByPhone 根据手机号码查询并返回成员信息
// 如果未找到成员，则返回 nil, nil
func (m *MemberDao) FindByPhone(ctx context.Context, phone string) (mem *model.Member, err error) {
	session := m.conn.Session(ctx)
	err = session.Model(&model.Member{}).
		Where("mobile_phone=?", phone).Limit(1).
		Take(&mem).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return mem, err
}

// NewMemberDao 创建并返回一个新的 MemberDao 实例
// 参数 db 是数据库连接对象
func NewMemberDao(db *db.DB) *MemberDao {
	return &MemberDao{
		conn: gorms.New(db.Conn),
	}
}
