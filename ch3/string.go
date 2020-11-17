package main

import (
	"fmt"
	"unicode/utf8"
	"strings"
	"bytes"
	"strconv"
)

func main() {
	// 
	s := "heello, 世界" // 字节和字符是不一样的,一个中文字符占2个字节
	fmt.Println(len(s)) // 14个字节
	fmt.Println(utf8.RuneCountInString(s))// 10个字符
	n := 0
	for _, _ = range s { // 这里都是2个空标识符,没有新的变量,所以是 =,而不是 :=
		n++ 
	}
	fmt.Println(n)

	n1 := 0
	for range s { // 可以忽略没用的变量
		n1++
	}
	fmt.Println(n1)

	// rune 文字符号类型
	fmt.Println(string(65))
	fmt.Println(string(0x4eac))
	fmt.Println(string(1234567)) // 如果文字符号非法,将被专门的替换字符取代 \uFFFD

	// 字符串和字节 slice
	// 4个标准包 对于字符串操作特别重要,bytes strings strconv unicode
	// 字节 slice []byte 类型 bytes.Buffer类型更高效
	sl := "a/b/c.xxx.go"
	fmt.Println(basename(sl))
	fmt.Println(basename2(sl))

	// path(url 地址的路径部分) path/filepath (系统文件名处理)

	s3 := "abc"
	b3 := []byte(s3)
	s4 := string(b3)
	fmt.Println(s3, b3, s4) // abc [97 98 99] abc

	// 字符串和字节 slice(允许随意修改)相互转换 bytes strings 包
	// 包含函数 Contains() Count() Fields() HasPrefix() Index() Join() 两个包都有对应的函数,操作对象 字符串 字节 slice
	// bytes.Buffer变量无须初始化,原因是零值本来就有效
	fmt.Println(intsToString([]int{1, 2, 3}))

	// strconv 字符串和数字的相互转换
	// 将整数转换为字符串,fmt.Sprintf(把格式化的字符串输出到指定字符串) strconv.Itoa()
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))
	// FormatInt FormatUint 可以按不同的进制格式化数字
	fmt.Println(strconv.FormatInt(int64(x), 2))
	// sx := 
	fmt.Println(fmt.Sprintf("x = %b", x)) // %b 输出二进制的格式,而且还可以包含数字以外的附加信息

	// Itoa() Atoi()表示整数的字符串
	x1, err := strconv.Atoi("123")
	y1, err := strconv.ParseInt("123", 10, 16) // 第三个参数: 0表示 int
	fmt.Println(x1, y1, err) 

}
// 	Printf(),是把格式字符串输出到标准输出（一般是屏幕，可以重定向）。
// 　　Printf()是和标准输出文件(stdout)关联的,Fprintf则没有这个限制.
// 　　Sprintf(),是把格式字符串输出到指定字符串中，所以参数比printf多一个char*。那就是目标字符串地址。
// 　　Fprintf(),是把格式字符串输出到指定文件设备中，所以参数笔printf多一个文件指针FILE*。主要用于文件操作。Fprintf()是格式化输出到一个stream，通常是到文件。
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[') // bytes.Buffer 添加任意文字符号的 UTF-8编码,使用 WriteRune,追加 ASCII字符, 使用 WriteByte 可
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		// fmt.Println(i, v)
		fmt.Fprintf(&buf, "%d", v) // 把格式字符串输出到指定字符串中,所以参数比,printf 多了一个文件指针 FILE,主要用于文件操作,Fprintf()是格式化输出到一个 stream,通常是到文件
	}
	buf.WriteByte(']')

	return buf.String()
}

// basename
func basename(s string) string  {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' { // 单个字符可以用''
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

// strings.LastIndex 最后一次出现的索引位置
func basename2(s string) string  {
	slash := strings.LastIndex(s, "/") // 如果没有找到,则 slash 取值 -1
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0{
		s = s[:dot]
	}

	return s
}

func comma(s string) string  {
	n := len(s)
	if n <= 3 {
		return s
	}
	// 只处理已经加了,前面的
	return comma(s[:n-3]) + "," + s[n-3:]
}

// func comma2(s string) string {
// 	var buff bytes.Buffer = 
// }