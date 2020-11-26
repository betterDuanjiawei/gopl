# 基本数据
* go 数据类型分为4大类:基础数据类型,聚合数据类型,引用类型,接口类型
    1. 基础类型: 数字(number) 字符串(string) 布尔(boolean) 
    2. 聚合类型: 数组(array) 结构体(struct)
    3. 引用类型: 指针(pointer) 切片(slice) 字典(map) 通道(channel) 函数(function)
    4. 接口类型: 接口(interface)
## 整数
* go同时具备有符号整数和无符号整数
* 有符号 int8(-127~128) int16 int32 int64
* 无符号 uint8(0~255) uint16 uint32 uint64
* int uint int是目前使用最广泛的数值类型,这两种类型大小相等,都是32位或64位,
* rune类型是 int32类型的同义词,常用于指明一个值是Unicode码点,这两个名称可互换使用
* byte类型是uint8类型的同义词,强调的是一个原始数据,而非量值.
* uintptr 无符号整数,其大小并不明确,但足以完整存放指针
* int uint uintptr 都有别于其大小明确的相似类型的类型. int32和 int是两种不同的类型
* n位数字有符号取值范围是-2(n-1)~2(n-1)-1. 无符号整数由全部位构成其非负值,范围是0~2(n)-1
* go 的二元操作符涵盖了算术 逻辑 比较等运算
```
优先级降序排序
// 前两行的运算符(+)都有对应的赋值运算符+=,用于简写赋值语句
* / % << >> & &^
+ - | ^
== != < <= > >=
&& 
||
```

* 算术运算符 + - * /可应用于整数 浮点数 复数,而取模运算符%仅能用于整数 取模余数的正负号总是与被除数一致,于是-5%3和 -5%-3都得-2.
* 除法运算/行为取决于是否都为整数,整数相除,商会舍弃小数部分,于是5.0/4.0=1.25 5/4=1
* 不论是有符号还是无符号数,若表示算术运算结果所需的位超出该类型的范围,就称为溢出,溢出的高位部分会无提示的丢弃.

* 二元运算符用于比较两个类型相同的整数,比较表达式本身的类型是布尔值
`< <= == != > >=`
全部基本类型的值(布尔值 数值 字符串)都可以比较,这意味着两个相同的类型的值可用== 和 != 运算符比较.整数 浮点数和字符串还能根据比较运算符排序
* 一元加法和一元减法运算符, + - +x(0+x) -x(0-x)
* 类型不匹配最直接的方法是:将全部操作数转换成同一类型 若允许转换,操作 T(x)会将 x的值转换成类型 T.
* 浮点数转成整型,会舍弃小数部分,趋零截尾(正值向下取整,负值向上取整)

* 八进制以0开头,0666,用于表示 POSIX文件系统的权限
* 十六进制 以0x或0X开头,0xdeadbeef, 用于强调其位模式,而非数值大小
* fmt 包 %d(十进制) %b(二进制) %o(八进制) %x(十六进制 0x 小写形式) %X(十六进制 0X 大写形式) 
    1. fmt.Printf()的格式化字符串含有多个%谓词,这要求提供相同数目的操作数,而%后的副词[1]告知 Printf 重复使用第一个操作数.
    2. %#o %#[1]x %#[1]X 之前的副词#告知 Printf 输出相应的前缀0, 0x, 0X
* 文字符号(rune literal)形式是字符写在一对单引号内 '', ASCII字符'a', Unicode码点 码值转义
* %c 输出文字符号 %q输出带有单引号的形式
```
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)
	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)

	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii) //97 a 'a'
	fmt.Printf("%d %[1]c %[1]q\n", unicode) //22269 国 '国'
	fmt.Printf("%d %[1]c %[1]q\n", newline)
	//10
	//'\n'
```
## 浮点数
* go具有 float32和 float64两种大小的浮点数
* math 包给出了浮点值的极限.常量 math.MaxFloat32和 math.MaxFloat64 最大值 math.MinFloat32 math.MinFloat64
* 十进制下float32的有效数字大约是6位,float64大约是15位,绝大多数情况下,应优先选用float64,因为除非格外小心,否则float32位的运算会迅速积累误差.
* 非常小或者非常大的数字最好用科学记数法表示,在数量级指数前写字母 e或 E
* 浮点数能方便的通过Printf的谓词%g输出,该谓词会自动保持足够的精度,并选择最简洁的表示方式,对于数据表%e(有指数) %f(无指数)形式可能更加合适.这三个谓词都能掌握输出宽度和数值精度.
```
    // float
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d, ex = %8.3f, %[2]g, %[2]e\n", x, math.Exp(float64(x)))
	}

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // 0 -0 +Inf(正无穷大) -Inf(负无穷大) NaN(not a number, 如0/0, Sqrt(-1))
	fmt.Println(math.IsNaN(z), math.IsNaN(z/z), math.NaN(), math.NaN() == math.NaN(), math.NaN() == z/z, math.NaN() != z/z) // false true NaN false false true
	//var test string = "0" // cannot convert test (type string) to type int
	var test float64  =  0.0
	fmt.Println(int(test))
```
* math.IsNaN()函数用于判断其参数是否是非数值,math.NaN()函数则返回非数值(NaN) 与 NaN的比较总是不成立的,除了!=

## 复数
* go具有complex64和 complex128两种大小的复数,两者分别由float32和 float64构成
* 内置的complex()函数根据给定的实部和虚部创建复数,而内置的 real() imag()函数分别提取复数的实部和虚部
* 源码中,如果在浮点数或十进制整数后面紧接着写字母 i,如:3.1415i 或2i,它就变成了一个虚数,表示一个实部为0的复数.
* 复数常量可以和其他常量相加(整数 浮点数 实数 虚数皆可) `x := 1+2i`
* 可以用== 或!=判断复数是否等值,若两个复数的实部和虚部都相等,则它们相等
* math.cmplx包提供了复数运算所需的库函数,`cmplx.Sqrt(-1)`

## 布尔值
* boolean值只有两种可能,true 或 false. if 和 for语句的条件就是布尔值,比较操作符 == > 也能得出布尔值结果.
* ! 一元逻辑符!表示取反,(!true == false)  == true.
* true == x,简写为 x
* 布尔值可以由 && 或 ||组合运算,这可能会引起短路行为,如果运算符左边的操作数已经直接确定结果,那右边的操作数不会计算在内.  
* && 较 || 优先级更高, 所以 x && y || a && b  || c && d 无须加()
* 布尔值无法隐式转化为数值(0或1),反之也不行.如果要转化必须要用显式转换
```
转换函数
func btoi(b bool) int {
    if b {
        return 1
    }
    return 0
}

func itob(i int) bool {
    return i != 0
}
``` 

## 字符串
* 字符串是不可变字节序列.它可以包含任意数据,包括0值字节,但主要是人类可读的文本.习惯上.文本字符串被解读成按 utf-8编码的Unicode码点(文字符号)序列.
* 内置 len()函数返回字符串的字节个数(并非文字符号的数目),下标访问操作 s[i]取得第 n 个字符,其中0<=i<len(s)
* 子串生成操作s[i:j]产生一个新字符串.内容取自原字符串的字节,下标从i(含边界),到 j(不含边界值),结果大小是j-i个字节
* 若下标越界或者j值小于i,将触发宕机异常
* 操作数 i 和 j默认值分别是0 字符串起始位置和len(s) 字符串终止位置,若省略i或 j,或两者,则取默认值 s[:]
* + 连接两个字符串生成一个新字符串
* 字符串可以通过比较运算符做比较. == <
* 尽管可以将新值赋值给字符串变量,但是字符串值无法改变.字符串值本身所包含的字节序列永不可变.
```
    // string
	s := "hello 世界"
	fmt.Println(len(s)) // 12
	fmt.Println(s[:5]) // hello
	fmt.Println(s[7:]) // ��界
	fmt.Println(s[0]) // 104
	fmt.Println(s[6:9]) // 世
	//fmt.Println(s[0:len(s)+1]) // 宕机:下标越界 runtime error: slice bounds out of range [:13] with length 12
	t := s
	s = "你好 world"
	fmt.Println(t, s) // hello 世界 你好 world
```
* 因为字符串不可变,所以字符串内部的数据不允许修改 s[0] = 'L' // 编译错误 cannot assign to s[0]
* 不可改变意味这两个字符串能安全的共用一段底层内存,使得复制任何长度字符串的开销都很低廉.
* 一个字符串和两个字符串子串.它们共用底层字节数组,不会分配新的内存

### 字符串字面量
* 字符串的值可以直接写成字符串字面量.形式上就是带双引号的字节序列 "hello 时间"
* 在源码中我们可以将任意值的字节插入字符串中.
```
转义字符.表示 ASCII控制码
\a 警告或响铃
\b 退格符
\f 换页符
\n 换行符(直接跳到下一行的同一个位置)
\r 回车符(返回行首)
\t 制表符
\v 垂直制表符
\' 单引号(仅用于文字字符字面量''\')
\" 双引号(仅用于"..."字符串字面量内部
\\ 反斜杠
```
* 原生的字符串字面量是``,使用反引号而不是双引号.原生字符串字面量中,转义序列不起作用.实质内容与字面量写法严格一致.包括反斜杠和换行符.因此在程序中.原生的字符串字面量可以展开多行.唯一的特殊处理是回车符会被删除.换行符会保留.
* 正则表达式往往包含大量\,可以方便的写成原生的字符串字面量.原生的字符串字面量也适用于 HTML 模板 json字面量 命令行提示信息,以及需要多行文本表达的场景

### Unicode
* Unicode码点的标准数字,这些字符记号称为文字符号(rune),数据类型实 int32. rune 类型是 int32的别名
* UTF-8 每个文字符号用1~4个字节表示.ASCII字符编码仅占一个字节,其他常用的文字符号的编码是2或3个字节.
* unicode 包 unicode/utf-8包
* strings包  Contains()
```
// HasPrefix tests whether the string s begins with prefix
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
}

// HasSuffix tests whether the string s ends with suffix.
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}
```
* unicode/utf-8包 utf8.RuneCountInString(str) 统计文字符号个数 r, size := utf8.DecodeRuneInString(s[i:]) // r 文字符号本身 size 按 utf-8编码所占用字节数.
* range循环也适用于字符串,按 utf-8隐式解码
* 每次UTF-8解码器读取一个不合理的字节,无论是显示调用 utf8.DecodeRuneInString() 还是 range循环内隐式读取,都会专门产生一个'\uFFFD'替换它.其输出通常是黑色六角形或类似砖石的形状.里面有个白色的问号.

* []rune r := []rune(str) string(r)
```
	// []rune
	str2 := "段佳伟"
	fmt.Printf("% x\n", str2)
	fmt.Println(str2)
	r2 := []rune(str2)
	fmt.Printf("%x\n", r2)
	fmt.Println(r2)
	fmt.Println(string(r2))
	fmt.Println(string(65))
	fmt.Println(string(0x4eac))
	fmt.Println(string(1234567))
//	e6 ae b5 e4 bd b3 e4 bc 9f
//	段佳伟
//	[6bb5 4f73 4f1f]
//	[27573 20339 20255]
//	段佳伟
//	A
//	京
//	�
```

### 字符串和字节 slice
* 对字符串操作的4个包: strings bytes strconv unicode
* strings 用于搜索 替换 比较 修整 切分 与连接字符串
* bytes 用于操作字节 slice []byte(其某些属性和字符串相同),由于字符串不可变.因此按照增量方式构建字符串会导致多次内存分配和赋值,因此使用 byte.Buffer类型会更高效
* strconv 主要用于整形 浮点型 布尔值 和字符串之间的互相转换,还有为字符串添加 去除引号的函数
* unicode 有判别文字符号值特性的函数,IsDigit IsLetter IsUpper IsLower
```
slash := strings.LastIndex(s, "/") // 如果没找 ,slash取值-1
```
* path包处理以'/'分段的路径字符串.不分平台 URL地址的路径部分
* path/filepath 包根据宿主平台的规则处理文件名

* 字符串和字节 slice 可以相互转换 string(b) []byte(s)
* []byte(s) 转换操作会分配新的字节数组,拷贝填入s含有的字节,并生成一个 slice 引用,指向整个数组.复制有必要确保 s的字节维持不变.
* string(b) 将字节 slice 转换为字符串也会产生一个副本.保证 s不可变
* 为避免转换和不必要的内存分配,bytes包和 strings 包都有很多使用的函数
```
func Contains(s, substr string) bool 
func Count(s, substr string) int 
func Fields(s string) []string 
func HasPrefix(s, prefix string) bool 
func HasSuffix(s, suffix string) bool 
func Index(s, substr string) int
func Join(elems []string, sep string) string 

bytes 包的操作对象由字符串变成了字节 slice
```
* bytes 包为高效处理字节 slice提供了 Buffer类型.Buffer起初为空,其大小随着各种类型的数据写入而增长.如 string byte []byte
* bytes.Buffer变量无须初始化,原因是其零值本来就是有效的
```
var buf bytes.Buffer
buf.WriteByte() // 追加 ASCII码值
buf.WrtieRune() // 追加任意文字符号的UTF-8编码
buf.WriteString() // 追加字符串
buf.String() // 将  bytes.Buffer类型 转换为字符串
```

### 字符串和数字的相互转换
* strconv 包
* 将整数转换为字符串,1:fmt.Sprintf("%d", x) 2:strconv.Itoa() // integer to ASCII
* 字符串转数值,strconv.FormatInt(int64(x), 2) strconv.FormatUint() 可以按照不同的进制位格式化数字 2: fmt.Printf() %b %d %o %x
* x, err := strconv.Atoi("123") 解释表示整数的字符串
* x, err := strconv.ParseInt("123", 10, 64) 

## 常量
* 常量是一种表达式,其可以保证在编译阶段就计算出其表达式的值.并不需要等到运行时,从而使编译器得以知晓其值.所有常量本质上都属于基本类型:布尔值 字符串 数字
* 常量的声明定义了具名的值,该值恒定,防止程序运行过程中的意外或恶意修改.
* 常量声明可以同时指定类型和值,如果没有显式指定类型,则类型根据右边的表达式推断.time.Duration 是一种具名类型.其基本类型是 int64.time.Minute 也是基于int64常量.
* 若同时声明一组常量,除了第一项之外,其他项在等号右侧的表达式都可以省略,这意味着回复用前一项的表达式及其类型
```
const (
    a = 1 // 1
    b     // 1
    c = 2 // 2
    d     // 2
)
```

### 常量生成器 iota
* iota 创建一系列相关值,而不是逐个值显示写出.常量声明中,itoa 从0开始取值,逐项加1
```
type Weekday int
const (
    Sunday Weekday = iota // 0 
    Monday                // 1
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)
```

### 无类型常量
* 从属类型待定的常量有6种,分别是无类型布尔 (true/false) 无类型整数 0 无类型文字符号 '\u0000' 无类型浮点数 0.0 无类型复数 0i 无类型字符串 (字符串字面量)
* 借助推迟确定从属类型,无类型常量不仅能暂时维持更高的精度.与类型已经确定的常量相比,它们还能写进更多表达式而无需转换类型. math.Pi
* 变量的声明(包括短变量声明)中,假如没有显式指定类型,无类型常量会隐式转换成该变量的默认类型. int rune float64 complex128 默认类型
* 各类型的不对称性.无类型整数可以转换成 int,其大小不确定,但是无类型浮点数和无类型复数被转换成大小明确的 float64和 complex128
* 想要将变量转换成不同的类型,我们必须将无类型常量显式转换为期望类型,或在声明变量是指明想要的类型`var i = int8(0) var i int8 = 0`
* 

