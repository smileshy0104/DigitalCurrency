package tran

import (
	"common/db"
	"common/db/gorms"
	"gorm.io/gorm"
)

type TransactionImpl struct {
	conn db.DbConn
}

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

func NewTransaction(db *gorm.DB) *TransactionImpl {
	return &TransactionImpl{
		conn: gorms.New(db),
	}
}
