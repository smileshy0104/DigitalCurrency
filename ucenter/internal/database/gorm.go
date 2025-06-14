package database

import (
	"common/db"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConn struct {
	Conn      sqlx.SqlConn
	CacheConn sqlc.CachedConn
}

func ConnMysql(dsn string) *db.DB {
	var err error
	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	sqlDb, _ := _db.DB()
	//连接池配置
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetMaxIdleConns(10)
	return &db.DB{
		_db,
	}
}
