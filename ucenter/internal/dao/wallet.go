package dao

import (
	"common/db"
	"common/db/gorms"
	"context"
	"gorm.io/gorm"
	"ucenter/internal/model"
)

// MemberWalletDao 结构体用于处理会员钱包的数据库操作
type MemberWalletDao struct {
	conn *gorms.GormConn // 数据库连接
}

// NewMemberWalletDao 创建一个新的 MemberWalletDao 实例
func NewMemberWalletDao(db *db.DB) *MemberWalletDao {
	return &MemberWalletDao{
		conn: gorms.New(db.Conn), // 初始化 GormConn
	}
}

// Save 保存会员钱包
func (m *MemberWalletDao) Save(ctx context.Context, mw *model.MemberWallet) error {
	session := m.conn.Session(ctx) // 获取数据库会话
	err := session.Save(&mw).Error // 保存钱包信息
	return err                     // 返回错误信息
}

// FindByIdAndCoinName 通过会员 ID 和币种名称查询会员钱包
func (m *MemberWalletDao) FindByIdAndCoinName(ctx context.Context, memId int64, coinName string) (mw *model.MemberWallet, err error) {
	session := m.conn.Session(ctx) // 获取数据库会话
	err = session.Model(&model.MemberWallet{}).
		Where("member_id=? and coin_name=?", memId, coinName).
		Take(&mw).Error // 查询钱包信息
	if err == gorm.ErrRecordNotFound {
		return nil, nil // 如果未找到记录，返回 nil
	}
	return // 返回钱包信息
}

// UpdateFreeze 更新会员钱包中的冻结金额
func (m *MemberWalletDao) UpdateFreeze(ctx context.Context, conn db.DbConn, memberId int64, symbol string, money float64) error {
	con := conn.(*gorms.GormConn) // 将连接对象转换为 GormConn 类型
	session := con.Tx(ctx)        // 获取事务会话
	// 准备 SQL 语句，用于更新余额和冻结余额
	sql := "update member_wallet set balance=balance-?, frozen_balance=frozen_balance+? where member_id=? and coin_name=?"
	// 执行 SQL 语句
	err := session.Model(&model.MemberWallet{}).Exec(sql, money, money, memberId, symbol).Error
	return err // 返回执行结果的错误信息
}

// UpdateWallet 更新会员钱包的余额和冻结余额
func (m *MemberWalletDao) UpdateWallet(ctx context.Context, conn db.DbConn, id int64, balance float64, frozenBalance float64) error {
	gormConn := conn.(*gorms.GormConn) // 将连接对象转换为 GormConn 类型
	tx := gormConn.Tx(ctx)             // 获取事务会话
	// 准备 SQL 语句
	updateSql := "update member_wallet set balance=?, frozen_balance=? where id=?"
	err := tx.Model(&model.MemberWallet{}).Exec(updateSql, balance, frozenBalance, id).Error // 执行更新
	return err                                                                               // 返回执行结果的错误信息
}

// FindByMemberId 根据会员 ID 查询用户钱包列表
func (m *MemberWalletDao) FindByMemberId(ctx context.Context, memId int64) (list []*model.MemberWallet, err error) {
	session := m.conn.Session(ctx)                                                           // 获取数据库会话
	err = session.Model(&model.MemberWallet{}).Where("member_id=?", memId).Find(&list).Error // 查询钱包列表
	return                                                                                   // 返回钱包列表
}

// UpdateAddress 更新用户钱包地址
func (m *MemberWalletDao) UpdateAddress(ctx context.Context, wallet *model.MemberWallet) error {
	updateSql := "update member_wallet set address=?, address_private_key=? where id=?"
	session := m.conn.Session(ctx)                                                                                         // 获取数据库会话
	err := session.Model(&model.MemberWallet{}).Exec(updateSql, wallet.Address, wallet.AddressPrivateKey, wallet.Id).Error // 执行更新
	return err                                                                                                             // 返回执行结果的错误信息
}

// FindAllAddress 查询所有用户钱包地址
func (m *MemberWalletDao) FindAllAddress(ctx context.Context, coinName string) (list []string, err error) {
	session := m.conn.Session(ctx)                                                                                // 获取数据库会话
	err = session.Model(&model.MemberWallet{}).Where("coin_name=?", coinName).Select("address").Find(&list).Error // 查询地址列表
	return                                                                                                        // 返回地址列表
}

// FindByAddress 根据地址查询用户钱包
func (m *MemberWalletDao) FindByAddress(ctx context.Context, address string) (mw *model.MemberWallet, err error) {
	session := m.conn.Session(ctx)                                                         // 获取数据库会话
	err = session.Model(&model.MemberWallet{}).Where("address=?", address).Take(&mw).Error // 查询钱包信息
	if err == gorm.ErrRecordNotFound {
		return nil, nil // 如果未找到记录，返回 nil
	}
	return // 返回钱包信息
}

// FindByIdAndCoinId 根据会员 ID 和币种 ID 查询用户钱包
func (m *MemberWalletDao) FindByIdAndCoinId(ctx context.Context, memId int64, coinId int64) (mw *model.MemberWallet, err error) {
	session := m.conn.Session(ctx) // 获取数据库会话
	err = session.Model(&model.MemberWallet{}).
		Where("member_id=? and coin_id=?", memId, coinId).
		Take(&mw).Error // 查询钱包信息
	if err == gorm.ErrRecordNotFound {
		return nil, nil // 如果未找到记录，返回 nil
	}
	return // 返回钱包信息
}
