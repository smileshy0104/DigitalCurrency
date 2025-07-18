package processor

import (
	"common/db"
	"context"
	"exchange/internal/database"
	"exchange/internal/model"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/market/mk_client"
	"grpc-common/market/types/market"
	"sync"
)

// CoinTradeFactory 工厂，专门生产对应 symbol 的交易引擎
type CoinTradeFactory struct {
	tradeMap map[string]*CoinTrade // 存储不同交易对的交易引擎
	mux      sync.RWMutex          // 读写锁，确保并发安全
}

// NewCoinTradeFactory 创建新的 CoinTradeFactory 实例
func NewCoinTradeFactory() *CoinTradeFactory {
	return &CoinTradeFactory{
		tradeMap: make(map[string]*CoinTrade), // 初始化交易引擎映射
	}
}

// Init 初始化 CoinTradeFactory，创建所有可见的交易引擎
func (c *CoinTradeFactory) Init(marketRpc mk_client.Market, cli *database.KafkaClient, db *db.DB) {
	// 初始化的操作
	// 查询所有的 exchange_coin 内容，循环创建交易引擎
	ctx := context.Background()
	// 查询所有可见的交易货币
	exchangeCoinRes, err := marketRpc.FindExchangeCoinVisible(ctx, &market.MarketReq{})
	if err != nil {
		logx.Error(err) // 记录错误信息
		return
	}
	// 循环创建交易引擎
	for _, v := range exchangeCoinRes.List {
		c.AddCoinTrade(v.Symbol, NewCoinTrade(v.Symbol, cli, db)) // 添加交易引擎到工厂
	}
}

// AddCoinTrade 添加一个新的 CoinTrade 实例到工厂
func (c *CoinTradeFactory) AddCoinTrade(symbol string, ct *CoinTrade) {
	c.mux.Lock()            // 锁定以进行写操作
	defer c.mux.Unlock()    // 解锁
	c.tradeMap[symbol] = ct // 将交易引擎添加到映射中
}

// GetCoinTrade 获取指定 symbol 的 CoinTrade 实例
func (c *CoinTradeFactory) GetCoinTrade(symbol string) *CoinTrade {
	c.mux.RLock()             // 锁定以进行读操作
	defer c.mux.RUnlock()     // 解锁
	return c.tradeMap[symbol] // 返回对应的交易引擎
}

// CoinTrade 撮合交易引擎，每一个交易对 BTC/USDT 都有各自的一个引擎
type CoinTrade struct {
	symbol string // 交易对符号
}

// NewCoinTrade 创建新的 CoinTrade 实例
func NewCoinTrade(symbol string, cli *database.KafkaClient, db *db.DB) *CoinTrade {
	return &CoinTrade{
		symbol: symbol, // 设置交易对符号
	}
}

// Trade  撮合交易核心代码
func (t *CoinTrade) Trade(exchangeOrder *model.ExchangeOrder) {

}
