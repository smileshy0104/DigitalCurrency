package dao

import (
	"common/db"
	"common/db/gorms"
	"context"
	"gorm.io/gorm"
	"market/internal/model"
)

type ExchangeCoinDao struct {
	conn *gorms.GormConn
}

func (e *ExchangeCoinDao) FindBySymbol(ctx context.Context, symbol string) (*model.ExchangeCoin, error) {
	session := e.conn.Session(ctx)
	data := &model.ExchangeCoin{}
	err := session.Model(&model.ExchangeCoin{}).Where("symbol=?", symbol).Take(data).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return data, err
}

func (e *ExchangeCoinDao) FindVisible(ctx context.Context) (list []*model.ExchangeCoin, err error) {
	session := e.conn.Session(ctx)
	err = session.Model(&model.ExchangeCoin{}).Where("visible=?", 1).Find(&list).Error
	return
}

func NewExchangeCoinDao(db *db.DB) *ExchangeCoinDao {
	return &ExchangeCoinDao{
		conn: gorms.New(db.Conn),
	}
}
