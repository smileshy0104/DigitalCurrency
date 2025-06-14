package db

import "gorm.io/gorm"

// DbConn db连接
type DbConn interface {
	Begin()    // 开启事务
	Rollback() // 回滚事务
	Commit()   // 提交事务
}

// DB db
type DB struct {
	Conn *gorm.DB
}
