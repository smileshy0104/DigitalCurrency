package task

import (
	"github.com/go-co-op/gocron"
	"job-center/internal/logic"
	"job-center/internal/svc"
	"time"
)

// Task 定时任务
// 该结构体用于管理定时任务，包括任务的创建、运行和停止
type Task struct {
	s   *gocron.Scheduler   // s 是调度器，用于控制任务的执行时间
	ctx *svc.ServiceContext // ctx 是服务上下文，包含任务执行所需的配置和依赖
}

// NewTask 创建新的定时任务
func NewTask(ctx *svc.ServiceContext) *Task {
	return &Task{
		s:   gocron.NewScheduler(time.UTC), // 创建一个新的调度器，并设置时区为UTC
		ctx: ctx,
	}
}

// Run 启动定时任务
// 该方法定义了各种时间间隔的定时任务，例如每分钟、每小时、每天等
func (t *Task) Run() {
	// 定义各种时间间隔的定时任务，例如每分钟、每小时、每天等
	t.s.Every(1).Minute().Do(func() {
		logic.NewKline(t.ctx.Config.Okx, t.ctx.MongoClient, t.ctx.KafkaClient, t.ctx.Cache).Do("1m")
	})
	t.s.Every(3).Minute().Do(func() {
		logic.NewKline(t.ctx.Config.Okx, t.ctx.MongoClient, t.ctx.KafkaClient, t.ctx.Cache).Do("3m")
	})
	t.s.Every(5).Minute().Do(func() {
		logic.NewKline(t.ctx.Config.Okx, t.ctx.MongoClient, t.ctx.KafkaClient, t.ctx.Cache).Do("5m")
	})
	//t.s.Every(15).Minute().Do(func() {
	//	logic.NewKline(t.ctx.Config.Okx, t.ctx.Cache).Do("15m")
	//})
	//t.s.Every(30).Minute().Do(func() {
	//	logic.NewKline(t.ctx.Config.Okx, t.ctx.Cache).Do("30m")
	//})
	//t.s.Every(1).Hour().Do(func() {
	//	logic.NewKline(t.ctx.Config.Okx, t.ctx.Cache).Do("1H")
	//})
	//t.s.Every(2).Hour().Do(func() {
	//	logic.NewKline(t.ctx.Config.Okx, t.ctx.Cache).Do("2H")
	//})
	//t.s.Every(4).Hour().Do(func() {
	//	logic.NewKline(t.ctx.Config.Okx, t.ctx.Cache).Do("4H")
	//})
	//t.s.Every(1).Day().Do(func() {
	//	logic.NewKline(t.ctx.Config.Okx, t.ctx.Cache).Do("1D")
	//})
	//t.s.Every(1).Week().Do(func() {
	//	logic.NewKline(t.ctx.Config.Okx, t.ctx.Cache).Do("1W")
	//})
	//t.s.Every(1).Month().Do(func() {
	//	logic.NewKline(t.ctx.Config.Okx, t.ctx.Cache).Do("1M")
	//})
	// 以下任务被注释掉，可在未来启用
	//t.s.Every(1).Minute().Do(func() {
	//	logic.NewRate(t.ctx.Config.Okx, t.ctx.Cache).Do()
	//})
	////十分钟生成一个区块
	//t.s.Every(10).Minute().Do(func() {
	//	logic.NewBitCoin(t.ctx.Cache, t.ctx.AssetRpc, t.ctx.MongoClient, t.ctx.KafkaClient).Do(t.ctx.BitCoinAddress)
	//})
}

// StartBlocking 启动定时任务并阻塞当前线程
// 该方法用于启动任务，并保持程序运行，直到任务被停止
func (t *Task) StartBlocking() {
	t.s.StartBlocking() // 调用调度器的StartBlocking方法来启动任务并阻塞当前线程
}

// Stop 停止定时任务
// 该方法用于停止所有正在运行的任务
func (t *Task) Stop() {
	t.s.Stop() // 调用调度器的Stop方法来停止所有任务
}
