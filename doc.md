# go 日记
关于gopl书籍的练习题可以[参考](https://www.cnblogs.com/ling-diary/tag/Go%E7%BB%83%E4%B9%A0%E9%A2%98/)

## struct{} 和 struct{}{} [参考](http://www.36nu.com/post/329)
* struct 是 go的关键字 结构体类型, 
* struct{} 是一个无元素的结构体类型,优点是大小为0,不需要内存来存储struct{}类型的值
* struct{}{} 是一个复合字面量,它构造了一个 struct{}的值,该值也是空的

## os.Stat() 用法和功能
* os.Stat() 获取文件属性 [参考](https://blog.csdn.net/weixin_43851310/article/details/87988648)

```
// 返回描述文件f 的 FileInfo 类型值,如果出错,错误底层类型是*PathError
func (f *File )Stat() (fi FileInfo, err error) == os.Stat(filename string)
type FileInfo interface {
    Name() string // 文件的名字,不含扩展名
    Size() int64 // 普通文件返回值标示其大小,其他文件的返回值含义各个系统不同
    Mode() FileMode // 文件的模式位
    ModTime() time.Time // 文件的修改时间
    IsDir() bool // 等价于 Mode.IsDir()
    Sys() interface{} // 底层数据来源(可以返回 nil)
}
```
## slice array 使用 append时候注意事项
append()是必须赋值给一个 slice 或 array 的

```
var names []string
append(names, name) // 错误写法,会报错:append(names, name) evaluated but not used
names = append(names, name) // 正确写法
```

## bufio.Scanner
bufio.Scanner 结构体 用于读取数据 [由bufio.Reader引出的问题](https://studygolang.com/articles/11436?fr=sidebar)
我们一般在读取数据到缓冲区而且想要使用分隔符分割数据流时候,我们一般使用bufio.Scanner数据结构,而不是 bufio.Reader.
```
func main() {
    scanner := bufio.NewScanner(strings.NewReader("ABCDEFG\nHKML"))
    scanner.Split(ScanWords) // 指定分割方式;也可以自定义实现 SplitFunc方法
    for scanner.Scan() {
        fmt.Println(scanner.Text()) // scanner.Bytes()
    }
}
// 对于scanner.Scan() 读取下一行,并将结尾的换行符去掉,Scan() 在读取到内容时返回true, 没有更多内容时候返回 false; 用法相当于其他语言的迭代器iterator,并且把迭代器指向的数据放到新的缓冲区中,新的缓冲区可以通过scanner.Text()或 scanner.Bytes()获取到
scanner.Scan()默认是以\n(换行符)作为分隔符.如果想要指定分割符,Go 提供了4种ScanBytes(返回单个字节作为 token),ScanLines(返回一行文本)(默认),ScanRunes(返回单个 utf8编码的 rune作为一个token);ScanWords(返回通过空格分割的单词)

```

##  os.exit() runtime.Goexit() return 有什么区别
os.Exit(x) 立即进行给定状态(非零状态)的退出
return:结束当前函数,并返回指定值
runtime.Goexit:结束当前 goroutine,其他 goroutine不受影响,主程序也不受影响
os.Exit():结束当前程序,不管三七二十一

## 内置函数 print() println()
print()在 go中是属于输出到标准错误流中并打印,官方不建议写程序时候使用它,可以在 debug时候用
println() 每次输出都会自动加换行

## fmt包Printf Sprintf Fprintf的区别
xxxf 格式化
xxxln 换行 ln 其实是 使用%v格式化之后,在最后追加换行符

fmt.Print 是属于标准输出流,一般使用它进行屏幕输出 不接受任何格式化操作
fmt.Printf 格式化输出,只可以输出字符串类型的变量
fmt.Println 输出后换行

fmt.Sprint 采用默认格式将参数格式化,串联所有输出生成并返回一个字符串,如果两个相邻的参数都不是字符串,会在它们的输出之间添加空格
```
s1 := fmt.Sprint("abc", "def", 100, true,  "ghy")
println(s1)
abcdef100 trueghy
```
fmt.Sprintf 返回一个格式化的字符串,它本身只返回数据,不打印屏幕 可以赋值给其他变量
func Sprintf(format string, a ...interface{}) string

fmt.Fprint Fprint采用默认格式将其参数格式化并写入w。如果两个相邻的参数都不是字符串，会在它们的输出之间添加空格。返回写入的字节数和遇到的任何错误。
fmt.Fprintf Fprintf根据 format参数生成格式化的字符串并写入 w.返回写入的字节数和遇到的任何错误,是输出到 io.Writer 而不是 os.Stdout
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)

## fmt.Printf 详解
* fmt.Printf 从一个表达式列表生成格式化的输出,第一个参数是格式化指示字符串, 每一个参数的格式是一个转义字符(verb) = % + 一个字符(特殊含义) 
* xxxf 格式化函数都以 f 结尾
```
verb                   描述
%d          十进制整数
%x,%o,%b    十六进制,八进制,二进制
%f,%g,%e(分别代表无指数形式 自动精度形式 有指数形式)(%8.3f 表示8位宽度,小数点后精确到3位的浮点数)    浮点数,如 3.141593, 3,1415926897755, 3.141593e+00
%t          布尔型 true/false
%c          字符 (Unicode码点) 文字符号
%s          字符串
%q          带引号的字符串:"abc" 'a'   fmt.Printf("%% %q %q %q", `a`, 'a', "a")  // % "a" 'a' "a"
%v          内置格式的任何值
%T          任何值的类型 
%%          百分号本身

\t 制表符 \n 换行符
%[1] 重复第一个参数
%#[1]x #输出相应的前缀 fmt.Printf("%v\n", w) // {{{8 8} 5} 20}
                	fmt.Printf("%#v\n", w2) // main.Wheel{Circle:main.Circle{Point:main.Point{X:8, Y:8}, Radius:5}, Spokes:20}
```

## 嵌套 map 出现panic: assignment to entry in nil map 报错
* 嵌套map m := make(map[string]map[string]int) // make只初始化的 map[string]T 部分, T 为 map[string]int, T 没有初始化,需要初始化 T
* [参考](https://blog.csdn.net/jason_cuijiahui/article/details/79410471)
* 常用做法

```
m := make(map[string]map[string]int)
if m[s] == nil {
    m[s] = make(map[string]int)
}
// 赋值
m[s][s1] = 1
```

## ioutil.ReadFile 和 ioutil.ReadAll 的区别
* ioutil.ReadFile(filename string) ([]byte, err)
* ioutil.ReadAll(r io.Reader) ([]byte, err)
* 如果是读取一个文件,那么用 ReadFile,因为它比 ReadAll 快,是因为它先计算了文件的大小,然后初始化对应的 size 大小的 buff,传入 readAll(f, n)来读取字节流
* 如果是 io.Reader 那么必须用 ReadAll 读取全部内容了,或者用 ioutil.Scanner()来逐行读取

## flag 包
* flag.Bool() 创建一个新的 bool标识变量,三个参数:标识的名字n,变量的默认值 false,以及当用户提供非法标识 非法参数或者-h -help参数是输出的消息
* flag.String() 使用名字 默认值 消息来创建一个新的字符串变量
* flag.Bool() 和 flag.String() 返回的都是指向标识变量的指针,必须通过*n和*sep来访问
* flag.Parse() 当程序运行时,在使用标识前,必须调用flag.Parse()来更新标识变量的默认值.如果 flag.Parse()遇到错误,它输出一条帮助信息(和-h -help 输出的信息一样),然后调用os.Exit(2)来结束程序
* flag.Args() 非标识参数可以从flag.Args()返回的字符串 slice 来访问.

## x(T) 类型转换
* 如果两个类型具有相同的底层类型或者二者都是指向相同底层类型变量的未命名指针类型.则二者是可以相互转换的.
* 数字类型之间的转换,字符串和一些slice类型间的转换是允许的.
```
//var test string = "0" // cannot convert test (type string) to type int
	var test float64  =  0.0
	fmt.Println(int(test))
```

## % 取模
* x%y 如果 x<y,那么结果就是 x,如果x>y,那么结果就是取余数 `fmt.Println(2%5) // 2`
* 浮点数和整数 取模结果不一样 只要存在浮点数就不会出现精度丢失 `fmt.Println(5.0/4.0, 5/4, 5.0/4, 5/4.0) // 1.25 1 1.25 1.25 `

## xx.ReadRune()
```
in := bufio.NewReader(os.Stdin)
r, n, err := in.ReadRune() // 返回 rune(解码的字符串 unicode.ReplacementChar 不合法的 utf-8字符,长度是1), nbytes(字节长度), error(错误  io.EOF 文件结束, 其他错误)
```

## fmt.Errorf()输出错信息 是 error类型
* fmt.Errorf() 格式化处理过的附加上下文信息. `return nil, fmt.Errorf("xxx %v", err)`

## 资源显式释放
* go的垃圾回收机制将回收未使用的内存,但不能指望它会释放未使用的操作系统资源.比如打开的文件以及网络连接.必须显式关闭他们.
* resp.Body.Close() 保证正确关闭使得网络资源得以释放,即使在发生错误的情况下也必须释放资源

## 多个复杂变量的同时命名
```
var (
	mu sync.Mutex
	mapping = make(map[string]string)
)
```