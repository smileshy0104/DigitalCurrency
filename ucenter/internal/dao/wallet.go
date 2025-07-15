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

// Save 保存会员钱包
func (m *MemberWalletDao) Save(ctx context.Context, mw *model.MemberWallet) error {
	session := m.conn.Session(ctx)
	err := session.Save(&mw).Error
	return err
}

// FindByIdAndCoinName 通过会员id和币种名称查询会员钱包
func (m *MemberWalletDao) FindByIdAndCoinName(ctx context.Context, memId int64, coinName string) (mw *model.MemberWallet, err error) {
	session := m.conn.Session(ctx)
	err = session.Model(&model.MemberWallet{}).
		Where("member_id=? and coin_name=?", memId, coinName).
		Take(&mw).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return
}

// UpdateFreeze 更新会员钱包中的冻结金额
// 该函数通过减少余额并增加冻结余额来更新会员的钱包信息
// 参数:
//
//	ctx context.Context: 上下文对象，用于传递请求范围的信息
//	conn db.DbConn: 数据库连接对象，用于执行数据库操作
//	memberId int64: 会员ID，用于标识钱包的拥有者
//	symbol string: 加密货币符号，用于指定钱包中的特定货币
//	money float64: 金额，表示需要冻结的资金量
//
// 返回值:
//
//	error: 错误对象，如果执行过程中发生错误，则返回该错误
func (m *MemberWalletDao) UpdateFreeze(ctx context.Context, conn db.DbConn, memberId int64, symbol string, money float64) error {
	// 将连接对象转换为GormConn类型
	con := conn.(*gorms.GormConn)
	// 获取上下文中的数据库会话
	session := con.Tx(ctx)
	// 准备SQL语句，用于更新钱包的余额和冻结余额
	sql := "update member_wallet set balance=balance-?, frozen_balance=frozen_balance+? where member_id=? and coin_name=?"
	// 执行SQL语句，减少余额并增加冻结余额
	err := session.Model(&model.MemberWallet{}).Exec(sql, money, money, memberId, symbol).Error
	// 返回执行结果的错误信息
	return err
}

func (m *MemberWalletDao) UpdateWallet(ctx context.Context, conn db.DbConn, id int64, balance float64, frozenBalance float64) error {
	//TODO implement me
	panic("implement me")
}

func (m *MemberWalletDao) FindByMemberId(ctx context.Context, memId int64) (list []*model.MemberWallet, err error) {
	session := m.conn.Session(ctx)
	err = session.Model(&model.MemberWallet{}).Where("member_id=?", memId).Find(&list).Error
	return
}

func (m *MemberWalletDao) UpdateAddress(ctx context.Context, wallet *model.MemberWallet) error {
	updateSql := "update member_wallet set address=?,address_private_key=? where id=?"
	session := m.conn.Session(ctx)
	err := session.Model(&model.MemberWallet{}).Exec(updateSql, wallet.Address, wallet.AddressPrivateKey, wallet.Id).Error
	return err
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
