package main

import (
	"fmt"
	"time"
)
// 常量
// 编译阶段就可以计算出表达式的值,不需要等到运行时,所有常量本质上都属于基本类型:数字 字符串 布尔值
// 常量值恒定,防止修改,
// 一组常量的命名方式:括号括起来,
// 常量定义了不使用,不会报错,但是变量不行,编译阶段就会报错
const (
	e = 2.71828
	pi = 3.141589
	IPv4Len = 4
)
// time.Duration  具名类型,基本类型是 int64, time.Mint
const noDelay time.Duration = 0
const timeout = 5 * time.Minute

const (
	// 若同时声明一组常量,除了第一项外,其他项在等号右侧的表达式都可以省略,这个意味着会复用前一项的表达式和类型
	a = 1
	b // 复用 a
	c = 2 
	d // 复用 c
)
// 常量生成器 iota, 可以创建一系列的相关值,,而不是逐个显式写出, 常量声明中 iota 从0开始取值,逐项+1
// time 包 Weekday
type Weekday int // Weekday 的具名类型
const (
	Sunday Weekday = iota // 9
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)
// 无类型常量 这种比基本类型的数字精度更高,精度至少达到256位,6种:无类型布尔(true false) 无类型整数(0) 无类型浮点数(0.0) 无类型复数(0i) 无类型文字符号('\u0000') 无类型字符串

func main() {
	// var p [IPv4Len]byte // 数组类型的长度
	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)
	fmt.Println(a, b, c, d)

	fmt.Println(Sunday,
		Monday,
		Tuesday,
		Wednesday,
		Thursday,
		Friday,
		Saturday)
}