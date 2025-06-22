package dao

import (
	"common/db"
	"common/db/gorms"
	"context"
	"gorm.io/gorm"
	"market/internal/model"
)

// ExchangeCoinDao 定义了交易所货币相关的数据访问对象
type ExchangeCoinDao struct {
	conn *gorms.GormConn
}

// FindBySymbol 根据交易货币名称symbol查询
// 该方法接收一个上下文和一个货币符号，返回该符号对应的交易货币信息或错误
func (e *ExchangeCoinDao) FindBySymbol(ctx context.Context, symbol string) (*model.ExchangeCoin, error) {
	session := e.conn.Session(ctx)
	data := &model.ExchangeCoin{}
	// 使用symbol查询数据库，如果找到则填充data，否则返回错误
	err := session.Model(&model.ExchangeCoin{}).Where("symbol=?", symbol).Take(data).Error
	// 如果错误是记录未找到，则返回nil, nil
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return data, err
}

// FindVisible 根据前台可见状态visible 查询
// 该方法接收一个上下文，返回所有前台可见的交易货币列表及可能的错误
func (e *ExchangeCoinDao) FindVisible(ctx context.Context) (list []*model.ExchangeCoin, err error) {
	session := e.conn.Session(ctx)
	// 查询所有可见的交易货币并填充list
	err = session.Model(&model.ExchangeCoin{}).Where("visible=?", 1).Find(&list).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return list, err
}

// NewExchangeCoinDao 创建并返回一个新的ExchangeCoinDao实例
// 该方法接收一个数据库连接，初始化并返回一个ExchangeCoinDao实例
func NewExchangeCoinDao(db *db.DB) *ExchangeCoinDao {
	return &ExchangeCoinDao{
		conn: gorms.New(db.Conn),
	}
}
