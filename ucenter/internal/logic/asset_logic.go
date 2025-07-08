package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"grpc-common/market/types/market"
	"grpc-common/ucenter/types/asset"
	"ucenter/internal/domain"
	"ucenter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	memberDomain       *domain.MemberDomain
	memberWalletDomain *domain.MemberWalletDomain
}

// FindWalletBySymbol 根据币种符号查找用户的钱包信息。
// 该方法首先通过市场RPC服务查询币种信息，然后根据用户ID和币种符号查找用户的钱包信息。
func (l *AssetLogic) FindWalletBySymbol(req *asset.AssetReq) (*asset.MemberWallet, error) {
	// 通过市场RPC服务查询币种信息。
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

func NewAssetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssetLogic {
	return &AssetLogic{
		ctx:                ctx,
		svcCtx:             svcCtx,
		Logger:             logx.WithContext(ctx),
		memberDomain:       domain.NewMemberDomain(svcCtx.Db),
		memberWalletDomain: domain.NewMemberWalletDomain(svcCtx.Db, svcCtx.MarketRpc, svcCtx.Cache),
	}
}
