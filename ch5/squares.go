package main

import (
	"strings"
	"fmt"
)

func main() {
	testStr := strings.Map(func (r rune) rune {
		return r + 1
	}, "HAL-9000")
	fmt.Println(testStr)
	
	// f := squares()
	
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	// 里层函数使用外层函数的变量
}
// 返回一个函数 类型是 func() int
func squares() func() int {
	var x int // 第二次 第三次 的 x值是不是一样的呢?不懂
	return func () int {
		x++
		return x * x
	}
}