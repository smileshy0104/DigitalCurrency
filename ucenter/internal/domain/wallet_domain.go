package domain

import (
	"common/db"
	"common/db/tran"
	"common/op"
	"common/tools"
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"grpc-common/market/mk_client"
	"ucenter/internal/dao"
	"ucenter/internal/model"
	"ucenter/internal/repo"
)

// MemberWalletDomain 是一个领域对象，负责处理与会员钱包相关的业务逻辑。
// 它依赖于以下组件：
// - memberWalletRepo: 用于访问会员钱包的数据库操作接口。
// - transaction: 用于管理数据库事务的操作接口。
// - marketRpc: 用于与市场服务进行远程调用的客户端。
// - redisCache: 用于缓存数据的接口。
type MemberWalletDomain struct {
	memberWalletRepo repo.MemberWalletRepo
	transaction      tran.Transaction
	marketRpc        mk_client.Market
	redisCache       cache.Cache
}

// NewMemberWalletDomain 创建并返回一个 MemberWalletDomain 实例。
// 参数:
// - db: 数据库连接对象，用于初始化数据库相关的仓库和事务管理器。
// - marketRpc: 市场服务的 RPC 客户端，用于与市场服务交互。
// - redisCache: 缓存接口，用于存储和读取缓存数据。
// 返回值:
// - *MemberWalletDomain: 初始化完成的 MemberWalletDomain 实例。
func NewMemberWalletDomain(db *db.DB, marketRpc mk_client.Market, redisCache cache.Cache) *MemberWalletDomain {
	// 初始化 MemberWalletDomain 的各个依赖组件。
	return &MemberWalletDomain{
		memberWalletRepo: dao.NewMemberWalletDao(db),   // 使用数据库连接初始化会员钱包仓库。
		transaction:      tran.NewTransaction(db.Conn), // 使用数据库连接初始化事务管理器。
		marketRpc:        marketRpc,                    // 直接注入市场服务 RPC 客户端。
		redisCache:       redisCache,                   // 直接注入缓存接口。
	}
}

// FindWalletBySymbol 根据用户ID和币种名称查找钱包信息。
// 如果钱包不存在，则创建新的钱包并存储。
func (d *MemberWalletDomain) FindWalletBySymbol(ctx context.Context, user_id int64, coin_name string, coin *mk_client.Coin) (*model.MemberWalletCoin, error) {
	// 尝试根据用户ID和币种名称查找钱包信息。
	mw, err := d.memberWalletRepo.FindByIdAndCoinName(ctx, user_id, coin_name)
	if err != nil {
		return nil, err
	}
	// 如果钱包信息不存在，新建并存储
	if mw == nil {
		// 创建新的钱包信息(若该币种不存在，则为该用户进行新增)
		mwnew, walletCoin := model.NewMemberWallet(user_id, coin)
		// 调用存储方法Save
		err = d.memberWalletRepo.Save(ctx, mwnew)
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

func (d *MemberWalletDomain) Freeze(ctx context.Context, conn db.DbConn, userId int64, money float64, symbol string) error {
	mw, err := d.memberWalletRepo.FindByIdAndCoinName(ctx, userId, symbol)
	if err != nil {
		return err
	}
	if mw.Balance < money {
		return errors.New("余额不足")
	}
	err = d.memberWalletRepo.UpdateFreeze(ctx, conn, userId, symbol, money)
	if err != nil {
		return err
	}
	return nil
}

func (d *MemberWalletDomain) FindWalletByMemIdAndCoin(ctx context.Context, memberId int64, coinName string) (*model.MemberWallet, error) {
	mw, err := d.memberWalletRepo.FindByIdAndCoinName(ctx, memberId, coinName)
	if err != nil {
		return nil, err
	}
	return mw, nil
}

func (d *MemberWalletDomain) UpdateAddress(ctx context.Context, wallet *model.MemberWallet) error {
	return d.memberWalletRepo.UpdateAddress(ctx, wallet)
}
