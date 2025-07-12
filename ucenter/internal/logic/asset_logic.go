package logic

import (
	"common/bc"
	"context"
	"github.com/jinzhu/copier"
	"grpc-common/market/types/market"
	"grpc-common/ucenter/types/asset"
	"ucenter/internal/domain"
	"ucenter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// AssetLogic 资产相关业务逻辑处理结构体
// 包含上下文、服务依赖、日志组件及成员资产领域对象
type AssetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	memberDomain       *domain.MemberDomain
	memberWalletDomain *domain.MemberWalletDomain
}

// NewAssetLogic 创建资产逻辑处理器实例
func NewAssetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssetLogic {
	return &AssetLogic{
		ctx:                ctx,
		svcCtx:             svcCtx,
		Logger:             logx.WithContext(ctx),
		memberDomain:       domain.NewMemberDomain(svcCtx.Db),
		memberWalletDomain: domain.NewMemberWalletDomain(svcCtx.Db, svcCtx.MarketRpc, svcCtx.Cache),
	}
}

// FindWalletBySymbol 根据币种符号查找用户的钱包信息。
// 该方法首先通过市场RPC服务查询币种信息，然后根据用户ID和币种符号查找用户的钱包信息。
func (l *AssetLogic) FindWalletBySymbol(req *asset.AssetReq) (*asset.MemberWallet, error) {
	// 通过MarketRpc服务中的查询币种信息。
	coinInfo, err := l.svcCtx.MarketRpc.FindCoinInfo(l.ctx, &market.MarketReq{
		Unit: req.CoinName,
	})
	if err != nil {
		return nil, err
	}

	// 根据用户ID、币种符号和币种信息查找用户的钱包信息。
	memberWalletCoin, err := l.memberWalletDomain.FindWalletBySymbol(l.ctx, req.UserId, req.CoinName, coinInfo)
	if err != nil {
		return nil, err
	}

	// 创建响应对象并将找到的钱包信息复制到响应对象中。
	resp := &asset.MemberWallet{}
	copier.Copy(resp, memberWalletCoin)

	// 返回用户的钱包信息。
	return resp, nil
}

// FindWallet 查找用户的所有钱包信息
func (l *AssetLogic) FindWallet(req *asset.AssetReq) (*asset.MemberWalletList, error) {
	//根据用户id查询用户的钱包 循环钱包信息 根据币种 查询币种详情
	memberWalletCoins, err := l.memberWalletDomain.FindWallet(l.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	var list []*asset.MemberWallet
	copier.Copy(&list, memberWalletCoins)
	return &asset.MemberWalletList{
		List: list,
	}, nil
}

func (l *AssetLogic) ResetAddress(req *asset.AssetReq) (*asset.AssetResp, error) {
	//查询用户的钱包 检查address是否为空 如果未空 生成地址 进行更新
	memberWallet, err := l.memberWalletDomain.FindWalletByMemIdAndCoin(l.ctx, req.UserId, req.CoinName)
	if err != nil {
		return nil, err
	}
	if req.CoinName == "BTC" {
		if memberWallet.Address == "" {
			wallet, err := bc.NewWallet()
			if err != nil {
				return nil, err
			}
			address := wallet.GetTestAddress()
			priKey := wallet.GetPriKey()
			memberWallet.AddressPrivateKey = priKey
			memberWallet.Address = string(address)
			err = l.memberWalletDomain.UpdateAddress(l.ctx, memberWallet)
			if err != nil {
				return nil, err
			}
		}
	}
	return &asset.AssetResp{}, nil
}
