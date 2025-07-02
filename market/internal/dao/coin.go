package dao

import (
	"common/db"
	"common/db/gorms"
	"context"
	"gorm.io/gorm"
	"market/internal/model"
)

// CoinDao 定义了操作货币相关数据库的接口
type CoinDao struct {
	conn *gorms.GormConn
}

// FindById 通过id查询货币
// 该方法接收一个上下文和一个货币ID，返回指定的货币实例或错误
func (d *CoinDao) FindById(ctx context.Context, id int64) (*model.Coin, error) {
	session := d.conn.Session(ctx)
	coin := &model.Coin{}
	// 通过id查询货币
	err := session.Model(&model.Coin{}).Where("id=?", id).Take(coin).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return coin, err
}

// FindAll 查询所有货币
// 该方法接收一个上下文，返回所有货币的列表及可能的错误
func (d *CoinDao) FindAll(ctx context.Context) (list []*model.Coin, err error) {
	session := d.conn.Session(ctx)
	// 查询所有货币并填充list
	err = session.Model(&model.Coin{}).Find(&list).Error
	return
}

// FindByUnit 通过单位unit查询货币
// 该方法接收一个上下文和一个货币单位，返回指定单位的货币实例或错误
func (d *CoinDao) FindByUnit(ctx context.Context, unit string) (*model.Coin, error) {
	session := d.conn.Session(ctx)
	coin := &model.Coin{}
	// 使用unit查询数据库，如果找到则填充coin，否则返回错误
	err := session.Model(&model.Coin{}).Where("unit=?", unit).Take(coin).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return coin, err
}

// NewCoinDao 创建一个新的CoinDao实例
// 该函数接收一个数据库连接，返回一个新的CoinDao实例
func NewCoinDao(db *db.DB) *CoinDao {
	return &CoinDao{
		conn: gorms.New(db.Conn),
	}
}
