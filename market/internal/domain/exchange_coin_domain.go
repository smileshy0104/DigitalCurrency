package domain

import (
	"common/db"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"market/internal/dao"
	"market/internal/model"
	"market/internal/repo"
)

// ExchangeCoinDomain 交易货币模块
type ExchangeCoinDomain struct {
	exchangeCoinRepo repo.ExchangeCoinRepo
}

// NewExchangeCoinDomain 创建交易货币模块
func NewExchangeCoinDomain(db *db.DB) *ExchangeCoinDomain {
	return &ExchangeCoinDomain{
		exchangeCoinRepo: dao.NewExchangeCoinDao(db),
	}
}

// FindVisible 查询所有可见的交易货币
func (d *ExchangeCoinDomain) FindVisible(ctx context.Context) []*model.ExchangeCoin {
	list, err := d.exchangeCoinRepo.FindVisible(ctx)
	if err != nil {
		logx.Error(err)
		return []*model.ExchangeCoin{}
	}
	return list
}

// FindBySymbol 根据交易对查询交易货币
func (d *ExchangeCoinDomain) FindBySymbol(ctx context.Context, symbol string) (*model.ExchangeCoin, error) {
	// 查询交易货币
	exchangeCoin, err := d.exchangeCoinRepo.FindBySymbol(ctx, symbol)
	if err != nil {
		return nil, err
	}
	if exchangeCoin == nil {
		return nil, errors.New("交易对不存在")
	}
	return exchangeCoin, nil
}
