package database

import "gorm.io/gorm"

type DbConn interface {
	Begin()
	Rollback()
	Commit()
}

type DB struct {
	Conn *gorm.DB
}
