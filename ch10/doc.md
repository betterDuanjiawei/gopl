# 10. 包和 go工具
* http://godoc.org

## 10.1 引言
* 每个包定义了一个不同的命名空间作为它的标识符. 每个名字关联一个具体的包
* 包通过控制名字是否导出使其对包外可见来提供封装能力. 限制包成员的可见性
* 修改一个文件时候,我们必须重新编译文件所在的包和所有潜在依赖他的包.
* go编译速度快的原因:
    1. 所有的导入都必须在一个源文件的开头进行显式列出.这样编译器在确定依赖性的时候就不需要读取和处理整个文件
    2. 包的依赖形成有向无环图,因为没有环,所以包可以独立甚至并行编译
    3. go 包编译输出的目标文件不仅记录了它自己的导出信息,还记录它所依赖包的导出信息.当编译一个包的时候,编译器必须从每一个导入中读取一个目标文件,但是不会超出这些文件
## 10.2 导入路径
* 每个包都通过一个唯一的字符串进行标识,它称为导入路径.它们用在 import声明中
* 对于准备共享或公开的包,导入路径需要全局唯一.
* 为了避免冲突,除了标准库中的包之外,其他包的导入路径都必须以互联网域名作为路径开始,这样也方便查找包
## 10.3 包声明
* 在每一个源文件的开头都需要进行包声明.主要目的是当该包被其他包引入的时候作为其默认的标识符(包名)
* 通常,包名是导入路径的最后一段,于是,即使导入路径不同的两个包,二者也可以拥有相同的名字.
```

```
* 最后一段的例外:
    1. 不管包的导入路径是什么,如果该包定义了一条命令(可以执行的 go 程序),那么它总是使用名称 main.这是告诉 go build 的信号,它必须调用连接器生成可执行文件
    2. 目录中可能有一些文件名字是_test.go 结尾.包名中会出现以_test 结尾. 
    这样的目录中有两个包,一个普通的,加上一个外部测试包._test后缀告诉 go test 两个包都需要构建,并且指明文件属于哪个包.外部测试包用来避免测试所依赖的导入图中的依赖循环
    3. 一些依赖管理工具回在包导入路径的尾部追加版本号后缀,如`gopkg.in/yaml.v2` 包名不包含后缀,因此这个情况下包名为 yaml

## 10.4 导入声明
* 一个 go源文件可以在 package声明的后面和第一个非导入声明语句前面紧接着包含零个或多个 import声明.每个导入都可以单独指定一条导入路径.也可以通过圆括号括起来的列表一次导入多个包
* 可以通过空行分组,这类分组通常表示不同领域和方面的包.导入的顺序不重要,但按照惯例每一组都按字母进行排序(gofmt 和 goimports工具都会自动进行分组并排序
* 重命名导入:如果需要把两个名字一样的包(math/rand 和 crypto/rand)导入到第三个包中,导入声明就必须至少为其中一个指定一个替代名字来避免冲突. 
    1. 替代名字仅影响当前文件.其他文件(即使是一个包中的文件)可以使用默认名字来导入包.或者一个替代名字也可以
    2. 重命名导入在没有冲突的时候也是非常有用的.如果有时用到自动生成的代码,导入的包名字非常冗长,使用一个替代名字可能更方便. 同样的缩写名字要一直用下去,以避免产生混淆
    3. 使用一个替代名字有助于规避常见的局部变量冲突.例如,如果一个文件可以包含许多以 path 命名的变量,我们就可以使用 pathpkg这个名字导入一个标准的"path"包
```
import "fmt"
import "os"

// 圆括号的更加常见
import (
    "fmt"
    "os"

    // 可以通过空行分组,这类分组通常表示不同领域和方面的包.导入的顺序不重要,但按照惯例每一组都按字母进行排序(gofmt 和 goimports工具都会自动进行分组并排序)
    "golang.org/x/net/hmtl"
    "golang.org/x/net/ipv4"

    "crypto/rand"
    mrand "math/rand" // 通过指定一个不同的名称 mrand就避免了冲突
)
```
* 每个导入声明从当前包向导入的包建立一个依赖.如果这些依赖形成一个循环,go build 工具会报错

## 10.5 空导入
* 如果导入的名字咩有在文件中引用,就会产生一个编译错误.但是有时候,我们必须导入一个包,这仅仅是为了利用其副作用:对包级别的变量执行初始化表达式求值,并执行它的 init 函数.
* 空白导入: 为了防止"未使用的导入"错误,我们必须使用一个重命名导入,它使用一个替代名字_,这表示导入的内容为空白标识符.通常情况下,空白标识符不可能被引用
* 多数情况下,它用来实现一个编译时的机制,使用空白引用导入额外的包,来开启主程序中的可选特性.
* 标准库提供了 GIF PNG JPEG 等格式的解码库,用户自己可以提供其他格式的,但是为了使可执行程序简短,除非明确需要,否则解码器不会被包含进应用程序
* image.Decode 对于每一种格式,通常通过在其支持的包的初始化函数中来调用 image.RegisterFormat 来向表格添加项.
```
package png

func init() {
	image.RegisterFormat("png", pngHeader, Decode, DecodeConfig)
}

```
* 这个效果就是,一个应用只需要空白导入格式化所需要的包,就可以让 image.Decode 函数具备对应格式的解码能力
* database/sql 包使用类型的机制让用户按需加入想要的数据库驱动程序:
```
import (
	"database/sql"
	_ "github.com/lib/pq" // 添加 Postgres支持
	_ "github.com/go-sql-driver/mysql" // 添加 mysql 支持
)
func main()  {
	db, err = sql.Open("postgres", dbname) // ok
	db, err = sql.Open("mysql", dbname) // ok
	db, err = sql.Open("sqlite3", dbname) // unknown driver sqlite3
}
```

## 10.6 包及其命名
* 当创建一个包时, 使用简短的名字,但是不要短到像加密了一样,比如:bufio os time io
* 尽可能的保持可读性和无歧义.例如,不要把一个辅助工具包命名为 util,使用 imageutil 或 ioutil 等名称更具体和清晰
* 避免选择经常用于相关局部变量的包名,或者迫使使用者使用重命名导入,比如 path命名 的包
* 包名经常使用统一的形式.标准包 bytes errors strings 使用复数来避免覆盖响应的预声明类型,使用 go/types 这个形式,来避免和关键字的冲突
* 避免使用由其他含义的包名. tempconv
* 包成员的命名. 因为对其他包的成员的每一个引用使用一个具体的标识符,fmt.Println.描述包的成员与描述包名与成员名同样复杂. 我们设计一个包时,要考两个有意义的部分如何一起工作,而不只是成员名.
```
bytes.Equal
flag.Int
http.Get
json.Marshal
```
* 单一类型包,例如 html/template math/rand,这些包导出一个数据类型及其方法,通常有一个 New()函数用来创建实例
```
package rand
type Rand struct{}
func New(source Source) *Rand
```
## 10.7 go 工具
* go 工具将不同种类的工具集合并为一个命名集.他是一个包管理器类似于 apt或 rpm.
```
Go is a tool for managing Go source code.
Usage:
        go <command> [arguments]
The commands are:
        bug         start a bug report
        build       compile packages and dependencies
        clean       remove object files and cached files
        doc         show documentation for package or symbol
        env         print Go environment information
        fix         update packages to use new APIs
        fmt         gofmt (reformat) package sources
        generate    generate Go files by processing source
        get         download and install packages and dependencies
        install     compile and install packages and dependencies
        list        list packages or modules
        mod         module maintenance
        run         compile and run Go program
        test        test packages
        tool        run specified go tool
        version     print Go version
        vet         report likely mistakes in packages

```
### 10.7.1 工作空间的组织
* GOPATH 环境变量,它指定工作空间的根.当需要切换到不同的工作空间时,更新 GOPATH变量的值即可.
`export GOPATH=$HOME/go/src` 将 GOPATH 设置为 xxx
* GOPATH 有3个子目录:
    1. src 目录包含源文件.每个包放在一个目录中,该目录相对于$GOPATH/src 的名字是包的导入路径,如:gopl.io/ch1/helloworld,一个 GOPATH工作空间在 src下包含多个源代码版本控制仓库,如 gopl.io golang.org
    2. pkg 子目录是构建工具存储编译后的包的位置.
    3. bin 子目录存放像 helloworld 这样的可执行程序.
* 第二个环境变量GOROOT用来指定 GO的安装目录,还有它自带标准包库的位置.GOROOT的目录结构和 GOPATH类似,因此存放 fmt包的源代码对应目录应该是$GOROOT/src/fmt
用户一般不需要设置 GOROOT,默认情况下 go语言安装工具回将其设置为安装的目录路径
* go env 用于查看 GO语言工具涉及的所有环境变量的值.包括未设置环境变量的默认值.
    1. GOOS 环境用于指定目标操作系统 
    2. GOARCH 环境变量用于指定处理器的类型 amd64 arm
### 10.7.2 下载包
* go get 可以下载一个单一的包或者用...下载整个子目录里面的每一个包. 同时计算并下载所依赖的每个包
* golint工具用于检测 go源代码的编程风格是否有问题,
```
安装
go get github.com/golang/lint/golint 
检查
$GOPATH/bin/golint gopl.io/ch2/popcount
```
* go get命令获取的代码是真实的本地存储仓库,而不仅仅是复制源文件.因此你依然可以使用版本管理工具比较本地代码的变更或切换到其他的版本.
* 需要注意的是导入路径含有的网站域名和本地 git仓库对应远程服务地址并不相同.真实的 git地址是go.googlesource.com.
* 可以让包用一个自定义的导入路径,但是真实的代码却是由更通用的服务提供.
* go get -u 命令行标志参数,将确保所有的包和依赖的包的版本都是最新的,然后重新编译和安装他们,如果不包含该标志参数的化,而且如果包已经在本地存在,那么代码将不会自动更新
* vendor的目录用于存储依赖包的固定版本的源代码. go help gopath
### 10.7.3 构建包
* go build 命令编译命令行参数指定的每个包,如果包是一个库,则忽略输出结果.还可以用于检测包是可以正确编译的,
* 如果包的名字是main,go build 将调用链接器在当前目录创建一个可执行程序;以导入路径的最后一段作为可执行程序的名字
* 由于每个目录只包含一个包,因此每个对应可执行程序或者叫 unix术语中的命令的包,会要求放到一个独立的目录中. cmd目录下面.
```
$ cd $GOPATH/src/gopl.io/ch1/helloworld
$ go build

$ cd anywhere
$ go build gopl.io/ch1/helloworld

$ cd $GOPATH
$ go build ./src/gopl.io/ch1/helloworld

$ cd $GOPATH
$ go build src/gopl.io/ch1/helloworld
Error: cannot find package "src/gopl.io/ch1/helloworld".
```
* 也可以指定包的源文件列表.这一般用于构建一些小程序或做一些临时性的实验.如果是 main包,将会以第一个 go源文件的基础文件名 作为最终可执行程序的名字
* go run xxx 构建并运行,适用于一次性运行的程序
* go install 命令和 go build 命令很相似,但是它会保存每个包的编译成果,而不是将他们都丢弃.被编译的包将会被保存到$GOPATG/pkg 目录下,目录路径和 src目录路径对应,可执行程序被保存到$GOPATH/bin目录
* go install 和 go build 命名都不会重新编译没有发生变化的包,这可以使后续构建更快捷,为了方便编译依赖的包,go build -i 命令将安装每个目标所依赖的包
* 因为编译对应不同的操作系统平台和 cpu架构.go install 命令会将编译结果安装到 GOOS和 GOARCH 对应的目录;mac $GOPATH/pkg/darwin_amd64目录
* 针对不同操作系统或 cpu 的交叉构建也很简单,只需要设置好目标对应的 GOOS和 GOARCH,然后运行构建命令即可
```
GOARCH=386 go build gopl.io/ch10/cross
```
* 有些包可能需要针对不同平台和处理器类型使用不同版本的代码文件,以便于处理底层的可移植性问题或为一些特定的代码提供优化.
* 如果一个文件名包含了一个操作系统或处理器类型名字,例如net_linux.go 或 asm_amd64.s go语言的构建工具将只在对应的平台编译这些文件.
* 还有一个特别的构建注释参数可以提供更多的构建过程控制.
```
// +build linux darwin
```
* 在包声明和包注释的前面, 该构建注释参数告诉 go build只在编译程序对应的目标操作系统是 linux 或 Mac OS X才编译这个文件.
* 下面的构建注释则表示不编译这个文件
```
// +build ignore
```
* 更多  go doc go/build

### 10.7.4 包文档
* 专门用于保存包文档的源文件叫做doc.go
* go doc,该命令打印其后所指定的实体的声明与文档注释.该实体可能是一个包
```
该实体可能是一个包
go doc time
某个具体包成员
go doc time.Since
或是一个方法
go doc time.Duration.Seconds
```
* 该命名并不需要输入完整的包导入路径或正确的大小写.
```
go doc json.decode
```
* godoc,可以提供相互交叉引用的HTML页面,但是包含和 go doc 相同以及更多信息
* godoc的在线服务 https://godoc.org 
* 自己的工作区目录运行 godoc服务,其中-analysis=type和-analysis=pointer命令行标志参数用于打开文档和代码中关于静态分析的结果。
```
godoc -http:8000
运行下面的命令，然后在浏览器查看 http://localhost:8000/pkg 页面：
```
执行上面的命名报错: command not found: godoc
[参考](https://stackoverflow.com/questions/63442354/godoc-command-not-found)
### 10.7.5 内部包
* 在 go中,包是最重要的封装机制.没有导出标识符只有在同一个包内部可以访问,而导出的标识符则是面向全宇宙都是可见的
* 有时候一个中间状态也是有用的,标识符对于一小部分包是信任可见的.但并不是对所有调用者都可见.
* go语言的构建工具对包含internal名字的路径段的包导入路径做了特殊处理.这种包叫做 internal包,一个 internal包只能被和internal目录有同一个父目录的包导入,
例如，net/http/internal/chunked内部包只能被net/http/httputil或net/http包导入，但是不能被net/url包导入。不过net/url包却可以导入net/http/httputil包。
```
net/http
net/http/internal/chunked
net/http/httputil
net/url
```
### 10.7.6 查询包
* go list命令可以查询可用包的信息.可以用来测试包是否在工作区并打印它的导入路径
```
go list github.com/go-sql-driver/mysql
can't load package: package github.com/go-sql-driver/mysql: 
    cannot find package "github.com/go-sql-driver/mysql" in any of:
    /usr/local/go/src/github.com/go-sql-driver/mysql (from $GOROOT)
```
* go list ... 用...匹配任意的包导入路径,可以列出工作区内的所有包
* go list gopl.io/ch10/... 特定子目录下的所有包
* go ...xml... 某一主题相关的包
* go list 还可以获取每个包完整的元信息,而不仅仅是导入路径,这些元信息可以以不同形式提供给用户,其中-json表示用 json 格式 打印每个包的元信息
```
go list -json hash
```
* -f 允许用户使用text/template包的模板语言定义输出文本的格式.
```
用join模板函数将结果链接为一行,连接时每个结果之间用空格分割
go list -f '{{join .Deps " "}}' strconv
所有包的导入包列表
go list -f '{{.ImportPath}} -> {{join .Imports " "}}' strconv/...
```
* go list 命令对于一次性交互式查询或自动化构建或测试脚本都很有帮助.