package database

import (
	"common/db"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConnMysql 连接到MySQL数据库并配置连接池。
func ConnMysql(dsn string) *db.DB {
	// 尝试连接到MySQL数据库。
	var err error
	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	// 如果连接失败，抛出panic。
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	// 获取底层的数据库连接。
	sqlDb, _ := _db.DB()
	// 连接池配置
	// 设置最大打开的连接数。
	sqlDb.SetMaxOpenConns(100)
	// 设置最大空闲的连接数。
	sqlDb.SetMaxIdleConns(10)
	// 返回数据库连接实例。
	return &db.DB{
		_db,
	}
}
