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

func (d *CoinDomain) FindCoinInfo(ctx context.Context, unit string) (*model.Coin, error) {
	coin, err := d.coinRepo.FindByUnit(ctx, unit)
	if err != nil {
		return nil, err
	}
	if coin == nil {
		return nil, errors.New("not support this coin:" + unit)
	}
	return coin, nil
}
func (d *CoinDomain) FindCoinId(ctx context.Context, id int64) (*model.Coin, error) {
	coin, err := d.coinRepo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	if coin == nil {
		return nil, errors.New("not support this coin")
	}
	return coin, nil
}

func (d *CoinDomain) FindAll(ctx context.Context) ([]*model.Coin, error) {
	return d.coinRepo.FindAll(ctx)
}

func NewCoinDomain(db *db.DB) *CoinDomain {
	return &CoinDomain{
		coinRepo: dao.NewCoinDao(db),
	}
}
