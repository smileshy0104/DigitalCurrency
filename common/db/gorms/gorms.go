package gorms

import (
	"context"
	"gorm.io/gorm"
)

// GormConn 是对 GORM 的连接封装，提供事务管理和会话处理功能
type GormConn struct {
	db *gorm.DB // 普通数据库连接
	tx *gorm.DB // 事务数据库连接
}

// Begin 开启一个新的事务
// 在需要执行一系列事务操作时调用此方法
func (g *GormConn) Begin() {
	g.tx = g.db.Begin() // 开启一个新的事务
}

// New 创建并返回一个新的 GormConn 实例
// db: 要封装的 GORM DB 实例
func New(db *gorm.DB) *GormConn {
	return &GormConn{db: db, tx: db}
}

// Session 使用给定的上下文创建一个新会话
// ctx: 上下文对象，通常包含请求相关的信息
func (g *GormConn) Session(ctx context.Context) *gorm.DB {
	return g.db.WithContext(ctx)
}

// Commit 提交当前事务
func (g *GormConn) Commit() {
	if g.tx != nil {
		g.tx.Commit()
		g.tx = nil
	}
}

// Rollback 回滚当前事务
func (g *GormConn) Rollback() {
	if g.tx != nil {
		g.tx.Rollback()
		g.tx = nil
	}
}

// WithTx 执行传入的函数，并在其内部使用事务进行数据库操作
// fn: 需要在事务中执行的函数
func (g *GormConn) WithTx(fn func(tx *gorm.DB) error) error {
	g.Begin()
	defer func() {
		if r := recover(); r != nil {
			g.Rollback()
		}
	}()

	if err := fn(g.tx); err != nil {
		g.Rollback()
		return err
	}

	g.Commit()
	return nil
}
