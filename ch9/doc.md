# 9 使用共享变量实现并发

## 9.1 静态
* 如果我们无法自信的说一个事件肯定先于另一个事件,那么这两个事件就是并发
* 考虑一个能在串行程序中正确工作的函数,如果这个函数在并发调用时仍能正确工作,那么这个函数是并发安全的.
在这里并发调用是指,在没有额外同步机制的情况下,从两个或多个 goroutine同时调用这个函数.如果一个类型的所有可访问方法和操作都是并发安全的时候,则它称为并发安全的类型
* 让一个程序并发安全的并不需要其中的每一个具体类型都是并发安全的.实际上,并发安全的类型其实是特例而不是普遍存在的,所以仅在文档指出类型是安全的情况下,才可以并发地访问一个变量.
对于觉大部分变量,如要回避并发访问,
    1. 要么限制变量只存在于一个 goroutine中,
    2. 要么维护一个更高层的互斥不变量
* 导出的包级别函数通常可以被认为是并发安全的.因为包级别的变量无法限制在一个 goroutine 内,所以那些修改这些变量的函数就必须采用互斥机制
* 函数并发调用不工作的原因有很多,包括死锁, 活锁(比如多个线程在尝试绕开死锁,却由于过分同步导致反复冲突), 以及资源耗尽
* 竞态: 指多个 goroutine按某些交错顺序执行时程序无法给出正确结果. 竞态对于程序是致命的,因为它可能潜伏在程序内,出现频率也很低,有可能仅在高负载环境或者使用特定的编译器,平台 架构时才出现.这些都让竞态很难再现和分析
* 数据竞态: 竞态的一种, 数据竞态发生于两个 goroutine 并发读写同一个变量并且至少其中一个时写入时.
* 当发生数据竞态的变量类型是大于一个机器字长的类型,比如接口 字符串 slice,事情就更复杂了
```
	var x []int
	go func() {
		x = make([]int, 10)
	}()

	go func() {
		x = make([]int, 1000000)
	}()
	// 雷区:未定义行为
	// x可能是 nil,一个长度为10的 slice 或一个长度为100000的 slice. slice:指针 长度 容量, 指针时第一个 make 调用而长度来自第二个 make 调用,x是嵌合体,
	// 它名义上长度为100000,但是底层的数组只有10个元素,在这种情况下,尝试存储到第99999个元素,会伤及一段很遥远的内存,其恶果无法预测看,问题也很难调试和定位.雷区:未定义行为
	x[99999] = 1 // 未定义行为,可能造成内存异常
```
* 避免数据竞态,3种方法:
1. 不要修改变量, 一旦完成了初始化,icons 就不再修改,那些从不修改的数据结构以及不可变数据结构本质上时并发安全的,也不需要做任何同步
```
// 下面有延迟初始化,对于每一个键,在第一次访问时才触发加载. 并行调用,Icon 的结果
var icons = make(map[string]image.Image)
func Icon(name string) image.Image  {
	icon, ok := icons[name]
	if !ok {
		icon = loadIcon(name)
		icons[name] = icon
	}
	return icon
}
func loadIcon(name string) image.Image  {
	// ...
}
// 如果在创建其他 goroutine之前就用完整的数据来初始化 map,并且不再修改.那么无论多少 goroutine也可以安全地并发调用 Icon,因为每个 goroutine都只读取这个 map
var icons = map[string]image.Image{
	"spades.png" : loadIcon("spades.png"),
	"hearts.png" : loadIcon("hearts.png"),
	"diamonds.png" : loadIcon("diamonds.png"),
	"clubs.png" : loadIcon("clubs.png"),
}

func Icon(name string) image.Image {
	return icons[name]
}
```
2. 避免多个 goroutine访问同一个变量.  把变量限制在一个 goroutine 内部
由于其他 goroutine无法直接访问相关变量,因此它们就必须使用通道来向受限 goroutine 发送查询请求或更新变量.
    * 不要通过共享内存来通信,而应该通过通信来共享内存.
    * 使用通道请求来代理一个受限变量的所有访问的 goroutine称为该变量的监控 goroutine.比如 broadcaster goroutine 监控了对 clients map 的访问
```
package bank

var deposits = make(chan int) // 发送存款额
var balances = make(chan int) // 接收余额

func Deposit(amount int) {
	deposits <- amount
}

func balance() int {
	return <-balances
}

func teller() {
	var balance int // balance 被限制在 teller goroutine 中
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:

		}
	}
}

func init() {
	go teller() // 启动监控 goroutine
}

```
* 串行受限: 即使一个变量无法在整个生命周期受限于单个 goroutine,加以限制仍然可以是解决并发访问的好办法.比如:可以通过借助通道来把共享变量的地址从上一步传到下一步,从而在流水线上的多个 goroutine之间共享该变量.
在流水线的每一步,在把变量地址传给下一步后就不再访问该变量了,这样对该变量的访问是串行的.换个说法,这个变量受限于流水线的下一步,再受限于下一步.以此类推,这种受限称为:串行受限
```
Cake 串行受限
type Cake struct {
	state string
}

func baker(cooked chan<- *Cake)  {
	for {
		cake := new(Cake)
		cake.state = "cooked"
		cooked <- cake // baker 不再访问cake变量
	}
}

func icer(iced chan<- *Cake, cooked <-chan *Cake)  {
	for cake := range cooked {
		cake.state = "iced"
		iced <- cake // icer 不再访问 cake变量
	}
}
```
* 第三种:避免数据竞态的方法是允许多个 goroutine 访问同一个变量,但在同一个时间内只有一个 goroutine访问.这种方法称为互斥机制.

## 9.2 互斥锁: sync.Mutex
* 使用一个缓冲通道实现了计数信号量,用于确认同时发起 HTTP请求的 goroutine数量不超过20个,同样,也可以用一个容量为1的通道来保证同一时间最多有一个 goroutine 能访问共享变量.
* 一个计数上限为1的信号量称为二进制信号量.
* 互斥锁模式
```
var (
	sema = make(chan struct{}, 1) // 用来保护 balance 的二进制信号量
	balance int
)
func Desposit(amount int)  {
	sema <- struct{}{} //获取令牌
	balance += amount
	<-sema // 释放令牌
}
func Balance() int {
	sema <- struct{}{} //获取令牌
	b := balance
	<-sema // 释放令牌
	return  b
}
```
* sync.Mutex 类型,Lock 方法用于获取令牌(token, 此过程称为上锁),Unlock方法用于释放令牌
```
var (
	mu  sync.Mutex
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance += amount
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}
```
* 每一个 goroutine 在访问银行的变量balance 之前,他都必须先调用互斥量的 Lock方法来获取一个互斥锁.如果其他 goroutine 已经取走了互斥锁,那么操作会一直阻塞到其他 goroutine 调用Unlock 之后(此时互斥锁再度可用)
* 互斥量保护共享变量.按照惯例,被互斥量保护的变量声明应当紧接在互斥量的声明之后.如果实际情况不是如此,请确认已加了注释来说明此事
* 在 Lock和 Unlock之间的代码,可以自由的读取和修改共享变量,这一部分称为临界区域.在锁的持有人调用 Unlock 之前,其他 goroutine不能获取锁.所以 goroutine 在使用完后就应当释放锁,另外,需啊包括函数的所有分支,特别是错误分支.
* 这种函数 互斥锁 变量的组合方式称为监控模式.
* defer mu.Unlock() 通过延迟执行 Unlock 就可以把临界区域隐式扩展到当前函数的结尾,避免了必须在一个或者多个远离 Lock 的位置插入一条条 Unlock 语句
```
mu.Lock()
defer mu.Unlock()
return xxx
```
* defer的执行成本比显式调用 Unlock 略大,但是不足以成为代码不清晰的理由.在处理并发程序时,永远应该优先考虑清晰度,并且拒绝过早优化,在可以使用的地方,就尽量使用 defer 来让临界区域扩展到函数结尾处
* 死锁, go语言的互斥量是不可再入的.
```
// 不是原子操作, 3个串行操作,每个操作都申请释放了互斥锁,但是对于整个序列没有上锁
func Withdraw(amount int) bool {
	Deposit(-amount)
	if Balance() < 0 {
		Deposit(amount)
		return false
	}
	return true
}

func Withdraw2(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	Deposit(-amount) // 会导致函数里再次取获取锁,死锁,Withdraw2会一直卡住
	if Balance() < 0 {
		Deposit(amount)
		return false
	}
	return true
}
```
* 常见的解决方案是吧 Deposit 这样的函数拆分为两部分:一个不导出的函数 desposite,它将定已经获得互斥锁,并完成实际的业务逻辑.以及一个导出的函数 Deposit,它用来获取锁并调用 deposit.
* 封装即通过在程序中减少对数据结构的非预期交互,来帮助我们保证数据结构的不变量.因为类似原因,封装也可以用来保持并发中的额不变性.所以无论是为了保护包级别的变量,还是结构中的字段,当你使用一个互斥量时,都请确保互斥量本身以及被保护的变量都没有导出.

## 9.3 读写互斥锁: sync.RWMutex
* 只允许读操作可以并发执行,但写操作需要获得完全独享的访问权限.多读单写锁.sync.RWMutex
```
var mu sync.RWMutex
var balance int
func Balance() int {
    mu.RLock() // 读锁
    defer mu.RUnlock()
    return balance
}
```
* 读锁 mu.RLock() 共享锁
* 写锁 mu.Lock() 互斥锁
* RLock 仅可用在临界区域内对共享变量无写操作的情形.如果是有更新,那么请用独享版本 Lock
* 仅在绝大部分goroutine都在获取读锁并且竞争激烈的时候(即 goroutine 需要等待后才能获取锁),RWMutex 才有优势,因为RWMutex需要更复杂的内部薄记工作,所以在竞争不激烈的时候,它比普通的互斥锁更慢.
* 应该是可以多个读操作进行,但是上锁了,就只能读不能写.

## 9.4 内存同步
* Balance()需要互斥锁的原因有:
    1. 防止 Balance()插到其他操作中去.
    2. 同步不仅涉及多个 goroutine 的执行顺序问题.同步还会影响内存
* 现代计算机多个处理器.每个处理器都有内存的本地缓存.为了提高效率,对内存的写入是缓存在每个处理器中的.只有必要时候才刷回内存,甚至刷回内存的顺序都可能和
goroutine 的写入顺序不一致.像通道通信或者互斥锁的操作这样的同步原语都会导致处理器把累积的写操作刷回内存并提交.所以这个时刻之前的goroutine的执行结果
就保证了对运行在其他处理器的goroutine可见.
* 特定的编译器 CPU或其他情况下,会有异常
```
var x, y int
go func () {
    x = 1
    fmt.Print("y:", y, " ") //A
}()
go func () {
    y = 1
    fmt.Print("x:", x, " ") //B
}()

x:0 y:0
y:0 x:0
```
1. A 读到 B 的过期值,B 读到 A的过期值.
2. 编译器:赋值和 Print对应不同的变量,所以编译器就会认为两个执行语句的顺序不会影响结果,然后就交换了这个两个语句的执行顺序.
3. CPU:如果两个 goroutine在不同的CPU上执行.每个 CPU都有自己的缓存,那么一个goroutine在写入操作同步到内存之前对另一个goroutine的 Print是不可见的.
* 这些并发问题,都可以通过简单 成熟的模式来避免.即在可能的情况下,将变量限制在单个 goroutine内,对于其他变量,使用互斥锁

## 9.5延迟初始化
* 延迟一个昂贵的初始化步骤到有实际需求的时刻是一个很好的实践.
* 预先初始化一个变量会增加程序的启动延时,并且如果实际执行有可能根本用到这个变量,那么初始化也不是必须的.
* 并发调用 Icon 是不安全的,你认为: 竞态带来的最严重的问题是loadIcons函数会被调用多次.当第一个 goroutine正忙着加载图标的时候,其他 goroutine进入 Icon函数,会发现
icons 仍然是 nil,所以仍然会调用loadIcons
```
var icons map[string]image.Image

func loadIcons()  {
	icons = map[string]image.Image{
		"spades.png" : Icon("spades.png"),
		"hearts.png" : Icon("hearts.png"),
		"diamonds.png" : Icon("diamonds.png"),
		"clubs.png" : Icon("clubs.png"),
	}
}
// 并发不安全
func Icon(name string) image.Image {
	if icons[name] == nil {
		loadIcons() // 一次性初始化
	}
	return icons[name]
}
```
* 但是这个直觉是错的.在缺乏显式同步的情况下,编译器和 cpu在能保证每个 goroutine都满足串行一致性的基础上,可以自由的重排访问内存的顺序.
填充数据之前将一个空的 map 赋值给 icons,因此一个 goroutine 发现icons不是 nil并意味着变量的初始化已经完成了
```
func loadIcons()  {
    icons = make(map[string]image.Image) // 填充数据之前将一个空的 map 赋值给 icons,
	icons["spades.png"] = loadIcon("spades.png")
	icons["hearts.png"] = loadIcon("hearts.png")
	icons["diamonds.png"] = loadIcon("diamonds.png")
	icons["clubs.png"] = loadIcon("clubs.png")
}
```
* 关于并发的直觉都不可靠
* 保证所有的goroutine都能观察到loadIcons效果最简单的正确方法就是使用互斥锁来做同步
```
var mu sync.Mutex
var icons map[string]image.Image
// 并发安全
func Icon(name string) image.Image {
    mu.Lock()
    defer mu.Unlock()
    if icons[name] == nil {
        loadIcons()
    }
    return icons[name]
}
```
* 采用互斥锁访问 icons的额外代价是两个 goroutine不能并发的访问这个变量了,即使在变量已经安全的完成初始化而且不再更改的情况下,也会造成这个后果,
可以用并发读的锁来改善这个问题:
```
var mu sync.RWMutex
var icons map[string]image.Image
func Icon(name string) image.Image {
    mu.RLock()
    if icons != nil {
        icon := icons[name]
        mu.RUnlock()
        return icon
    }
    mu.RUnlock()
    // 由于不先释放一个共享锁就无法直接把他升级到互斥锁,为了避免在过渡期其他 goroutine已经初始化了 icons,所以我们必须再次检查icons的值
    mu.Lock()
    if icons == nil { // 注意:必须重新检查 icons的值,有可能这个锁获取到之前,锁被其他写锁获取过去,已经初始化了
        loadIcons()
    }
    icon := icons[name]
    mu.Unlock()
    return icon
}
```
* 由于不先释放一个共享锁就无法直接把他升级到互斥锁,为了避免在过渡期其他 goroutine已经初始化了 icons,所以我们必须再次检查icons的值

* sync包提供了针对一次性初始化问题的特化解决方案:sync.Once
* Once包含了一个布尔值和一个互斥量,布尔值记录初始化是否已经完成,互斥量则负责保护这个布尔变量和客户端的数据结构
* Once 唯一的方法Do以初始化函数作为它的参数. 每次调用Do(loadIcons) 都会先锁定互斥量,然后检查里边的布尔变量.
在第一次调用的时候,这个布尔值是 false,Do 会调用loadIcons然后把布尔值设为 true.后续的调用相当于空操作,只是通过互斥量的同步来保证所有的 loadIcons
对内存产生的效果(icons变量)对所有的 goroutine可见.以这种方式来使用 sync.Once,避免了变量在正确构造之前就被其他 goroutine 分享
```
var loadIconsOnce sync.Once
var icons map[string]image.Image
// 并发安全
func Icon(name) image.Image {
    loadIconsOnce.Do(loadIcons)
    return icons[name]
}
``` 

## 9.6 竞态检测器
* 动态分析工具: 竞态检测器
* 简单的将-race命令行参数加到go build go run go test 命令里即可使用该功能.
* 构建一个修改后的版本,记录执行时对于共享变量的访问,读写这些变量的 goroutine 的标识.记录所有的同步事件,包括 go语句,通道操作,(*sync.Mutex).Lock 调用,(*sync.WaitGroup).Wait 的调用
* 竞态检测器报告了所有的实际运行了的数据竞态.

## 9.7 并发非阻塞缓存
* 函数记忆问题:缓存函数调用的结果,达到多次调用只需要计算一次的效果
* 解决方案:并发是安全的,并且要避免简单的对整个缓存使用单个锁带来的锁争夺问题
* 重复抑制:两个 goroutine同时先查缓存,发现缓存中没有需要的数据,然后调用那个慢函数f,最后又都用获得的结果来更新map,其中一个结果会被另一个给覆盖掉
要避免这种情况发生.
* 构建并发结构:共享变量上锁,通信顺序进程

## 9.8 goroutine 和线程
* 差异:量变
### 9.8.1 可增长的栈
* 每个 os线程都有一个固定大小的栈内存(2MB),保存在其他函数调用期间正在执行或临时暂停的函数中的局部变量, 既太大又太小,
* 递归复杂的程序又太小,固定的栈始终是不够大.
* goroutine最开始只是一个很小的栈,2kb,不固定,按需增大或缩小.最大可达到1GB

### 9.8.2 goroutine 调度
* os 线程由 os内核来调度. 调度器的内核函数. 更新调度器的数据结构
* go 运行时包含自己的调度器.m:n 的调度技术.可以复用/调度m个 goroutine到 n个线程
* go 调度器不是由硬件时钟来定时触发的,而是由特定的 go语言结构触发的.比如当一个 goroutine调用 time.Sleep或通道阻塞或对互斥量操作时,
调度器就会把这个 goroutine设为休眠模式.并运行其他的goroutine直到前一个可重新唤醒为止,因为它不需要切换到内核语境,所以调用一个 goroutine比调度一个线程成本低很多

### 9.8.3 GOMAXPROCS
* go调度器使用 GOMAXPROCS参数来确定需要使用多少个os线程来同时执行go代码,默认是机器上的CPU数量,GOMAXPROCS是 m:n 中的 n
* 可以用环境变量GOMAXPROCS或者 runtime.GOMAXPROCS函数来显式控制这个参数,
```
GOMAXPROCS=1 go run xxx.go
GOMAXPROCS=2 go run xxx.go
```
### 9.8.2 goroutine 没有标识
* 其他支持多线程的语言或者操作系统里,当前线程都有一个独特的标识,它通常可以取一个整数或者指针,根据这个可以构建线程的局部存储,
本质上就是全局的 map,以线程的标识作为键,这样每个线程都可以独立的用这个 map存储和获取值,不受其他线程干扰
* 不健康的超距作用:函数的行为不仅取决于它的参数,还取决于他的线程标识,
* go 能影响一个函数行为的参数应该是显式指定的.不仅程序易读,还可以自由的把一个函数的子任务分发到多个不同的goroutine中而无须担心这些 goroutine标识
``


