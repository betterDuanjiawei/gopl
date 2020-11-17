package main

import (
	"fmt"
	"flag"
	"strings"
)
// 指正 flag包 flag.Bool 函数创建一个新的布尔标识变量, 标识的名字, 变量的默认值, 以及当用户提供非法标识 非法参数 或者 -h -help参数时候输出的消息
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator") // flag.String 名字 默认值 消息

func main() {
	flag.Parse() // 更新标识变量的默认值
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println() // 输出空行
	}

	// new 函数
	// 表达式 new(T)创建一个未命名的 T类型变量,初始化为 T 类型的零值,并返回其地址(地址类型为*T)
	p := new(int) // 使用new创建的变量和取其地址的普通局部变量没有什么不同,只是不需要引入和声明一个虚拟的名字,通过 new(T)就可以直接在表达式中使用,new 只是语法上的便利,不是一个基础概念
	// p := new(1) 必须为类型,不能为具体的值,
	fmt.Println(p, *p)
	*p = 2
	fmt.Println(p, *p)

	// 每一次调用 new 返回一个具有唯一地址的不同变量
	p1 := new(int)
	q1 := new(int)
	fmt.Println(p1 == q1)
	// 例外:两个变量的类型不携带任何信息而且是零值,struct{} [0]int 他们具有相同的地址
	p2 := new(struct{})
	q2 := new([0]int)
	fmt.Println(p2, q2)

	// new 是一个预声明函数, 不是一个关键词,所以它可以重定义为另外的其他类型
	var new int
	new = 3
	fmt.Println(new)

	g := f()
	fmt.Println(global, g, *g)
	// 变量的生命周期

}
// 下面那两个函数效果相同
// func newInt() *int {
// 	return new(int)
// }

// func newIng() *int {
// 	var dummy int
// 	return &dummy
// }

var global *int
func f() *int {
	var x int
	x = 1
	global = &x // x 从 f()中逃逸, 在性能优化时候是由好处的,因为每一次变量逃逸都需要一次额外的内存分配空间
	return global
}