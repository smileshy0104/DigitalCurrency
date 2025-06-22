package logic

import (
	"common/tools"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"job-center/internal/database"
	"job-center/internal/domain"
	"log"
	"strings"
	"sync"
	"time"
)

// OkxConfig 获取okx配置
type OkxConfig struct {
	ApiKey    string // API密钥
	SecretKey string // 秘密密钥
	Pass      string // 密码
	Host      string // 主机地址
	Proxy     string // 代理地址
}

// OkxResult 结构体用于解析OKX返回的结果
type OkxResult struct {
	Code string     `json:"code"` // 响应代码
	Msg  string     `json:"msg"`  // 消息描述
	Data [][]string `json:"data"` // 数据内容
}

// Kline 结构体用于处理K线数据
type Kline struct {
	wg          sync.WaitGroup      // 用于同步控制
	c           OkxConfig           // OKX配置
	klineDomain *domain.KlineDomain // K线处理模块
	queueDomain *domain.QueueDomain // 队列处理模块
	ch          cache.Cache         // 缓存接口
}

// Do 方法用于并发获取指定交易对的K线数据。
// 该方法接收一个周期参数，用于指定所需K线数据的时间周期。
// 参数 period: K线数据的时间周期，如"1m"代表1分钟周期。
func (k *Kline) Do(period string) {
	// 使用WaitGroup来等待两个并发任务完成。
	k.wg.Add(2)

	// 并发获取BTC-USDT和BTC/USDT交易对的K线数据。
	go k.getKlineData("BTC-USDT", "BTC/USDT", period)

	// 并发获取ETH-USDT和ETH/USDT交易对的K线数据。
	go k.getKlineData("ETH-USDT", "ETH/USDT", period)

	// 等待两个并发任务完成后再继续执行。
	k.wg.Wait()
}

// getKlineData 获取K线数据
func (k *Kline) getKlineData(instId string, symbol string, period string) {
	// 发起http请求 获取数据
	// TODO GET / 获取交易产品K线数据 接口
	api := k.c.Host + "/api/v5/market/candles?instId=" + instId + "&bar=" + period
	// 例子：sign=CryptoJS.enc.Base64.stringify(CryptoJS.HmacSHA256(timestamp + 'GET' + '/api/v5/account/balance?ccy=BTC', SecretKey))
	timestamp := tools.ISO(time.Now())
	sign := tools.ComputeHmacSha256(timestamp+"GET"+"/api/v5/market/candles?instId="+instId+"&bar="+period, k.c.SecretKey)
	// 设置请求头
	header := make(map[string]string)
	header["OK-ACCESS-KEY"] = k.c.ApiKey
	header["OK-ACCESS-SIGN"] = sign
	header["OK-ACCESS-TIMESTAMP"] = timestamp
	header["OK-ACCESS-PASSPHRASE"] = k.c.Pass
	// 发起http请求
	resp, err := tools.GetWithHeader(api, header, k.c.Proxy)
	if err != nil {
		log.Println(err)
		k.wg.Done()
		return
	}
	// 解析数据
	var result = &OkxResult{}
	err = json.Unmarshal(resp, result)
	if err != nil {
		log.Println(err)
		k.wg.Done()
		return
	}
	log.Println("==================执行存储mongo====================")
	if result.Code == "0" {
		//代表成功
		k.klineDomain.SaveBatch(result.Data, symbol, period)
		if "1m" == period {
			//把这个最新的数据result.Data[0] 推送到market服务，推送到前端页面，实时进行变化
			//->kafka->market kafka消费者进行数据消费-> 通过websocket通道发送给前端 ->前端更新数据
			if len(result.Data) > 0 {
				data := result.Data[0]
				k.queueDomain.Send1mKline(data, symbol)
				//放入redis 将其最新的价格
				key := strings.ReplaceAll(instId, "-", "::")
				k.ch.Set(key+"::RATE", data[4])
			}
		}
	}
	k.wg.Done()
	log.Println("==================End====================")
}

// NewKline 创建Kline实例
func NewKline(c OkxConfig, mongoClient *database.MongoClient, kafkaCli *database.KafkaClient, cache2 cache.Cache) *Kline {
	return &Kline{
		c:           c,
		klineDomain: domain.NewKlineDomain(mongoClient),
		queueDomain: domain.NewQueueDomain(kafkaCli),
		ch:          cache2,
	}
}
