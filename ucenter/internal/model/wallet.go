package model

import (
	"github.com/jinzhu/copier"
	"grpc-common/market/mk_client"
	"grpc-common/market/types/market"
)

// MemberWallet 用户钱包结构体
// 用于表示用户在系统中的钱包信息，包括余额、冻结余额等
type MemberWallet struct {
	Id                int64   `gorm:"column:id"`              // 钱包ID
	Address           string  `gorm:"column:address"`         // 钱包地址
	Balance           float64 `gorm:"column:balance"`         // 可用余额
	FrozenBalance     float64 `gorm:"column:frozen_balance"`  // 冻结余额
	ReleaseBalance    float64 `gorm:"column:release_balance"` // 待释放余额
	IsLock            int     `gorm:"column:is_lock"`         // 是否锁定
	MemberId          int64   `gorm:"column:member_id"`       // 用户ID
	Version           int     `gorm:"column:version"`         // 版本号
	CoinId            int64   `gorm:"column:coin_id"`         // 币种ID
	ToReleased        float64 `gorm:"column:to_released"`     // 待释放的金额
	CoinName          string  `gorm:"column:coin_name"`       // 币种名称
	AddressPrivateKey string  `gorm:"-"`                      // 地址私钥，不在数据库中存储
}

// TableName 表名
// 返回MemberWallet对应的数据库表名
func (*MemberWallet) TableName() string {
	return "member_wallet"
}

// Copy 复制钱包信息到MemberWalletCoin
// 参数:
//
//	coinInfo *mk_client.Coin: 币种信息
//
// 返回值:
//
//	*MemberWalletCoin: 复制后的钱包信息对象
func (w *MemberWallet) Copy(coinInfo *mk_client.Coin) *MemberWalletCoin {
	mc := &MemberWalletCoin{}
	copier.Copy(mc, w)
	coin := &market.Coin{}
	copier.Copy(coin, coinInfo)
	mc.Coin = coin
	return mc
}

// MemberWalletCoin 用户钱包及币种信息结构体
// 用于表示用户钱包信息以及关联的币种详细信息
type MemberWalletCoin struct {
	Id             int64        `json:"id" from:"id"`                         // 钱包ID
	Address        string       `json:"address" from:"address"`               // 钱包地址
	Balance        float64      `json:"balance" from:"balance"`               // 可用余额
	FrozenBalance  float64      `json:"frozenBalance" from:"frozenBalance"`   // 冻结余额
	ReleaseBalance float64      `json:"releaseBalance" from:"releaseBalance"` // 待释放余额
	IsLock         int          `json:"isLock" from:"isLock"`                 // 是否锁定
	MemberId       int64        `json:"memberId" from:"memberId"`             // 用户ID
	Version        int          `json:"version" from:"version"`               // 版本号
	Coin           *market.Coin `json:"coin" from:"coinId"`                   // 币种信息
	ToReleased     float64      `json:"toReleased" from:"toReleased"`         // 待释放的金额
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
