package main

import (
	"fmt"
	// "os"
)
// 包级别的初始化在main 开始之前进行,局部变量初始化和声明在函数执行期间进行
func main() {
	// 变量声明通用形式
	// var name type  = expression 可以只有左边是声明, 如果加上右边的表达式那么就是声明+初始化了
	// 类型和表达式可以省略,但是不能同时省略, 如果类型省略,那么它的类型将由初始化表达式决定,如果表达式省略,那么其初始值对应于类型的零值
	// 零值: 对于数字是0 对于 bool类型的是 false 对于string 是 "", 对于接口和引用类型(slice 指针 map 通道 函数)是 nil,对于一个数组或者结构体这样的复合类型,零值是其所有元素或成员的零值
	var i int
	var s string
	var b bool
	var sl []int
	var pt *int
	var m map[string]int
	var c chan(int)
	var f func()
	var arr [3]int
	type djw struct {} 

	// fmt.Printf("i : %d, s : %s, b : %t, sl : %v, pt %v , m : %v, c : %v, f : %v, arr : %v, djw : %v", i, s, b, sl, pt, m, c, f, arr, djw{})
	fmt.Printf("i : %d, s : %q, b : %t, sl : %v, pt %v , m : %v, c : %v, f : %v, arr : %v, djw : %T\n", i, s, b, sl, pt, m, c, f, arr, djw{})
	// i : 0, s : "", b : false, sl : [], pt <nil> , m : map[], c : <nil>, f : <nil>, arr : [0 0 0], djw : main.djw

	// 变量列表
	// var v1, v2, v3 int
	// var b1, f1, s1 = true, 2.3, "four" // 忽略类型声明允许声明多个不同类型的变量

	// var f, err = os.Open(name) var 和 := 不能同时出现
	// 短变量声明 : name := expression,name 的类型由 expression 决定
	//  因其短小 灵活,故在局部变量的声明和初始化中主要使用短声明, var声明通常是为那些跟初始化表达式类型不一致的局部变量保留的,或者用于后面才对变量赋值以及变量初始值不重要的情况
	// := 表示声明, = 表示赋值,一个多变量的声明不能和多重赋值搞混,后者将右边的值赋值给左边的对应变量
	// i1, j1 := 0, 1
	// i1, j1 = j1, i1
	// f2, err := os.Open(name)
	// if err != nil {
	// 	return err
	// }
	// defer f2.Close()
	// in, err := os.Open(infile)
	// out, err := os.Create(outfile) // 短变量声明最少声明一个新变量,否则,代码编译将无法通过
	// err := xxx 错误的,因为上面的 err,已经声明的变量了,你下面只能赋值了

	// 指针 指针类型的零值是 nil,
	// 指针的值是一个变量的地址,一个指针指示值所保存的位置,不是所有的值都有地址,但是所有的变量都有地址,使用指针,可以在无需知道变量名字的情况下,间接读取或更新变量的值.
	// & 取地址符 * 间接引用符,获取指针指向的变量值
	x2 := 1
	p := &x2 // p 是整形指针 指向x2 *p是x2的别名, 指针别名允许我们通过不同变量的名字来访问变量.
	if p != nil {
		fmt.Println(p) // != nil,说明 p 指向一个变量,指针是可比较的.
	}
	fmt.Println(p, *p)// p 是地址值, *p 是指针指向的变量值
	*p = 2 // 等于 x2 = 2
	fmt.Println(*p, x2)
	// 每一个聚合类型变量的组成(结构体的成员或数组中的元素)都是变量,所以也由一个地址
	var x3, y3 int
	var p2 *int
	var p3 *int
	fmt.Println(&x3 == &x3, &x3 == &y3, &x3 == nil, p2 == p3)// 两个指当且仅当指向同一个变量或者两者都是 nil的情况下才会相等(x3和 y3虽然都是0值,但是他们是两个不同的变量,不相等, &x3有指向的变量 x3,所以也不等于 nil,只有没有指向变量时候才是 nil) p2和 p3都是 nil,所以相等

	var p4 = fc()
	var p5 = fc()
	
	fmt.Println(p4, p5, p4 == p5) // 0xc000136040 0xc000136048 false 每次调用 fc都会返回一个不同值
	// 因为一个指针包含变量的地址值,所以传递一个指针参数给函数,能够让函数更新间接传递的变量值.
	v2 := 1
	incr(&v2)
	fmt.Println(incr(&v2)) // return的其实是 v2

	//
	
}
func fc() *int {
	v := 1 // 函数返回局部变量的地址是非常安全的, 通过调用 f 产生的局部变量v即使在调用返回后依然存在,指针 p依然引用 它
	return &v
}

func incr(p *int) int {
	*p++ // 递增 p 所指向的值,p 本身保持不变
	return *p
}