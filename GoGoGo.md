# GoStudy
## 一、GMP模型
### 1. 模型介绍
GMP模型是Go语言并发编程的核心，旨在通过高效管理Goroutine的调度和执行，充分利用多核CPU的计算能力。GMP的三个组成部分分别是G（Goroutine）、M（Machine）和P（Processor）
### 2. 组成部分
#### 2.1  G（Goroutine）
* 定义：G（Goroutine）是Go语言中的基本单位，Goroutine是一个轻量级的线程，Goroutine的创建和销毁非常快，Goroutine的创建和销毁不会产生额外的开销。
* 状态：Goroutine的状态包括：
* * 运行中：当前正在执行。
* * 就绪：准备好执行，但等待分配到M上。
* * 阻塞：等待某个事件（如I/O操作）完成。
* * 死亡：执行完毕，资源被释放。
#### 2.2 M（Machine）
* 定义：M代表操作系统线程，Go运行时通过M来执行Goroutine。M是与操作系统直接交互的实体。
* 数量管理：M的数量通常与CPU核心数相匹配，Go运行时会根据系统负载动态调整M的数量。
* 生命周期：M的生命周期由Go运行时管理，可以被创建、销毁和重用。
#### 2.3 P（Processor）
* 定义：P代表逻辑处理器，是Goroutine的调度器。每个P负责管理一组G，并将其分配到M上执行。
* G队列：每个P有一个G队列，用于存储就绪的G。P从队列中选择G进行调度。
* 数量配置：P的数量可以通过GOMAXPROCS环境变量设置，通常设置为系统的CPU核心数，以最大化并发性能。


## 二、Go Struct 能否比较
### 1. 结构体比较的条件
* 可比较性：结构体的所有字段都必须是可比较的（即可以使用==和!=操作符）。如果结构体中包含**不可比较的字段（如切片、映射、函数等）**，则该结构体也不可比较。
* 直接比较：可以直接**使用==和!=操作符**比较结构体实例。
```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p1 := Person{"Alice", 30}
	p2 := Person{"Alice", 30}
	p3 := Person{"Bob", 25}

	fmt.Println(p1 == p2) // true
	fmt.Println(p1 == p3) // false
}
```
## 三、Go defer
### 1. defer 的基本用法
defer语句用于**在函数返回之前执行特定的操作**。它通常用于**资源清理**（如关闭文件、解锁mutex等）。
### 2. defer 在循环中的使用
在循环中使用defer时，要注意defer的**执行顺序是后进先出（LIFO）**，并且defer语句的参数在声明时就被求值，而不是在执行时。
```go
package main

import "fmt"

func main() {
    for i := 0; i < 3; i++ {
        defer fmt.Println(i) // 在循环结束后按逆序打印
    }
}
// 输出 2 1 0
```
* 注意事项<br>
资源管理：在循环中使用defer可能会导致资源未及时释放，尤其是在大循环中。建议在需要时使用defer，而在高频率的循环中，考虑直接在循环结束时释放资源。

## 四、select作用
### 1. select作用<br>
select语句用于处理多个通信操作。select会等待直到某个case准备好了，然后执行该case。如果没有case准备好，select会阻塞。<br>
### 2. select使用场景
* 多路复用：select可以同时处理多个通信操作。
* 负载均衡：select可以处理多个服务器的请求，并选择一个处理请求的服务器。
* 超时控制：可以结合time.After实现超时机制。
* 默认分支：如果没有通道准备好，可以使用default分支来执行某些操作。
```go
// 多路复用
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "from channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "from channel 2"
	}()
    // select可以同时等待多个通道的操作（发送或接收），当其中一个通道准备好时，就会执行对应的代码块。
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}
/* 输出：
    from channel 1
    from channel 2
 */
```
```go
// 超时控制
select {
    case msg := <-ch:
        fmt.Println("Received:", msg)
    case <-time.After(2 * time.Second):
        fmt.Println("Timeout!")
}
```
## 五、context包的用途
### 1. 概念
context包在Go中用于在**多个Goroutine之间传递上下文信息**。<br>
它主要用于控制Goroutine的生命周期、传递请求范围的值、取消信号和超时控制。
### 2. 使用场景
1. **取消信号**：在某些场景下，需要取消某个Goroutine，可以使用context包中的CancelFunc来取消。
```go
package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // 监听取消信号
			fmt.Println("Worker stopped:", ctx.Err())
			return
		default:
			// 模拟工作
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background()) // 创建取消上下文
	go worker(ctx) // 启动工作协程

	time.Sleep(2 * time.Second) // 主协程等待
	cancel() // 发送取消信号
	time.Sleep(1 * time.Second) // 等待协程处理完毕
}
```
2. **值传递**：可以在上下文中传递请求范围的值，比如用户ID、请求ID等，方便在多个Goroutine之间共享数据。
```go
package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "userID", 123) // 创建带值的上下文
	userID := ctx.Value("userID").(int) // 从上下文中获取值
	fmt.Println("User ID:", userID) // 输出: User ID: 123
}
```
3. **超时处理**：可以使用`context.WithTimeout`创建一个带有超时的上下文，当超时后，会自动取消该上下文。
```go
package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second): // 模拟长时间工作
		fmt.Println("Worker finished work")
	case <-ctx.Done(): // 监听取消信号
		fmt.Println("Worker stopped due to timeout:", ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second) // 设置超时
	defer cancel() // 确保在主协程结束前调用cancel

	go worker(ctx) // 启动工作协程
	time.Sleep(3 * time.Second) // 主协程等待
}
```
## 六、Client 如何实现长连接
### 1. 概念
长连接是指在TCP连接上进行多次数据传输，而不需要每次传输都重新建立连接。在Go语言中，我们可以使用net包中的Conn类型来实现长连接。
### 2. 使用TCP长连接
* 使用net包：可以使用net.Dial建立TCP连接，并保持该连接以便后续通信。
* 保持连接：通过循环发送数据或定期心跳保持连接。
```go
package main

import (
    "fmt"
    "net"
    "time"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8080") // 建立TCP连接
    if err != nil {
        fmt.Println("Error connecting:", err)
        return
    }
    defer conn.Close() // 确保连接在结束时关闭

    for {
        _, err := conn.Write([]byte("Hello Server\n")) // 发送数据
        if err != nil {
            fmt.Println("Error writing:", err)
            break
        }
        time.Sleep(1 * time.Second) // 模拟间隔发送
    }
}
```
### 3. 使用HTTP长连接
* HTTP Keep-Alive：HTTP/1.1默认支持长连接，可以通过设置请求头来实现。
* 使用http.Client：可以配置HTTP客户端以保持连接。
```go
package main

import (
    "fmt"
    "net/http"
    "time"
)

func main() {
    client := &http.Client{
        Transport: &http.Transport{
            DisableKeepAlives: false, // 启用Keep-Alive
        },
    }

    for {
        resp, err := client.Get("http://localhost:8080") // 发送HTTP请求
        if err != nil {
            fmt.Println("Error:", err)
            break
        }
        resp.Body.Close() // 关闭响应体
        time.Sleep(1 * time.Second) // 模拟间隔请求
    }
}
```
## 七、WaitGroup
### 1. 概念
在Go中，主协程可以使用sync.WaitGroup来等待其他协程完成。WaitGroup允许主协程等待一组Goroutine的完成。
### 2. 使用WaitGroup
```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // 完成时通知WaitGroup
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(2 * time.Second) // 模拟工作
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup // 创建WaitGroup实例
    for i := 1; i <= 3; i++ {
        wg.Add(1) // 增加计数
        go worker(i, &wg) // 启动Goroutine
    }
    wg.Wait() // 等待所有Goroutine完成
    fmt.Println("All workers completed")
}
```