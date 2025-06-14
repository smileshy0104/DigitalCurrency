package tran

import (
	"common/db"
	"common/db/gorms"
	"gorm.io/gorm"
)

// TransactionImpl 提供了事务管理的功能，用于在数据库操作中确保数据的一致性。
type TransactionImpl struct {
	conn db.DbConn
}

// Action 执行一个数据库事务。
// 参数 f 是一个函数，它接收一个数据库连接并执行数据库操作。
// 如果 f 中的数据库操作失败，事务将回滚并返回错误。
// 如果操作成功，事务将被提交。
func (t *TransactionImpl) Action(f func(conn db.DbConn) error) error {
	t.conn.Begin()
	err := f(t.conn)
	if err != nil {
		t.conn.Rollback()
		return err
	}
	t.conn.Commit()
	return nil
}

// NewTransaction 创建一个新的 TransactionImpl 实例。
// 参数 db 是一个 GORM 数据库实例，用于执行数据库事务。
// 该函数返回一个初始化了数据库连接的 TransactionImpl 实例。
func NewTransaction(db *gorm.DB) *TransactionImpl {
	return &TransactionImpl{
		conn: gorms.New(db),
	}
}
