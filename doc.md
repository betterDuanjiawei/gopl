# go 日记
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
// 对于scanner.Scan()的用法相当于其他语言的迭代器iterator,并且把迭代器指向的数据放到新的缓冲区中,新的缓冲区可以通过scanner.Text()或 scanner.Bytes()获取到
scanner.Scan()默认是以\n(换行符)作为分隔符.如果想要指定分割符,Go 提供了4种ScanBytes(返回单个字节作为 token),ScanLines(返回一行文本)(默认),ScanRunes(返回单个 utf8编码的 rune作为一个token);ScanWords(返回通过空格分割的单词)

```

##  os.exit() runtime.Goexit() return 有什么区别
os.Exit(x) 立即进行给定状态(非零状态)的退出
return:结束当前函数,并返回指定值
runtime.Goexit:结束当前 goroutine,其他 goroutine不受影响,主程序也不受影响
os.Exit():结束当前程序,不管三七二十一