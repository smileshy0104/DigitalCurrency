package dao

import (
	"common/db"
	"common/db/gorms"
	"context"
	"gorm.io/gorm"
	"ucenter/internal/model"
)

type MemberWalletDao struct {
	conn *gorms.GormConn
}

func (m *MemberWalletDao) Save(ctx context.Context, mw *model.MemberWallet) error {
	//TODO implement me
	panic("implement me")
}

func (m *MemberWalletDao) FindByIdAndCoinName(ctx context.Context, memId int64, coinName string) (mw *model.MemberWallet, err error) {
	//TODO implement me
	panic("implement me")
}

func (m *MemberWalletDao) UpdateFreeze(ctx context.Context, conn db.DbConn, memberId int64, symbol string, money float64) error {
	//TODO implement me
	panic("implement me")
}

func (m *MemberWalletDao) UpdateWallet(ctx context.Context, conn db.DbConn, id int64, balance float64, frozenBalance float64) error {
	//TODO implement me
	panic("implement me")
}

func (m *MemberWalletDao) FindByMemberId(ctx context.Context, memId int64) ([]*model.MemberWallet, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MemberWalletDao) UpdateAddress(ctx context.Context, wallet *model.MemberWallet) error {
	//TODO implement me
	panic("implement me")
}

func (m *MemberWalletDao) FindAllAddress(ctx context.Context, name string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MemberWalletDao) FindByAddress(ctx context.Context, address string) (*model.MemberWallet, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MemberWalletDao) FindByIdAndCoinId(ctx context.Context, memId int64, coinId int64) (mw *model.MemberWallet, err error) {
	session := m.conn.Session(ctx)
	err = session.Model(&model.MemberWallet{}).
		Where("member_id=? and coin_id=?", memId, coinId).
		Take(&mw).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return
}

func NewMemberWalletDao(db *db.DB) *MemberWalletDao {
	return &MemberWalletDao{
		conn: gorms.New(db.Conn),
	}
}
