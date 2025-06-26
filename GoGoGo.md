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
## 八、Slice相关内容
### 1. 概念
Slice是一种动态数组，它可以存储任意数量的元素。在Go中，slice是一个结构体，包含三个字段：指针、长度和容量。指针指向数组的起始位置，长度表示slice的长度，容量表示slice可以存储的元素数量。
### 2. 创建和初始化
```go
slice := []int{1, 2, 3} // 创建一个长度和容量均为3的切片
```
### 3. len和cap
* len：返回切片的长度。
* cap：返回切片的容量。
```go
fmt.Println(len(slice)) // 输出: 3
fmt.Println(cap(slice)) // 输出: 3
```
### 4. 共享
* 切片是引用类型，多个切片可以共享同一个底层数组。
* 修改一个切片的元素会影响到所有共享该底层数组的切片。
```go
slice1 := []int{1, 2, 3}
slice2 := slice1 // slice2与slice1共享底层数组
slice2[0] = 10 // 修改slice2的元素
fmt.Println(slice1) // 输出: [10 2 3]
```
### 5. 扩容
* 当向切片中添加元素超出其容量时，Go会自动扩容，通常会将容量翻倍。
* 扩容会创建一个新的底层数组，并将旧数据复制到新数组中。
```go
slice := make([]int, 0, 2) // 初始长度为0，容量为2
slice = append(slice, 1, 2) // 添加两个元素
fmt.Println(len(slice), cap(slice)) // 输出: 2 2
slice = append(slice, 3) // 添加第三个元素，触发扩容
fmt.Println(len(slice), cap(slice)) // 输出: 3 4
```
## 九、map如何顺序读取
### 1. 概念
在Go语言中，map是无序的，这意味着在插入元素时，它们的顺序并不一定会被保留。
### 2. 实现
```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    // 创建一个map
    myMap := map[string]int{
        "apple":  5,
        "banana": 2,
        "orange": 3,
    }

    // 提取map的键到切片
    keys := make([]string, 0, len(myMap))
    for key := range myMap {
        keys = append(keys, key)
    }

    // 对键进行排序
    sort.Strings(keys)

    // 顺序读取map
    for _, key := range keys {
        fmt.Printf("%s: %d\n", key, myMap[key])
    }
}
```
## 十、自定义实现Set
### 1. 实现
```go
package main

import "fmt"

// 定义一个Set类型（元素无序且不重复）
type Set struct {
    items map[string]struct{}
}

// 创建新的Set
func NewSet() *Set {
    return &Set{
        items: make(map[string]struct{}),
    }
}

// 添加元素
func (s *Set) Add(item string) {
    s.items[item] = struct{}{}
}

// 删除元素
func (s *Set) Remove(item string) {
    delete(s.items, item)
}

// 检查元素是否存在
func (s *Set) Contains(item string) bool {
    _, exists := s.items[item]
    return exists
}

// 获取集合大小
func (s *Set) Size() int {
    return len(s.items)
}

// 打印集合
func (s *Set) Print() {
    for key := range s.items {
        fmt.Println(key)
    }
}

func main() {
    set := NewSet()
    set.Add("apple")
    set.Add("banana")
    set.Add("orange")

    fmt.Println("Set contains apple:", set.Contains("apple")) // true
    fmt.Println("Set size:", set.Size())                     // 3

    set.Remove("banana")
    fmt.Println("Set contains banana:", set.Contains("banana")) // false

    fmt.Println("Elements in set:")
    set.Print() // 打印集合中的元素
}
```
## 十一、实现消息队列（多生产者，多消费者）
### 1. 概念
在Go中，可以使用通道（channel）来实现消息队列。通过创建一个缓冲通道，可以允许多个生产者和多个消费者进行并发操作。
### 2. 实现
```go
package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

const (
    numProducers = 3
    numConsumers = 2
)

func producer(queue chan<- int, id int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 5; i++ {
        item := rand.Intn(100) // 生成随机数
        queue <- item          // 发送到队列
        fmt.Printf("Producer %d produced %d\n", id, item)
        time.Sleep(time.Millisecond * 500) // 模拟工作
    }
}

func consumer(queue <-chan int, id int, wg *sync.WaitGroup) {
    defer wg.Done()
    for item := range queue {
        fmt.Printf("Consumer %d consumed %d\n", id, item)
        time.Sleep(time.Millisecond * 1000) // 模拟处理时间
    }
}

func main() {
    queue := make(chan int, 10) // 创建缓冲通道
    var wg sync.WaitGroup

    // 启动生产者
    for i := 1; i <= numProducers; i++ {
        wg.Add(1)
        go producer(queue, i, &wg)
    }

    // 启动消费者
    for i := 1; i <= numConsumers; i++ {
        wg.Add(1)
        go consumer(queue, i, &wg)
    }

    wg.Wait() // 等待所有生产者完成
    close(queue) // 关闭队列

    wg.Wait() // 等待所有消费者完成
}
/*
在这个示例中，我们创建了一个缓冲通道queue，允许多个生产者将数据发送到队列中，同时多个消费者从队列中接收数据。
生产者生成随机数并发送到队列，消费者从队列中接收数据并处理。
 */
```
## 十二、 大文件排序
### 1. 实现
```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
)

func sortFileChunk(filePath string, chunkSize int) ([]string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var lines []string
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
        if len(lines) == chunkSize {
            sort.Strings(lines) // 对当前块进行排序
            break
        }
    }
    return lines, scanner.Err()
}

func main() {
    filePath := "largefile.txt" // 假设这是一个大文件
    chunkSize := 1000           // 每次读取1000行进行排序

    sortedLines, err := sortFileChunk(filePath, chunkSize)
    if err != nil {
        fmt.Println("Error sorting file chunk:", err)
        return
    }

    // 将排序后的结果写入新文件
    outputFile, err := os.Create("sortedfile.txt")
    if err != nil {
        fmt.Println("Error creating output file:", err)
        return
    }
    defer outputFile.Close()

    writer := bufio.NewWriter(outputFile)
    for _, line := range sortedLines {
        writer.WriteString(line + "\n")
    }
    writer.Flush()
}
```
## 十三、HTTP能不能一次连接多次请求，不等后端返回
### 1. HTTP/1.1的持久连接
在HTTP/1.1中，默认启用了持久连接（Keep-Alive），允许在同一TCP连接上发送多个请求，而无需为每个请求重新建立连接。这意味着客户端可以在一个连接上连续发送多个请求，但通常仍然需要等待服务器响应。
### 2. HTTP/2的多路复用
HTTP/2引入了多路复用技术，允许在同一连接上并行发送多个请求和响应，而不需要等待前一个请求完成。这种方式显著提高了性能，减少了延迟。
### 3. HTTP/2多路复用实现
```go
package main

import (
    "fmt"
    "net/http"
    "golang.org/x/net/http2"
)

func main() {
    server := &http.Server{
        Addr: ":8080",
    }
    
    http2.ConfigureServer(server, nil) // 配置HTTP/2支持

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, HTTP/2!")
    })

    server.ListenAndServeTLS("server.crt", "server.key") // 启动HTTPS服务器
}
```
## 十四、TCP与UDP的区别，UDP优点，适用场景
### 1. 概念
TCP和UDP是两种不同的网络传输协议，它们之间的主要区别是：
* 特性	TCP	UDP
* 连接性	面向连接	无连接
* 可靠性	提供可靠的数据传输，保证顺序	不保证数据的可靠性和顺序
* 速度	较慢，因需建立连接和确认	较快，因无连接建立和确认
* 流量控制	有流量控制机制	无流量控制
* 适用场景	适用于需要可靠传输的应用	适用于实时性要求高的应用
## 十五、死锁条件及其避免
### 1. 概念
死锁是指两个或多个进程在执行过程中，因为**争夺资源而造成的一种互相等待的状态**，导致所有进程都无法继续执行。
### 2. 四个必要条件
* 互斥条件：至少有一个资源是非共享的，即某一时刻只能被一个进程使用。
* 保持并等待条件：一个进程持有至少一个资源，并等待获取其他资源。
* 不剥夺条件：已经分配给进程的资源在未使用完之前不能被剥夺。
* 循环等待条件：存在一个进程资源的循环等待链。
### 3. 避免死锁
* 资源分配图：使用资源分配图检测可能的死锁情况。
* 避免循环等待：为资源分配设定一个顺序，确保进程按照顺序请求资源。
* 资源预分配：进程在开始时请求所有所需资源，避免在运行过程中请求资源。
* 使用时间限制：如果进程在一定时间内未能获取资源，则放弃并重试。
### 4. 实例
```go
package main

import (
    "fmt"
    "sync"
)

// 可以通过锁的顺序来避免死锁：
var (
    mutexA sync.Mutex
    mutexB sync.Mutex
)

func processA() {
    mutexA.Lock()
    defer mutexA.Unlock()
    mutexB.Lock()
    defer mutexB.Unlock()
    fmt.Println("Process A completed")
}

func processB() {
    mutexA.Lock()
    defer mutexA.Unlock()
    mutexB.Lock()
    defer mutexB.Unlock()
    fmt.Println("Process B completed")
}
```