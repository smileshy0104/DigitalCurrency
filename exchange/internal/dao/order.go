package dao

import (
	"common/db"
	"common/db/gorms"
	"context"
	"exchange/internal/model"
)

type OrderDao struct {
	conn *gorms.GormConn
}

// FindOrderHistory 查询指定会员的订单历史记录。
// 该方法根据符号（symbol）、页码（page）、每页大小（size）和会员ID（memberId）来查询订单信息。
// 返回值包括订单列表（list）、总订单数（total）和可能的错误（err）。
func (d *OrderDao) FindOrderHistory(ctx context.Context, symbol string, page int64, size int64, memberId int64) (list []*model.ExchangeOrder, total int64, err error) {
	// 创建数据库会话。
	session := d.conn.Session(ctx)

	// 查询符合条件的订单列表。
	err = session.Model(&model.ExchangeOrder{}).
		Where("symbol=? and member_id=?", symbol, memberId).
		Limit(int(size)).
		Offset(int((page - 1) * size)).Find(&list).Error

	// 查询符合条件的总订单数。
	err = session.Model(&model.ExchangeOrder{}).
		Where("symbol=? and member_id=?", symbol, memberId).
		Count(&total).Error

	// 返回查询结果和总订单数。
	return
}

// FindOrderCurrent 查询指定会员的订单（当前交易中）。
// 该方法根据符号（symbol）、页码（page）、每页大小（size）和会员ID（memberId）来查询订单信息。
// 返回值包括订单列表（list）、总订单数（total）和可能的错误（err）。
func (d *OrderDao) FindOrderCurrent(ctx context.Context, symbol string, page int64, size int64, memberId int64) (list []*model.ExchangeOrder, total int64, err error) {
	// 创建数据库会话。
	session := d.conn.Session(ctx)

	// 查询符合条件的订单列表。
	err = session.Model(&model.ExchangeOrder{}).
		Where("symbol=? and member_id=? and status=?", symbol, memberId, model.Trading).
		Limit(int(size)).
		Offset(int((page - 1) * size)).Find(&list).Error

	// 查询符合条件的总订单数。
	err = session.Model(&model.ExchangeOrder{}).
		Where("symbol=? and member_id=? and status=?", symbol, memberId, model.Trading).
		Count(&total).Error

	// 返回查询结果和总订单数。
	return
}

// FindCurrentTradingCount 查询当前用户在指定交易对下的订单数（交易中的订单）。
func (d *OrderDao) FindCurrentTradingCount(ctx context.Context, id int64, symbol string, direction int) (total int64, err error) {
	session := d.conn.Session(ctx)
	err = session.Model(&model.ExchangeOrder{}).
		Where("symbol=? and member_id=? and direction=? and status=?", symbol, id, direction, model.Trading).
		Count(&total).Error
	return
}

func (d *OrderDao) Save(ctx context.Context, conn db.DbConn, order *model.ExchangeOrder) error {
	//TODO implement me
	panic("implement me")
}

func NewOrderDao(db *db.DB) *OrderDao {
	return &OrderDao{
		conn: gorms.New(db.Conn),
	}
}
