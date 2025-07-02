package domain

import (
	"common/db"
	"context"
	"errors"
	"market/internal/dao"
	"market/internal/model"
	"market/internal/repo"
)

type CoinDomain struct {
	coinRepo repo.CoinRepo
}

// FindCoinInfo 获取币种信息
func (d *CoinDomain) FindCoinInfo(ctx context.Context, unit string) (*model.Coin, error) {
	// 通过币种单位查询币种信息
	coin, err := d.coinRepo.FindByUnit(ctx, unit)
	if err != nil {
		return nil, err
	}
	if coin == nil {
		return nil, errors.New("not support this coin:" + unit)
	}
	return coin, nil
}

// FindCoinId 获取币种信息
func (d *CoinDomain) FindCoinId(ctx context.Context, id int64) (*model.Coin, error) {
	// 通过币种id查询币种信息
	coin, err := d.coinRepo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	if coin == nil {
		return nil, errors.New("not support this coin")
	}
	return coin, nil
}

// FindAll 获取所有币种信息
func (d *CoinDomain) FindAll(ctx context.Context) ([]*model.Coin, error) {
	// 获取所有币种信息
	return d.coinRepo.FindAll(ctx)
}

func NewCoinDomain(db *db.DB) *CoinDomain {
	return &CoinDomain{
		coinRepo: dao.NewCoinDao(db),
	}
}
