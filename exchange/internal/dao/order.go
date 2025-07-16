package dao

import (
	"common/db"
	"common/db/gorms"
	"context"
	"exchange/internal/model"
	"gorm.io/gorm"
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
	// 查询指定用户、指定交易对、指定方向的订单数。
	err = session.Model(&model.ExchangeOrder{}).
		Where("symbol=? and member_id=? and direction=? and status=?", symbol, id, direction, model.Trading).
		Count(&total).Error
	return
}

// Save 保存订单信息到数据库。
// 该方法使用传入的数据库连接执行保存操作，并处理事务。
// 参数:
//
//	ctx context.Context: 上下文对象，用于传递请求范围的值和控制取消操作。
//	conn db.DbConn: 数据库连接对象，用于执行数据库操作。
//	order *model.ExchangeOrder: 待保存的订单对象。
//
// 返回值:
//
//	error: 如果保存操作失败，则返回错误信息；否则返回nil。
func (d *OrderDao) Save(ctx context.Context, conn db.DbConn, order *model.ExchangeOrder) error {
	// 将传入的数据库连接转换为GormConn类型
	d.conn = conn.(*gorms.GormConn)
	// 获取上下文中的事务对象
	tx := d.conn.Tx(ctx)
	// 使用事务对象保存订单信息，并处理可能发生的错误
	err := tx.Save(&order).Error
	// 返回保存操作的结果，如果有错误则返回错误信息
	return err
}

// FindOrderByOrderId 根据订单ID查找订单。
func (e *OrderDao) FindOrderByOrderId(ctx context.Context, orderId string) (order *model.ExchangeOrder, err error) {
	// 创建数据库会话。
	session := e.conn.Session(ctx)
	// 使用订单ID查询订单信息。
	err = session.Model(&model.ExchangeOrder{}).
		Where("order_id=?", orderId).Take(&order).Error
	// 如果未找到订单且没有其他错误，返回nil, nil。
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	// 返回查询到的订单和可能的错误。
	return
}

// UpdateStatusCancel 更新订单状态为取消。
func (e *OrderDao) UpdateStatusCancel(ctx context.Context, orderId string) error {
	// 创建数据库会话。
	session := e.conn.Session(ctx)
	// 更新订单状态为取消。
	err := session.Model(&model.ExchangeOrder{}).
		Where("order_id=?", orderId).Update("status", model.Canceled).Error
	// 返回更新操作的结果。
	return err
}

// UpdateOrderStatusTrading 更新订单状态为交易中。
func (e *OrderDao) UpdateOrderStatusTrading(ctx context.Context, orderId string) error {
	// 创建数据库会话。
	session := e.conn.Session(ctx)
	// 更新订单状态为交易中。
	err := session.Model(&model.ExchangeOrder{}).
		Where("order_id=?", orderId).Update("status", model.Trading).Error
	// 返回更新操作的结果。
	return err
}

// FindOrderListBySymbol 根据交易对符号和订单状态查找订单列表。
// 该方法主要用于获取特定交易对和特定状态下的所有订单，常用于订单查询和管理场景。
func (e *OrderDao) FindOrderListBySymbol(ctx context.Context, symbol string, status int) (list []*model.ExchangeOrder, err error) {
	// 创建数据库会话
	session := e.conn.Session(ctx)
	// 执行查询操作，并处理错误
	err = session.Model(&model.ExchangeOrder{}).
		Where("symbol=? and status=?", symbol, status).Find(&list).Error
	// 返回查询结果和可能的错误
	return
}

// UpdateOrderComplete 更新订单为完成状态。
// 该方法接收订单ID、交易量、营业额和订单状态作为参数，更新数据库中的订单信息。
// 主要用于在订单完成后，更新订单的详细信息和状态。
func (e *OrderDao) UpdateOrderComplete(ctx context.Context, orderId string, tradedAmount float64,
	turnover float64, status int) error {
	// 创建数据库会话
	session := e.conn.Session(ctx)
	// 准备更新SQL语句
	updateSql := "update exchange_order set traded_amount=?,turnover=?,status=? where order_id=? and status=?"
	// 执行更新操作，并处理错误
	err := session.Model(&model.ExchangeOrder{}).Exec(updateSql, tradedAmount, turnover, status, orderId, model.Trading).Error
	// 返回可能的错误
	return err
}

func NewOrderDao(db *db.DB) *OrderDao {
	return &OrderDao{
		conn: gorms.New(db.Conn),
	}
}
