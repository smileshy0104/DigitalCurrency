package domain

import (
	"common/db"
	"common/db/tran"
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"grpc-common/market/mk_client"
	"ucenter/internal/dao"
	"ucenter/internal/model"
	"ucenter/internal/repo"
)

type MemberWalletDomain struct {
	memberWalletRepo repo.MemberWalletRepo
	transaction      tran.Transaction
	marketRpc        mk_client.Market
	redisCache       cache.Cache
}

// FindWalletBySymbol 根据用户ID和币种名称查找钱包信息。
// 如果钱包不存在，则创建新的钱包并存储。
func (d *MemberWalletDomain) FindWalletBySymbol(ctx context.Context, id int64, name string, coin *mk_client.Coin) (*model.MemberWalletCoin, error) {
	// 尝试根据用户ID和币种名称查找钱包信息。
	mw, err := d.memberWalletRepo.FindByIdAndCoinName(ctx, id, name)
	if err != nil {
		return nil, err
	}
	if mw == nil {
		// 如果钱包信息不存在，新建并存储
		mw, walletCoin := model.NewMemberWallet(id, coin)
		err := d.memberWalletRepo.Save(ctx, mw)
		if err != nil {
			return nil, err
		}
		return walletCoin, nil
	}
	// 如果钱包信息已存在，复制到新的结构体并返回
	nwc := &model.MemberWalletCoin{}
	copier.Copy(nwc, mw)
	nwc.Coin = coin
	return nwc, nil
}

func NewMemberWalletDomain(db *db.DB, marketRpc mk_client.Market, redisCache cache.Cache) *MemberWalletDomain {
	return &MemberWalletDomain{
		memberWalletRepo: dao.NewMemberWalletDao(db),
		transaction:      tran.NewTransaction(db.Conn),
		marketRpc:        marketRpc,
		redisCache:       redisCache,
	}
}
