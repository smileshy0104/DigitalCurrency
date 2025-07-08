package domain

import (
	"common/db"
	"common/db/tran"
	"common/op"
	"common/tools"
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
	// 如果钱包信息不存在，新建并存储
	if mw == nil {
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

// FindWallet 根据用户ID查找会员钱包信息。
// 该方法首先通过会员ID获取钱包列表，然后查询相关币种的汇率信息，
// 并计算出每个币种对应的人民币和美元价值，最后返回处理后的钱包列表。
func (d *MemberWalletDomain) FindWallet(ctx context.Context, userId int64) (list []*model.MemberWalletCoin, err error) {
	// 通过会员ID获取钱包列表
	memberWallets, err := d.memberWalletRepo.FindByMemberId(ctx, userId)
	if err != nil {
		return nil, err
	}

	// 查询cny的汇率
	var cnyRateStr string
	d.redisCache.Get("USDT::CNY::RATE", &cnyRateStr)
	var cnyRate float64 = 7
	if cnyRateStr != "" {
		cnyRate = tools.ToFloat64(cnyRateStr)
	}

	// 需要查询 币种的详情
	for _, v := range memberWallets {
		coinInfo, err := d.marketRpc.FindCoinInfo(ctx, &mk_client.MarketReq{
			Unit: v.CoinName,
		})
		if err != nil {
			return nil, err
		}
		// 根据币种类型设置汇率信息
		if coinInfo.Unit == "USDT" {
			coinInfo.CnyRate = cnyRate
			coinInfo.UsdRate = 1
		} else {
			var usdtRateStr string
			var usdtRate float64 = 20000
			d.redisCache.Get(v.CoinName+"::USDT::RATE", &usdtRateStr)
			if usdtRateStr != "" {
				usdtRate = tools.ToFloat64(usdtRateStr)
			}
			coinInfo.UsdRate = usdtRate
			coinInfo.CnyRate = op.MulFloor(cnyRate, coinInfo.UsdRate, 10)
		}
		// 将处理后的钱包信息添加到返回列表中
		list = append(list, v.Copy(coinInfo))
	}
	// 返回处理后的钱包列表
	return list, nil
}

func NewMemberWalletDomain(db *db.DB, marketRpc mk_client.Market, redisCache cache.Cache) *MemberWalletDomain {
	return &MemberWalletDomain{
		memberWalletRepo: dao.NewMemberWalletDao(db),
		transaction:      tran.NewTransaction(db.Conn),
		marketRpc:        marketRpc,
		redisCache:       redisCache,
	}
}
