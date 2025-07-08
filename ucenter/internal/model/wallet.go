package model

import (
	"github.com/jinzhu/copier"
	"grpc-common/market/mk_client"
	"grpc-common/market/types/market"
)

type MemberWallet struct {
	Id                int64   `gorm:"column:id"`
	Address           string  `gorm:"column:address"`
	Balance           float64 `gorm:"column:balance"`
	FrozenBalance     float64 `gorm:"column:frozen_balance"`
	ReleaseBalance    float64 `gorm:"column:release_balance"`
	IsLock            int     `gorm:"column:is_lock"`
	MemberId          int64   `gorm:"column:member_id"`
	Version           int     `gorm:"column:version"`
	CoinId            int64   `gorm:"column:coin_id"`
	ToReleased        float64 `gorm:"column:to_released"`
	CoinName          string  `gorm:"column:coin_name"`
	AddressPrivateKey string  `gorm:"address_private_key"`
}

func (*MemberWallet) TableName() string {
	return "member_wallet"
}

func (w *MemberWallet) Copy(coinInfo *mk_client.Coin) *MemberWalletCoin {
	mc := &MemberWalletCoin{}
	copier.Copy(mc, w)
	coin := &market.Coin{}
	copier.Copy(coin, coinInfo)
	mc.Coin = coin
	return mc
}

type MemberWalletCoin struct {
	Id             int64        `json:"id" from:"id"`
	Address        string       `json:"address" from:"address"`
	Balance        float64      `json:"balance" from:"balance"`
	FrozenBalance  float64      `json:"frozenBalance" from:"frozenBalance"`
	ReleaseBalance float64      `json:"releaseBalance" from:"releaseBalance"`
	IsLock         int          `json:"isLock" from:"isLock"`
	MemberId       int64        `json:"memberId" from:"memberId"`
	Version        int          `json:"version" from:"version"`
	Coin           *market.Coin `json:"coin" from:"coinId"`
	ToReleased     float64      `json:"toReleased" from:"toReleased"`
}

// NewMemberWallet 创建一个新的会员钱包及其对应币种信息。
// 这个函数接收一个会员ID和一个币种对象，然后创建并返回一个MemberWallet实例和一个MemberWalletCoin实例。
// 参数:
//
//	memId - 会员ID，用于标识会员钱包的拥有者。
//	coin - 币种对象，包含了需要创建钱包的币种信息。
//
// 返回值:
//
//	*MemberWallet - 会员钱包实例，包含了会员ID、币种ID和币种名称。
//	*MemberWalletCoin - 会员钱包币种实例，包含了会员ID和币种对象。
func NewMemberWallet(memId int64, coin *market.Coin) (*MemberWallet, *MemberWalletCoin) {
	// 创建会员钱包实例，设置会员ID、币种ID和币种名称。
	mw := &MemberWallet{
		MemberId: memId,
		CoinId:   int64(coin.Id),
		CoinName: coin.Unit,
	}

	// 创建会员钱包币种实例，设置会员ID和币种对象。
	mwc := &MemberWalletCoin{
		MemberId: memId,
		Coin:     coin,
	}

	// 返回会员钱包实例和会员钱包币种实例。
	return mw, mwc
}
