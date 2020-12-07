# goroutine 和 通道
* goroutine 和通道 channel,他们支持通信顺序进程(CSP),CSP 是一个并发模式,在不同的执行体(goroutine)之间传递值.但是变量本身局限于单一的执行体.

## 8.1 goroutine
* 在 go中每一个并发执行的活动称为 goroutine.
* 当程序启动时候,只有一个 goroutine 来调用 main 函数.称它为主 goroutine. 新的 goroutine通过 go语句进行创建.
* 语法上,一个 go 语句是在普通函数或者方法调用前面加上 go关键词前缀.
* go语句使函数在一个新建的 goroutine中调用,go语句本身的执行立即完成
```
f() // 调用 f(),等待他的返回
go f() // 新建一个调用 f()的 goroutine,不用等待
```
* 当main函数返回,所有的 goroutine 都暴力地直接终结,然后程序退出

## 8.2 示例: 并发时钟服务器
```
func main()  {
	listener, err := net.Listen("tcp", "localhost:8000") // net.Listener 对象,它在一个网络端口撒花姑娘监听进来的连接.
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept() // 阻塞,直到有连接请求进来.然后返回net.Conn 对象来代表一个连接
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
		go handleConn(conn) // 并发支持给多个连接返回
	}
}

func handleConn(c net.Conn)  {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n")) 
		// net.Conn 满足 io.Writer 接口. 
		// time.Time.Format 方法提供了格式化日期和时间信息的方式.它的参数是一个模板.指示如何格式化一个参考时间,具体如 Mon Jan 2 03:04:05PM 2006 UTC-0700这样的形式
		if err != nil {
			return
		}
		time.Sleep(1* time.Second)
	}
}
```
## 8.3 示例 并发回声服务器
* 真正的并发,不仅可以处理来自多个客户端的连接,还可以包括在一个连接处理中,使用多个 go 关键字
```
func main()  {
	listener, err := net.Listen("tcp", "localhost:8000") // net.Listener 对象,它在一个网络端口撒花姑娘监听进来的连接.
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept() // 阻塞,直到有连接请求进来.然后返回net.Conn 对象来代表一个连接
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func echo(c net.Conn, shout string, delay time.Duration)  {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn)  {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}
```

## 8.4 通道
* goroutine 是go并发的执行体, channel 是它们之间的连接, 通道是可以让一个 goroutine发送特定值到 另一个 goroutine的通信机制. 每个通道都是一个具体类型的导管,叫做通道的元素类型.
* 创建: 
内置make()创建通道
`ch := make(chan int) // ch的元素类型是 chan int, int类型元素的通道`
* 引用类型: 和 map一样,make创建的是一个引用类型,和其他引用类型一样,通道的零值是 nil
* 可比较性 同种类型的通道可以用==来比较,当两者都是同一类型的引用时候,比较结果为 true;通道也可以和 nil进行比较
* 操作: send 发送 receive 接收 统称通信

```
<-操作符
ch <- x // 发送 往通道里发送
y = <-ch // 赋值语句中的接收表达式 接收时候<-ch,可以连接 从通道中接收
<-ch //接收语句,丢弃结果
```

close 关闭操作 内置的 close(ch)函数来关闭通道 关闭后的发送操作将导致宕机. 在一个已经关闭的通道上进行接收操作,将获取所有已经发送的值,直到通道为空.这时候任何接收操作都会立即完成,同时获取到一个通道元素类型对应的零值
* 无缓冲通道和有缓冲通道: 通过 make()创建的第二个参数来指定缓冲通道的容量

```
make 函数初始化变量时候,有:,除非是之前已经声明过变量了
ch1 := make(chan int) // 无缓冲通道
ch2 := make(chan int, 0) // 无缓冲通道
ch3 := make(chan int, 3) // 容量为3的缓冲通道
```
### 8.4.1 无缓冲通道

* 无缓冲通道(同步通道): 发送和接收同步化,发送阻塞,直到另一个 goroutine 在对应的通道上执行接收操作,这时候传送完成,两个 goroutine可以继续执行.
* 通道可以用来传递数据,消息的值,这种消息称为事件,如果没有任何额外信息,单纯就是进行同步,可以用 struct{} bool int 来做通道的元素类型
* 每一个消息有一个值,但有时候通信本身以及通信发生的时间也很重要,. 当我们强调这方面的时候,把消息叫做事件.当事件没有携带额外信息时候,它单纯的目的是进行同步.
我们通过使用一个 struct{}元素类型的通道来强调它.尽管通常使用 bool或 int类型的通道来做同样的事情,因为done<-1比 done<- struct{}{}要短.

### 8.4.2 管道
* 通道可以用来连接 goroutine,这样一个的输出是另一个的输入,这个叫管道(pipeline) 
* 在通道关闭后,任何后续的发送操作都会导致应用崩溃,
* range 语法可以在通道上进行迭代.这个语法更方便接收在通道上发送的值,接受完最后一个值后,关闭循环
* 只有在通知接收方 goroutine所有的数据是否发送完成时候才需要关闭通道, 不是必须的.
* close() 通道的关闭 和 x.Close() 资源的关闭 不要混淆了
* 试图关闭一个已经关闭的通道会导致宕机
`naturals := make(chan int)
close(naturals)
close(naturals)
`
* 接收变种, 产生2个结果,接收到的通道元素,以及一个布尔值,它为 true 的时候,代表接受成功,false表示当前的接受操作在一个关闭的并且读完的通道上.
```
go func () {
    for {
        x,ok := <- naturals  
        if !ok {
            break
        }
        squares <- x*x
    }
    close(squares)
}

```
* range 循环语法在通道上迭代.接收完最后一个值后关闭循环
```
没有 x,_ 里面都是通道元素
for x := range naturals {
    squares <- x * x 
}
```
* 结束时候,关闭每一个通道不是必需的.只有在通知接收方 gorountine 所有的数据都发送完毕时候才需要关闭通道.通道也是可以通过垃圾回收器根据它是否可以访问来决定是否回收它,而不是根据它是否关闭
(不要将这个 close()操作和对于文件的 close 操作混淆,当结束的时候对每一个文件调用 Close 方法是非常重要的)
* 试图关闭一个已经关闭的通道会导致宕机,就像关闭一个空通道一样.


### 8.4.3 单向通道类型
* 函数的形参,总是有意的限制不能发送或者接受
* out chan<-  int 发送通道  in <-chan int 接收通道
* 双向可以转单向(隐式),但是逆向操作不 ok 
```
func counter(out chan<- int)  {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}
func printer(in <-chan int)  {
	for v := range in {
		fmt.Println(v)
	}
}
```

### 8.4.4 缓冲通道
* make()第二个参数指定缓冲通道的容量
* cap(ch) 内置 cap函数获取通道的容量
* len(ch) 获取当前通道的元素个数(debug 性能优化时候会用到)
* 队列的使用:使用 slice 创建就可以了,不要用缓冲通道来实现,因为如果没有另一个 goroutine 从通道进行接收,那么发送者有被永久阻塞的风险
* 通道和 goroutine的调度深度关联,如果没有另一个 goroutine从通道进行接收,发送者(也许是整个程序)有被永久阻塞的风险.
* goroutine 泄露, 发送响应结果到通道的时候没有 goroutine来接收,泄露的 goroutine 是不会自动回收的,是 bug,不像回收变量,泄露的 goroutine 不会自动回收,所以确保 goroutine在不需要的时候可以自动结束
```
// 并发的像3个镜像(指相同但是分布在不同地理区域的服务器)地址发请求,它将他们的响应通过一个缓冲通道进行发送,然后只接收第一个返回的响应,因为它是最早到达的
func mirroedQuery() string {
    responses := make(chan string, 3)
    go func () { responses <- request("asia.gopl.io")} ()
    go func () { responses <- request("europe.gopl.io")} ()
    go func () { responses <- request("americas.gopl.io")} ()
    return <-responses
}
```
* 无缓冲通道和缓冲通道的选择: 无缓冲通道提供强同步保障, 缓存通道 发送和接受是解耦的,一般会创建一个容量是使用上限的缓冲通道.

## 8.5 并行循环
* 处理文件的顺序没有关系,因为每一个缩放操作和其他的操作独立.像这样由一些完全独立的子问题组成的问题称为高度并行.
* 计数器类型 sync.WaitGroup,它可以被多个 goroutine安全的操作, wd.Add(1); wd.Done() == wd.Add(-1); wd.Wait()
```
func makeThumbnails6(filenames <-chan string) int64  {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for f := range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			thumb, err  := thumbnail.ImageFile(f)
			if err != nil {
				log.Print(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}

	go func() {
		wg.Wait()
		close(sizes)
	}()
	var total int64
	for size := range sizes { // 主 goroutine把大部分时间花在range循环休眠上了,等待工作者发送值或等待closer()关闭通道
		total += size
	}
	return total
}
```

## 8.6 示例 并发的 web爬虫
* 程序的并行度太高了,无限制的并行通常不是一个好的主意,因为系统总有各种限制因素.解决方法是根据资源可用情况限制并发个数,以匹配合适的并行度.
```
func main()  {
	worklist := make(chan []string)

	go func() {
		worklist <- os.Args[1:]
	}()

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) { // 循环变量捕获问题,显式传参
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
```
* 缓冲通道来建立一个并发原语,称为计数信号量.通道的元素类型不重要,使用 struct{},因为它的空间大小是0
```
func main()  {
	worklist := make(chan []string)
	// 计数器 n跟踪发送到任务列表的个数
	var n int
	n++
	go func() {
		worklist <- os.Args[1:]
	}()
	seen := make(map[string]bool)
	for ; n >0; n-- { // 如果第一个条件为空,那么就是;
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}

		}
	}
}

var token = make(chan struct{}, 20)
func crawl(url string) []string {
	fmt.Println(url)
	token <- struct{}{} // 获取令牌
	list, err := links.Extract(url)
	<-token
	if err != nil {
		log.Print(err)
	}
	return list
}
```
* 死锁:死锁是一种卡住的情况,其中主 goroutine 和爬取 goroutine同时发送给对方但是双方都没有接收
```
// 发送给任务列表的命令行参数必须在它自己的goroutine中运行来避免死锁.
go func(link string) {
    worklist <- crawl(link)
}(link)
```
* 没有计数信号量,通过长期存活的20个 goroutine来调用它,这样确保最多20个HTTP请求并发执行
```
func main()  {
	worklist := make(chan []string) // 可能有重复的 url列表
	unseenLinks := make(chan string) // 去重后的url列表

	go func() {
		worklist <- os.Args[1:]
	}()

	for i := 0; i < 20; i++  {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				// 为啥这里要加 go呢? 如果不加 go,这里进行发送,但是 worklist 是满的,那么这里就会阻塞
				go func() {
					worklist <- foundLinks // crawl发现的链接通过精心设计的 goroutine发送到任务列表来避免死锁
				}()
			}
		}()
	}

	// 爬取 goroutine 使用同一个通道 unseenLinks进行接收,主 goroutine 负责从任务列表接收到的条目进行去重,然后发送每一条
	// 没有爬去过的 url到 unseenLinks 通道,然后被爬取goroutine接收
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}

}

func crawl(url string) []string  {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		fmt.Print(err)
	}
	return list
}
```

## 8.7 使用 select多路复用
*  time.Tick 返回一个通道,它定期发送事件,像一个节拍器一样.每个事件的值是一个时间戳
```
func main()  {
	fmt.Println("Commencing countdown.")
	// time.Tick 返回一个通道,它定期发送事件,像一个节拍器一样.每个事件的值是一个时间戳
	tick := time.Tick(2*time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	launch()
}
func launch() {
	fmt.Println("Lift off!")
}
```
* select 语句:每个情况指定一次通信(在一些通道上进行发送或接收操作)和关联的一块代码, 第二种情况短变量声明,可以引用接收的值
select 一直在等待,直到一次通信告知有一些情况可以执行,然后它进行这次通信,执行此情况所对应的语句,其他的通信将不会发生.
* 对于没有对应情况的 select{},select{}将会永远等待
* 
```
select {
case <-ch1:
//...
case x := <-ch2:
// use x...
case ch3 <- y:
// ...
default:
// ...
}
```
* time.After 在立即返回一个通道,然后启动一个新的goroutine在间隔指定时间之后,发送一个值到它上面
```
func main()  {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // 读取单个字节
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown. pleas return to abort")
	select {
	// time.After 在立即返回一个通道,然后启动一个新的goroutine在间隔指定时间之后,发送一个值到它上面
	case <- time.After(10*time.Second):
		//不执行任何操作
	case <- abort:
		fmt.Println("Launch aborted!")
		return
	}
	launch()
}
func launch() {
	fmt.Println("Lift off!")
}
```
* ch 的缓冲区大小是1,他要么是空,要么是满,因此只有在其他的一种情况下执行.要么 i 是偶数时候发送,要么 i是奇数时候接收
```
// ch 的缓冲区大小是1,他要么是空,要么是满,因此只有在其他的一种情况下执行.要么 i 是偶数时候发送,要么 i是奇数时候接收
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case ch <- i:
			//
		case x := <-ch:
			fmt.Println(x)
		}

	}
```
* 如果多种情况同时满足,select会随机选择一个,这样保证每个一个通道都有相同的机会被选中.
```
ch := make(chan int, 4) 结果随机
```
* time.Tick()函数行为很像创建一个 goroutine在循环里调用 time.Sleep(),然后在每次醒来时候发送事件.当上面的倒计时函数返回时候,它停止从tick通道接收事件,
但是计时器 goroutine还在运行,徒劳的向一个没有goroutine接收的通道不断发送,会导致goroutine 泄露
```
func main() {
	// 创建终止通道
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown. pleas return to abort")
	tick := time.Tick(1*time.Second)
	for countdown := 0; countdown < 10; countdown++ {
		fmt.Println(countdown)
		select {
		case <-tick:
			//
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
}
func launch() {
	fmt.Println("Lift off!")
}
```
* time.Tick()仅仅在应用的整个生命周期都需要时才合适,否则我们需要使用这个模式:
```
ticker := time.NewTicker(1*time.Second)
<-ticker.C // 从 ticker 通道接收
ticker.Stop() // 造成 ticker的 goroutine 终止
```
* 有时候我们试图在一个通道上发送或接收,但是不想在通道没有准备好的情况下被阻塞---非阻塞通信,这时候使用 select语句也可以做到.
select可以有一个默认情况,它用来指定在没有其他的通信发生时可以立即执行的动作
```
// 尝试从 abort通道中接收一个值,如果没有值,它什么都不做.这是一个非阻塞操作,重复这个动作称为对通道轮询.
    abort := make(chan struct{})
	select {
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	default:
		//fmt.Println("default")
	}
	fmt.Println("end") // 跑一次的结果是:end,如果想一直监听,需要循环
```
* 通道的零值是 nil,有时候 nil通道很有用.因为在 nil通道上发送和接收将永远阻塞,
对于 select 的情况,如果其通道是 nil,它将永远不会被选择,这可以让我们用 nil来开启或禁用特性所对应的情况,比如超时处理或取消操作,响应其他的输入事件或发送事件

## 8.8 示例 并发目录遍历
* 见 dux包 代码

## 8.9 取消
* 有时候我们需要让一个 goroutine 停止它当前的任务
* 通常任何时刻,很难知道有多少 goroutine正在工作
* 对于取消操作,我们需要一个可靠的机制在一个通道上广播一个事件,这样很多 goroutine可以认为它发生了,然后可以看到它已经发生了
* 当一个通道关闭而且已经取完所有发送的值之后,接下来的接收操作立即返回,得到零值,我们可以利用它创建一个广播机制:不在通道上发送值,而是关闭它
```
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
//go run du4/main.go -v / /usr /bin /etc
var verbose = flag.Bool("v", false, "show verbose progress messages")
func main()  {
	// 确定初始目录
	flag.Parse()
	roots := flag.Args()
	//os.Args[1:]
	if len(roots) == 0 {
		roots = []string{"."}
	}
	// 遍历文件树
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes) // 这里是& 指针
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()
	// 当检测到输入的时候,取消遍历
	go func ()  {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	// 定期输出结果
	var tick <-chan time.Time // 变量 tick 类型是单向(接收)通道 类型是 time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	var nfiles, nbytes int64
loop: // 只有 loop:
	for {
		select {
		case <-done:
			// 耗尽fileSize 通道,丢弃他的所有值,直到通道关闭,这么做是为了保证所有的walkDir调用可以执行完,不会卡在向 fileSizes通道发送消息上
			for range fileSizes {
				// 不执行任何操作
			}
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loop // 标签化的 break语句,将跳出 select 和 for 循环,没有标签的 break, 只能跳出 select,导致循环的下一次迭代. fileSizes 关闭,跳出
			}
			nfiles++
			nbytes += size
		case <-tick :
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfile, nbytes int64)  {
	fmt.Printf("%d files %.1fGB\n", nfile, float64(nbytes)/1e9)
}

func walkDir(dir string, n *sync.WaitGroup,fileSizes chan<- int64 )  {
	defer n.Done()
	// 轮询取消状态,如果设置状态,什么都不做,立马返回
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1) // 这里调用了也不能少
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes) // 这里也有 go
		} else {
			fileSizes <- entry.Size() // 文件,发送一条消息到通道,消息是文件所占的字节数
		}
	}
}
var sema = make(chan struct{}, 20)
// 高峰期会有很多很多 goroutine,
func dirents(dir string) []os.FileInfo {
	//sema <- struct{}{}
	select {
	case sema <- struct{}{}: // 获取令牌
	case <-done:
		return nil
	}
	defer func() { // 释放令牌
		<-sema
	}()
	// defer <-sema 错误写法
	entries, err := ioutil.ReadDir(dir) // ioutil.ReadDir()返回一个os.FileInfo类型的slice.针对单个文件同样的信息可以通过调用 os.Stat()来返回.
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
```

## 8.10 聊天服务器
* 



