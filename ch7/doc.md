# 7.接口
* 接口类型是对其他类型行为的抽象和概括.通过接口.我们可以写出更加灵活和通用的函数,这些函数不用绑定在一个特定的类型实现上.
* go 的接口是隐式实现的. 对于一个具体的类型,无须声明它实现了哪些接口,只要提供接口所必须的方法即可.

## 7.1 接口即约定
* 之前介绍的类型都是具体类型,具体类型指定的它所包含数据的精确布局,还暴露了基于这个精确布局的操作,比如数值有算术操作,slice类型有索引 append range 等操作
具体类型还会提供其方法来提供额外的能力.总之,如果你知道了一个具体类型的数据,就知道了它是什么以及它能干什么
* 接口类型是抽象类型,它并没有暴露所含数据的布局或者内部结构,当然也没有那些数据的基本操作,它能提供的仅仅是一些方法而已,接口类型,你需要知道它提供了哪些方法
* 可取代性:可以把一种类型替换为满足同一接口的另一种类型的特性称为可取代性.面向对象的典型特征

## 7.2接口类型
* 一个接口类型定义了一套方法,一个具体的类型要实现该接口,那么必须实现接口类型中定一个的所有方法
* io.Writer是一个广泛使用的接口,它负责所有可写入字节的类型抽象,包括文件 内存缓冲区 网络连接 http客户端 打包器(archiver) 散列器(hasher)等
* io.Reader 抽象了所有可以读取字节的类型, io.Closer 抽象了可以关闭的类型,文件或网络连接
* go 单方法接口命名的约定
```
package io

type Writer interface {
    Write()
}
type Reader interface {
    Read(p []byte) (n int, err error)
}
type Closer interface {
    Close() error
}
```
* 可以通过组合已有的接口得到新的接口, 下面的语法称为嵌入式接口,与嵌入式结构类似,让我们可以直接使用一个接口,而不用逐一写出这个接口所报含的方法.
```
type ReadWriter interface {
    Reader
    Writer
}
type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}

```
* 组合接口也可以不用嵌入式接口来声明,也可以用两种方式组合声明
```
type ReadWriter interface {
    Read(p []byte) (n int, err error)
    Write(p []byte) (n int, err error) 
}
type ReadWriter interface {
    Read(p []byte) (n int, err error)
    Writer // 这里是 Writer 而不是 Write
}
```
* 上面三种声明的效果是一样的.方法定义的顺序也是无意义的.真正有意义的只有接口的方法集合

## 7.3实现接口
* 如果一个类型实现了接口要求的所有方法,那么这个类型实现了该接口.
* 一个具体类型"是一个"(is-a)特定的接口类型,这其实代表着该具体类型实现了该接口. *byte.Buffer 是一个io.Writer. 代表*byte.Buffer 实现了 io.Writer 接口
* 赋值只有在值对于变量类型是可赋值的时候才合法
* 仅当一个表达式实现了一个接口时候,这个表达式才可以赋值给该接口
* 当右侧表达式也是一个接口时候,该规则也有效
```
func demo1() {
	var w io.Writer
	w = os.Stdout // *os.File
	w = new(bytes.Buffer) // *bytes.Buffer
	//w = time.Second
	fmt.Printf("%T\n", w)

	var rwc io.ReadWriteCloser
	rwc = os.Stdout // *os.File 有 Read() Write() Close()方法
	//rwc = new(bytes.Buffer) // *bytes.Buffer缺少 Close 方法
	fmt.Printf("%T\n", rwc)

	w = rwc // ok, io.ReadWriteCloser有 write方法
	rwc = w // io.Write 缺少 Close 方法

}
```
* 如果 B接口包含了 A接口的所有方法,那么任何实现了 B接口的类型都必然实现了 A接口
* interface{} 空接口类型,空接口类型对于其实现类型没有任何要求,所以我们可以将任何值赋值给空接口类型.
```
func demo2() {
	var any interface{}
	any = true
	any = 12.34
	any = "hello"
	any = map[string]int{"one":1}
	any = new(bytes.Buffer)
}
```
* 需要类型断言来从一个空接口中还原实际值

* 判断是否实现一个接口,只需要比较具体类型和接口类型的方法,所以没必要在具体类型的定义中声明这种关系,所以偶尔在注释中标注也不坏,但是对于程序来说这种声明不是必须的.
```
    var w io.Writer = new(bytes.Buffer)
	var _ io.Writer = (*bytes.Buffer)(nil) // 将 nil 转换为指针类型的 bytes.Buffer
```
* 非空的接口类型,通常由一个指针类型来实现.特别是当接口类型的一个或多个方法暗示会修改接收者的情形. 一个指向结构的指针才是最常见的方法接收者
* 指针类型肯定不是实现接口方法的唯一类型,也可以是其他类型,slice类型的方法 map类型的方法 函数类型的方法 基础类型也可以实现方法
* 一个具体的类型可以实现很多不相关的接口.
* 从具体类型出发,提取其共性而得出的每一个分组的方式都可以表示为这一种接口类型.

## 7.4使用 flag.Value 来解析参数
* 使用 flag.Value来帮助我们定义命令行标志
```
package flag
type Value interface {
    String() string
    Set(string) error
}
```

## 7.5接口值
* 从概念上来讲,一个接口类型的值(接口值)其实有2部分,一个具体类型和该类型的一个值,即动态类型和动态值
```
var w io.Writer // nil
w = os.Stdout
w = new(bytes.Buffer)
w = nil //nil
类型: nil
值: nil
```
* 变量总是初始化为一个特定的值,接口也不例外,接口就是把它的动态类型和值都设置为 nil.
* w == nil 或 w != nil来检测一个接口值是否是 nil,调用 nil接口的任何一个方法都会导致崩溃
```
w.Write([]byte("hello")) // 崩溃,对空指针取引用值
// = os.Stdout.Write([]byte("hello"))
w = os.Stdout
类型: *os.File
值: 
```
* 接口值可以用==和!=操作符来比较.如果两个接口值都是 nil或二者的动态类型完全一致而且动态值相等(==比较),那么两个接口值相等.
* 因为接口值是可以比较的,所以可以作为 map的键或 switch 的操作数
* 比较两个接口值时候,如果两个接口值的动态类型一致,但是对应的动态值是不可比较的,如 slice.那么这种比较会以崩溃的方式失败:
```
var x interface{} = []int{1, 2, 3}
fmt.Println(x == x) //  comparing uncomparable type []int
```
* 用 fmt包里 f %T 来拿到接口的动态类型
```
    var w io.Writer
	fmt.Printf("%T\n", w) // <nil>

	w = os.Stdout
	fmt.Printf("%T\n", w) //*os.File

	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w) //*bytes.Buffer
```
### 含有空指针的非空接口 没看懂,有需要回来继续看吧
* 空的接口值(其中不包含任何信息)与仅仅动态值为 nil的接口值是不一样的.

## 7.6 使用 sort.Interface来排序
* go中 sort.Sort函数对序列和其中元素的布局无任何要求,它使用sort.Interface接口来指定通用排序算法和每一个具体的序列类型之间的协议,
这个接口的实现确定了序列的具体布局(经常是一个 slice)以及元素期望的排序方式
* 一个原地排序算法需要知道3个信息:序列长度 比较2个元素的含义和如何交换两个元素.
```
package sort
type Interface interface {
    Len() int
    Less(i, j int) bool 
    Swap(i, j int)
}

type StringSlice []string
func (p StringSlice) Len() int { return len(p)}
func (p StringSlice) Less(i, j int) {return p[i] < p[j]}
func (p StringSlice) Swap(i, j int) {p[i], p[j] = p[j], p[i]}

sort.Sort(StringSlice(names)) //将 names 普通[]string转换为 StringSlice,增加了三个额外用于排序的方法
sort.Strings(names) // 内置排函数
```
* sort.Interface的具体类型并不一定都是 slice.有可能是其他类型:结构体
[参考](https://www.cnblogs.com/haima/p/12202032.html)
```
    sort.Sort(byArtist(track))
	printTrack(track)
	sort.Sort(sort.Reverse(byArtist(track))) // 这个这样用,
	printTrack(track)
```
* sort的其他方法:IntsAreSorted Ints IntSlice
```
    values := []int{3, 1, 4, 1}
	fmt.Println(sort.IntsAreSorted(values))
	sort.Ints(values)
	fmt.Println(sort.IntsAreSorted(values))

	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println(values)
	fmt.Println(sort.IntsAreSorted(values))
	//false
	//true
	//[4 3 1 1]
	//false
```
* sort 包提供了[]int, []string []float64自然排序的函数和相关类型. []uint []int64其他等需要自己实现

## 7.7 http.Handler接口
* http.Handler 接口
```
net/http
package http

type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}
func ListenAndServe(addres string, h Handler) error // 需要一个服务器地址:localhost:8080和一个 handler实例
```
```
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	// 基于 URl的路径部分:req.URL.Path
	switch req.URL.Path {
	case "/list" :
		for item, price := range  db {
			fmt.Fprintf(w, "%s:%s\n", item, price)
		}
	case "/price" :
		item := req.URL.Query().Get("item") // req.URL.Query()解析为一个 map(multimap), url.Values
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound) //返回一个 HTTP错误 404
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}
```
* net/http 包 请求多工转发器 ServeMux,用来简化 url和处理程序之间的关系
* HandlerFunc
```
package http
type HandlerFunc func(w WriterResponse, r Request)
func (f HandlerFunc) ServeHTTP(w WriterResponse, r Request) {
    f(w, r)
}
```
* ServeMux引入了一个便捷的HandleFunc方法来简化调用
```
func main()  {
	db := database{"shote" : 5, "sock" : 50}
	mux := http.NewServeMux()
	//mux.Handle("/list", http.HandlerFunc(db.list)) // 将 db.list 函数转化为 http.HandleFunc类型
	//mux.Handle("/price", http.HandlerFunc(db.price))
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type dollars float32

func (d dollars) String() string  {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request)  {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request)  {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %s\n", item)
		return //要记得加
	}
	fmt.Fprintf(w, "%s : %s\n", item, price)
}
```
* net/http 包提供了一个全局的 ServeMux实例DefaultServeMux以及包级别的注册函数http.Handle 和 http.HanleFunc.
要让DefaultServeMux作为服务器的主处理程序,无须将它传递给ListenAndServe,直接传 nil 即可
```
    db := database{"shote" : 5, "sock" : 50}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
```

## 7.8 error 接口
* error 类型是一个接口.包含一个返回错误消息的方法
```
tyep error interface{
    Error() string
}
```
* 构造error最简单的方法就是调用errors.New().它会返回一个包含指定错误消息的新 error实例
```
package errors
type errorString struct{ text string} // 没有直接用字符串,为了避免将来无意间的布局变更

func New(text string) error {
    return &errorString{text}
}
// 满足error 接口的是*errorString 指针,而不是原始的errorString,主要是为了让每次 New 分配的error实例都互不相等.
// fmt.Println(errors.New("EOF") == errors.New("EOF")) // false
func (e *errorString) Error() string {
    return e.text
}
```
* 直接调用 errors.New 比较罕见.一般用fmt.Errorf()额外提供了字符串格式化的功能
```
package fmt

func Errorf(format string, args ...interface{}) error {
    return errors.New(Sprintf(format, ...args))
}
```
* syscall.Errno类型 操作系统错误码
```
package syscall
tyep Errno uintptr 

var errors = [...]string{
    1 : "operation not permitted", // EPERM
    2 : "no such file or directory", // ENOENT
    3 : "no such process", // ESRCH
    ...
}
func (e Errno) Error() string {
    if 0 <= int(e) && int(e) < len(errors) {
        return errors[e]
    }

    return fmt.Sprintf("errno %d", e)
}

	var err error = syscall.Errno(15)
	fmt.Println(err.Error())
	fmt.Println(err)
```

## 7.9 示例:表达式求值器 看了一遍,没啥收获,后续可以继续复习

## 7.10 类型断言
* 类型断言是一个作用在接口值上的操作,写出来类型 x.(T),其中 x是一个接口类型的表达式,T是一个类型(断言类型)
* 类型断言会检查作为操作数的动态类型是否满足指定的断言类型
* 断言返回2个值,多一个布尔型的返回值来指示是否断言成功
```
    var w io.Writer = os.Stdout
	if f, ok := w.(*os.File); ok {
		fmt.Printf("%T\n", f) // *os.File
	}
	if b, ok := w.(*bytes.Buffer); !ok {
		fmt.Printf("%v%[1]T\n", b) // <nil>*bytes.Buffer
	}
```

## 7.11使用断言类型来识别错误
```
    os.IsExist() // 文件创建
	os.IsNotExist() // 文件读取
	os.IsPermission() // 权限不足
```
```
_, err := os.Open("no such file")
	fmt.Println(err)
	fmt.Printf("%#v\n", err)
	//open no such file: no such file or directory
	//&os.PathError{Op:"open", Path:"no such file", Err:0x2}
```
## 7.12 通过接口类型断言来查询特性

## 7.13 类型分支
* 接口有两种不同的风格,1. io.Writer 这种风格强调了方法,而不是具体类型
2.可识别联合接口 利用了接口值能够容纳各种类型的能力.类型断言用来在运行时区分这些类型并分别处理.强调的是这个接口的具体类型,而且这种接口经常没有方法.
* type switch
```
switch x := x.(type) { // 新变量 x与外部块中的变量 x不冲突
case nil:
case int, uint:
case string:
case bool:
default:
// 不允许使用 fallthrough``
}
```

## 7.14 基于标记 XML 解析
```
func main()  {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect : %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if containAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

// 注意这个函数和巧妙
func containAll(x, y []string) bool {
	for len(y) <= len(x) { //for 循环
		if len(y) == 0 {
			return true
		}
		// 有啥用呢?
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
```



