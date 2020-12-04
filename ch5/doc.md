# 5. 函数
* 函数包含连续执行的语句,函数对它的使用者隐藏了实现细节
## 5.1 函数声明
* 每一个函数声明都包含一个名字 一个形参列表 一个可选的返回值列表以及函数体
```
func name(parameter-list) (resutl-list) {
    body
}
```
* 形参列表指定了一组变量的参数名和参数类型,这些局部变量由调用者提供的实参传递而来.返回列表指定了函数返回值的类型.当函数返回一个未命名的返回值或没有返回值的时候,返回列表的圆括号可以省略.如果一个函数即省略返回列表也没有任何返回值,那么设计这个函数的目的是调用函数之后所带来的附加效果
* 返回值也可以像形参一样命名.这个时候,每一个命名的返回值会声明为一个局部变量,并根据变量类型初始化为相应的零值
* 当函数存在返回列表时,必须显式的的以 return语句结束,除非函数明确不会走完整个执行流程,比如在函数中抛出宕机异常或者函数体内存在一个没有 break退出条件的无限 for循环
* 如果几个形参或者返回值类型相同,那么类型只需要写一次.
```
// 空白标识符用来强调这个形参在函数中未使用
func add(x int, y int) int {return x+y}
func sub(x, y int) (z int) {z = x -y ; return }
func first(x int, _ int) int {return x}
func zero(int, int) int {return 0} 
fmt.Printf("%[1]T\n", add) 都是 func(int, int) int
```
* 函数的类型称为函数签名.当两个函数拥有相同的形参列表和返回值列表时,认为这个两个函数的类型或签名是相同的.而形参和返回值的名字不会影响到函数类型,采用简写同样也不会影响到函数的类型.
* 每一次调用函数都需要提供实参来对应函数的每一个形参,包括参数的调用顺序也必须一致.go 语言没有默认参数值的概念,(a = true )也不能指定实参名.
* 形参变量都是函数的局部变量,初始值由调用者提供的实参传递,函数形参以及命名返回值同属于函数最外层作用域的局部变量
* 函数的声明没有函数体, 那说明这个函数使用了除了 go 以外的语言实现. 这样的声明定义了该函数的签名
```
package math
func Sin(x float64) float64 // 使用汇编语言
```
## 5.2 递归
* 函数可以递归调用,这意味着函数可以直接或间接的调用自己.
* golang.org/x/net/html html.Parse() 读入一段字节序列,解析它们,然后返回 HTML 文档树的根节点 html.Node.
* go 语言实现了可变长度的栈,栈的大小会随着使用而增长,可达到1GB左右的上限,这使得我们可以安心的使用递归,而不用担心溢出的问题

## 5.3 多返回值
* 一个函数不止能返回一个结果.返回一个期望得到的结果和一个错误值或者一个表示函数调用是否正确的布尔值.
* fmt.Errorf() 格式化处理过的附加上下文信息. `return nil, fmt.Errorf("xxx %v", err)`
* 显式的将多个返回值赋值给变量, 如果要忽略某个一个返回值,可以将它赋值给一个空标识符
```
links, err := findLinks(url)
links, _  := findLinks(url) //
```
* 返回一个多值结果可以是调用另一个多值返回的函数. `return xxx()`
* 传递到拥有多个形参的函数中,方便调试 fmt.Println(xxx())
* 一个函数如果有命名的返回值.可以省略return 语句的操作数,称为裸返回. (word string)  return 
* 裸返回是将每个命名返回结果按照顺序返回的快捷方法.
* 裸返回可以消除重复代码,但是不能使代码易理解.保守使用裸返回

## 5.4 错误
* 如果当函数调用发生错误时返回一个附加的结果作为错误值,习惯上将错误值作为最后一个结果返回,如果错误只有一种情况,结果通常设置未bool 值,
* i/o操作,错误的原因可能多种多样.调用者需要详细的信息,这种情况下错误类型是:error
* error是内置的接口类型.一个错误可能是空值和非空值,空值意味着成功,而非空值意味着失败.而且非空的错误类型有一个错误消息字符串.
* go 使用普通的值而非异常来报告错误

### 5.4.1 错误处理策略
* 1.将错误传递下去
* 2.构建一个新的错误,添加其他有用的错误信息

* 对于不固定或不可预测的错误,在短暂间隔之后,对操作进行重试是合理的.超出一定的重试次数和限定的时间后再报错退出
* 指数退避策略
```
    const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tires := 0; time.Now().Before(deadline); tires++ {
		log.Println(tires, time.Second, time.Second << uint(tires))
		time.Sleep(time.Second << uint(tires)) // 指数退避策略
		//2020/12/02 12:25:52 0 1s 1s
		//2020/12/02 12:25:53 1 1s 2s
		//2020/12/02 12:25:55 2 1s 4s
		//2020/12/02 12:25:59 3 1s 8s
		//2020/12/02 12:26:07 4 1s 16s

	}
```
* fmt.Errorf() 输出 error类型的错误, log.Fatalf()输出日志,会将时间和日期作为前缀添加到错误信息前
* log包的函数都会以日期和时间作为前缀输出.所有的 log函数都会为缺少换行符的日志补充一个换行符  使用 log.SetPrefix() log.SetFlags()可以自定义名称作为 log包的前缀.并将日期和时间略去
```
log.SetPrefix("wait...")
log.SetFlags(0)
log.Println(tires, time.Second, time.Second << uint(tires)) 
log.Printf("ping failed: %v; networking disabled", err) //记录日志,然后程序继续运行
log.Fprintf(os.Stderr, "xxx %v", err) // 直接输出到标准错误流中
```
* go进行错误检查之后,检测到失败的情况往往都在成功之前.如果检测到的失败导致函数返回,成功的逻辑一般不会放在 else块,而是在外层作用域中. 函数会有一个通常的形式,就是在开头会有一连串的检查用来返回错误.之后 跟着实际的函数体一直到最后.

### 5.4.2 文件结束标识符
* io.EOF 文件结束标识符,如果文件读取失败,而且错误==io.EOF 代表文件读取完啦. 

## 5.5 函数变量
* 函数变量也有类型,而且他们可以赋给变量或传递或从其他函数中返回
* 函数类型不一致,不能赋值
* 函数类型的零值是 nil,调用一个空的函数变量将导致宕机
* 函数类型可以和空值 nil 比较,但是函数本身不可比较,所以不可以互相进行比较或者作为键值出现在 map中
* 函数可以作为参数传递. strings.Map(func, s) 对字符串的每一个字符使用一个函数,将结果连接起来变成另一个字符串

```
	f := square
	fmt.Println(f(3))
	f = negative
	fmt.Println(f(5))
	//f = product // cannot use product (type func(int, int) int) as type func(int) int in assignment

	var fc func(int) int
	fc(3) // panic: runtime error: invalid memory address or nil pointer dereference
	if fc != nil {
		fc(3)
	}
```

* fmt.Printf()来缩进输出. `fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)` %*s中的*号输出带有可变数量空格的字符串,输出的宽度和字符串则由参数 depth*2和""提供

## 5.6 匿名函数
* 命名函数只能在包级别的作用域进行声明,但我们能够使用函数字面量在任何表达式内指定函数变量.
* 函数字面量就像函数声明,但是在 func 关键字之后没有函数名称,而且没有(),它是一个表达式,它的值称为匿名函数
* 函数字面量在我们使用的时候才定义 `strings.Map(func (r rune) rune {return r + 1 }, "HAL-9000)`
* 匿名函数可以获取到整个词法环境,因此里层的函数可以使用外层函数的变量.
* 里层的匿名函数可以获取和更新外层squares函数的局部变量,这些隐藏变量的引用就是函数归类为引用类型而且函数变量无法进行比较的原因
* 函数变量 闭包
* 变量的生命周期不是由其作用域决定的
* 当一个匿名函数需要进行递归,必须线声明一个变量,然后将匿名函数赋给这个变量.如果将这两个变量合成一个声明.函数字面量将不能存在于 visitAll 变量的作用域中,这样也就不能递归调用自己了
```
    var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	// 当一个匿名函数需要进行递归,必须线声明一个变量,然后将匿名函数赋给这个变量.如果将这两个变量合成一个声明.函数字面量将不能存在于 visitAll 变量的作用域中,这样也就不能递归调用自己了
	/**
	visitAll := func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	 */

```
* 网络爬虫的核心是解决图的遍历.拓扑排序的实例展示了深度优先遍历,对于网络爬虫,我们使用广度优先遍历

* 警告捕获迭代变量
```
var rmdirs []func()
func bhbl() {
	for _, d := range tempDirs() {
		dir := d // 在循环体内将循环变量赋给一个新的局部变量 dir.
		os.MkdirAll(dir, 0755)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir)
		})
	}

	for _, dir := range tempDirs() {
		os.MkdirAll(dir, 0755)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir) // 不正确,匿名函数可以获取到外部的变量 dir变量的实际取值是最后一次迭代时的值并且所有的 os.RemoveAll 调用最终都试图删除同一个目录
		})
	}
	//
	for _, rmdir := range rmdirs {
		rmdir()
	}
}
```
* 引入一个内部变量来解决变量迭代导致更新的问题
```
for _, dir := range tempDirs() {
    dir := dir // 声明内部 dir,并以外部 dir初始化
    // ...
}
```
* 这样隐患不仅仅存在于使用 range的 for循环里, for循环捕获索引变量 i也会导致同样的问题
```
var rmdirs []func()
dirs := temDirs()
for i:= 0; i < len(dirs); i++ {
    os.MkdirAll(dirs[i], 0755) //ok
    rmdirs = append(rmdirs, func(){
        os.RemoveAll(dirs[i]) // 不正确
    })
}
```
* 迭代变量捕获的问题是最频繁的:go defer,因为这两个逻辑都会推迟函数的执行时机.直到循环结束.


## 5.7 变长函数
* 变长函数被调用的时候,可以有可变个数的参数个数.fmt.Printf()开头需要提供一个固定的参数,后续可以接受任意数目的参数
* 在参数列表最后的类型名称之前使用省略号...表示一个变长函数,调用这个函数的时候可以传递该类型任意数目的参数
*  vals ...int 在形参里数据类型之前加...表示变长参数, 如果vals 在函数体内继续使用,当做实参传递时候应该是:val... 来表示
```
func main()  {
	fmt.Println(sum())
	fmt.Println(sum(3))
	fmt.Println(sum(1, 2, 3, 4, 5))
	
	values := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(values...)) //当实参已经存在于一个 slice 中的时候在最后一个参数后面放一个省略号...,
}

func sum(vals ...int) int  {
	total := 0
	for _, val := range vals { // 在函数体内 vals 是一个 int类型slice
		total += val
	}
	return total
}
```
* 尽管...int参数就像函数体内的 slice.但变长函数的类型和一个带有普通 slice 参数的函数类型不相同
```
func f(...int) {}
func g([]int) {}

//  vals ...int 在形参里数据类型之前加...表示变长参数, 如果vals 在函数体内继续使用,当做实参传递时候应该是:val... 来表示
func errorf(linenum int, format string, args ...interface{}){
    fmt.Fprintf(os.Stderr, format, args...)
}
```
* interface{}类型意味着这个函数的最后一个参数可以接受任意值.

## 5.8 延迟函数调用
* defer语句就是一个普通的函数或方法调用.在调用之前加上关键字 defer. 函数和参数表达式会在语句执行时求值,但是无论是正常情况下,执行 return语句或函数执行完毕,还是不正常的情况下,比如发生宕机,实际的调用推迟到包含 defer语句的函数结束后才执行.
* defer语句没有限制次数,执行的时候以调用 defer 语句的顺序的倒序进行 相当于栈中的数据
* defer语句经常使用于成对的操作,比如打开和关闭,连接和断开,加锁和解锁,即使是再复杂的控制流,资源在任何情况下都能够正确释放. 正确使用 defer语句的地方是在成功获得资源之后.
```
defer resp.Body.Close() 关闭网络资源
defer f.Close() // 关闭一个打开的文件
defer mu.Unlock() // 解锁一个互斥锁
```
* 延迟执行的函数在 return语句之后执行,并且可以个更新函数的结果变量. 因为匿名函数可以得到其外层函数作用域内的变量(包括命名的结果),所以延迟执行的匿名函数可以观察到函数的返回结果
* 延迟函数不到最后一刻是不会执行的. 要注意循环里defer语句的使用.
```
for _,filename := range filenames {
    f, err := os.Open(filename)
    if err != nil {
        return nil
    }
    defer f.Close() // 可能会用尽文件描述符
}
// 解决方式:将循环体(包括 defer的语句)放到另一个函数中,每次循环迭代都会调用文件关闭函数
for _, filename := range filenames {
    if err := doFile(filename); err != nil {
        return err
    }
}

func doFile(filename string) error {
    f, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer f.Close()
}
```

## 5.9宕机
* go运行时候检测到错误,就会触发宕机
* 一个典型的宕机发生的时候,正常的程序执行会终止,goroutine中的所有延迟函数都会执行.然后程序会异常退出并留下一条日志消息.日志消息包括宕机的值,显示一个函数调用的栈追踪消息.
* 可以直接调用内置的宕机函数.内置的宕机函数可以接受任何值作为参数. panic()
* 宕机会引起程序异常退出,因此只有发生严重错误的时候才会使用宕机.
* regexp.Compile() 高效的正则表达式. 大部分的正则表达式是字面量,regexp 包有包装函数 regexp.MustCompile()
```
packeage regexp
func Compile(expr string) (*Regexp, error) { /** */}
func MustCompile(expr string) *Regexp {
    re, err := Compile(expr)
    if err != nil {
        panic(err)
    }
    return re
}
```
* 包装函数使得初始化一个包级别的正则表达式变量(带有一个编译的正则表达式)变得更加方便.
```
var httSchemeRE = regexp.MustCompile(`^http?:`) // http: https:
```
* 宕机发生时,所有的延迟函数以倒序执行.从栈最上面的函数开始.一直返回至 main 函数
* runtime 包提供了转储栈的方法,可以用来诊断错误. runtime.Stack()

## 5.10 恢复
* recover() 如果内置的recover函数在延迟函数的内部调用,而且这个包含 defer语句的函数发生宕机,recover 会终止当前宕机状态,并且返回宕机的值.函数不会从之前宕机的地方继续运行而是正常返回.如果 recover 在其他任何情况下运行则它没有任何效果且返回 nil
* 将宕机的错误看做一个解析错误,不要立即终止运行,而是将一些有用的附加信息提供给用户来报告这个 bug
* 
```
// Parse函数中的延迟函数会从宕机中恢复,并使用宕机的值,组成一条错误消息. 延迟函数将错误赋值给err结果变量,从而返回给使用者
func Parse(input string) (s *Syntax, err error) {
    defer func() {
        if p := recover(); p != nil {
            err = fmt.Errorf("internal error: %v", p)    
        }
    }
    // 解析器
}
```
* 最安全的做法还是选择性使用recover,宕机过后,需要恢复的情况本来就不多.