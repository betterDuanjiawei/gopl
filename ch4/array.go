package main

import (
	"fmt"
	"crypto/sha256"
)
// array
// 新数组中的元素初始值为元素类型的零值,数字:0
func main() {
	var q [3]int = [3]int{1, 2, 3}  // 数组字面量来初始化一个数组
	var r [3]int = [3]int{1, 2}
	fmt.Println(q, r[2])
	
	q1 := [...]int{1, 2, 3, 4, 5} // 如果省略号...出现在数组长度的位置,那么数组的长度由初始化数组的元素个数决定
	fmt.Printf("%T\n", q1) // %T 任何值的类型 [5]int

	// 数组的长度是数组类型的一部分,所以[3]int 和 [4]int是两种不同的数组类型,数组的长度必须是常量表达式,这个表达式的值在编译阶段就可以确定
	q2 := [3]int{1, 2, 3}
	// q2 = [4]int{1, 2, 3, 4} // cannot use [4]int literal (type [4]int) as type [3]int in assignment
	fmt.Println(q2)
	// go 数组 索引数组, 索引都是数字  0 1 2...
	r1 := [...]int{99 : -1} // 100个元素,前99个是默认值0 ,最后一个是-1
	fmt.Println(r1)

	// c1 := sha256.Sum256([]byte('x')) // cannot convert 'x' (type rune) to type []byte 
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1) // %x 十六进制 %t bool %T 类型
// 	2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
// false
// [32]uint8
// 'x'和 "x" 不同
	x1 := 'x'
	x2 := "x"
	// x1 == x2 // invalid operation: x1 == x2 (mismatched types rune and string)
	fmt.Printf("%v, %v,  %T, %T", x1, x2,  x1, x2) // 20, x,  int32, string
}