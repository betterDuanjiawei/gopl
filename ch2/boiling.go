package main

import (
	"fmt" // 包名总是小写
)

// 包级别的实体名字不仅对于包含其声明的源文件可见,而且对于同一个包里面的所有源文件都可见
const boilingF = 212.0 // const 常量 包级别的声明 

// 函数包括一个名字 一个参数列表 一个可选的返回值列表 以及函数体
func main()  {
	// f c 是函数的局部比变量
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %gF or %gC\n", f, c) // %g 浮点数
}
