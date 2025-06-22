package database

import (
	"context"
	"errors"
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
	"sync"
	"time"
)

// KafkaConfig 配置文件解析
// 用于解析Kafka客户端的配置信息
type KafkaConfig struct {
	Addr     string `json:"addr,optional"`
	WriteCap int    `json:"writeCap,optional"`
	ReadCap  int    `json:"readCap,optional"`
	// optional 是可以忽略的
	ConsumerGroup string `json:"ConsumerGroup,optional"`
}

// KafkaData Kafka数据结构
// 用于存储待发送或已接收的Kafka消息
type KafkaData struct {
	Topic string // Kafka主题
	Key   []byte // Kafka消息的键
	Data  []byte // Kafka消息的数据
}

// KafkaClient Kafka客户端结构体
// 管理Kafka的读写操作
type KafkaClient struct {
	w         *kafka.Writer  // Kafka写入器
	r         *kafka.Reader  // Kafka读取器
	readChan  chan KafkaData // 用于接收数据的通道
	writeChan chan KafkaData // 用于发送数据的通道
	c         KafkaConfig    // Kafka配置
	closed    bool           // 是否已关闭
	mutex     sync.Mutex     // 锁
}

// NewKafkaClient 创建Kafka客户端实例
// 参数c: Kafka的配置信息
func NewKafkaClient(c KafkaConfig) *KafkaClient {
	// 创建Kafka客户端实例
	return &KafkaClient{
		c: c,
	}
}

// StartWrite 启动写
// 初始化Kafka Writer并开始发送消息
func (k *KafkaClient) StartWrite() {
	// 创建Kafka Writer
	w := &kafka.Writer{
		Addr:     kafka.TCP(k.c.Addr), // Kafka地址
		Balancer: &kafka.LeastBytes{}, // 负载均衡
	}
	k.w = w                                          // 将Kafka Writer保存到KafkaClient实例中
	k.writeChan = make(chan KafkaData, k.c.WriteCap) // 创建一个容量为k.c.WriteCap的通道
	go k.sendKafka()                                 // 启动一个goroutine来发送消息
}

// Send 发送数据到写通道
// 参数data: 待发送的Kafka消息数据
func (w *KafkaClient) Send(data KafkaData) {
	// 捕获panic，防止发送消息时出现异常导致程序崩溃
	defer func() {
		if err := recover(); err != nil {
			w.closed = true
		}
	}()
	w.writeChan <- data // 将数据写入写通道
	w.closed = false    // 设置已关闭状态为false
}

// Close 关闭Kafka客户端的连接。
// 该方法首先检查是否有写连接(w.w)和读连接(w.r)，然后分别关闭它们。
// 在关闭写连接之前，使用互斥锁来确保并发安全，并且只有在客户端尚未关闭的情况下才会关闭写通道。
// 这种设计确保了即使在并发环境下，写操作的安全性和正确性也能得到保证。
func (w *KafkaClient) Close() {
	// 检查并关闭写连接。
	if w.w != nil {
		w.w.Close()
		// 获取互斥锁以保护对共享资源的访问。
		w.mutex.Lock()
		defer w.mutex.Unlock()
		// 仅在客户端尚未关闭时执行关闭写通道和设置关闭状态的操作。
		if !w.closed {
			close(w.writeChan)
			w.closed = true
		}
	}
	// 检查并关闭读连接。
	if w.r != nil {
		w.r.Close()
	}
}

// sendKafka 是 KafkaClient 的一个方法，负责监听写通道并发送消息到 Kafka
// 它使用一个无限循环来持续监听写通道中的消息，并尝试将这些消息发送到 Kafka
func (w *KafkaClient) sendKafka() {
	for {
		// 使用 select 语句监听写通道
		select {
		case data := <-w.writeChan:
			// 将从写通道接收到的消息封装成 Kafka 消息格式
			messages := []kafka.Message{
				{
					Topic: data.Topic,
					Key:   data.Key,
					Value: data.Data,
				},
			}
			// 初始化错误变量
			var err error
			// 定义重试次数常量
			const retries = 3
			// 创建一个带有超时的上下文
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			// 在方法退出时取消上下文
			defer cancel()
			// 初始化成功标志变量
			success := false
			// 开始重试循环
			for i := 0; i < retries; i++ {
				// 尝试发送消息到 Kafka
				err = w.w.WriteMessages(ctx, messages...)
				// 如果发送成功，设置成功标志并跳出循环
				if err == nil {
					success = true
					break
				}
				// 如果遇到特定错误，进行重试前的短暂休眠
				if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
					time.Sleep(time.Millisecond * 250)
					success = false
					continue
				}
				// 如果遇到其他错误，记录错误信息
				if err != nil {
					success = false
					log.Printf("kafka send writemessage err %s \n", err.Error())
				}
			}
			// 如果所有重试都失败，将消息重新发送到写通道
			if !success {
				w.Send(data)
			}
		}
	}
}

// StartRead 初始化 Kafka 客户端为读取模式，并开始读取消息。
// 该方法配置了一个 Kafka 读取器，设置了基本的读取参数如代理地址、消费者组ID以及读取数据的最小和最大字节数。
// 它还初始化了一个用于传递读取到的消息的通道，并启动了一个协程来实际执行消息读取操作。
func (k *KafkaClient) StartRead() {
	// 创建一个 Kafka 读取器，配置关键参数。
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{k.c.Addr}, // 指定 Kafka 代理地址。
		GroupID:  k.c.ConsumerGroup,  // 指定消费者组ID。
		MinBytes: 10e3,               // 设置单次读取数据的最小字节数为10KB。
		MaxBytes: 10e6,               // 设置单次读取数据的最大字节数为10MB。
	})
	// 将读取器赋值给实例变量，以便后续使用。
	k.r = r
	// 创建一个具有指定容量的通道，用于存储读取到的 Kafka 数据。
	k.readChan = make(chan KafkaData, k.c.ReadCap)
	// 启动一个协程，用于执行实际的消息读取操作。
	go k.readMsg()
}

// readMsg 是 KafkaClient 的一个方法，用于持续从 Kafka 中读取消息
// 该方法没有输入参数和返回值
func (k *KafkaClient) readMsg() {
	for {
		// 尝试从 Kafka 主题读取消息
		m, err := k.r.ReadMessage(context.Background())
		// 如果发生错误，记录错误信息并继续下一次循环
		if err != nil {
			logx.Error(err)
			continue
		}
		// 将读取到的消息封装为 KafkaData 结构
		data := KafkaData{
			Key:  m.Key,
			Data: m.Value,
		}
		// 将封装好的消息发送到 readChan 通道，供其他协程处理
		k.readChan <- data
	}
}

// Read 从读通道读取消息
// 返回读通道中的消息数据
func (k *KafkaClient) Read() KafkaData {
	// 从 readChan 读取消息
	msg := <-k.readChan
	return msg
}
