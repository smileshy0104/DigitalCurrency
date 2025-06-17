package logic

import (
	"common/tools"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"log"
	"sync"
	"time"
)

type OkxConfig struct {
	ApiKey    string
	SecretKey string
	Pass      string
	Host      string
	Proxy     string
}

type OkxResult struct {
	Code string     `json:"code"`
	Msg  string     `json:"msg"`
	Data [][]string `json:"data"`
}
type Kline struct {
	wg sync.WaitGroup
	c  OkxConfig
	ch cache.Cache
}

func (k *Kline) Do(period string) {
	k.wg.Add(2)
	//获取某个币 BTC-USDT ETH-USDT  BTC/USDT
	go k.getKlineData("BTC-USDT", "BTC/USDT", period)
	go k.getKlineData("ETH-USDT", "ETH/USDT", period)
	k.wg.Wait()
}

func (k *Kline) getKlineData(instId string, symbol string, period string) {
	//发起http请求 获取数据
	api := k.c.Host + "/api/v5/market/candles?instId=" + instId + "&bar=" + period
	timestamp := tools.ISO(time.Now())
	sign := tools.ComputeHmacSha256(timestamp+"GET"+"/api/v5/market/candles?instId="+instId+"&bar="+period, k.c.SecretKey)
	header := make(map[string]string)
	header["OK-ACCESS-KEY"] = k.c.ApiKey
	header["OK-ACCESS-SIGN"] = sign
	header["OK-ACCESS-TIMESTAMP"] = timestamp
	header["OK-ACCESS-PASSPHRASE"] = k.c.Pass
	resp, err := tools.GetWithHeader(api, header, k.c.Proxy)
	if err != nil {
		log.Println(err)
		k.wg.Done()
		return
	}
	var result = &OkxResult{}
	err = json.Unmarshal(resp, result)
	if err != nil {
		log.Println(err)
		k.wg.Done()
		return
	}
	log.Println("==================执行存储mongo====================")
	k.wg.Done()
	log.Println("==================End====================")
}

func NewKline(c OkxConfig, cache2 cache.Cache) *Kline {
	return &Kline{
		c:  c,
		ch: cache2,
	}
}
