package main

import (
	"fmt"
)
// 赋值
func main() {
	// 赋值语句用来更新变量所值的值,最简单的形式: 赋值符号= 和左边的变量 和右边的表达式组成
	// x = 1 // 有名称的变量
	// *p = true // 间接变量
	// person.name = "djw" // 结构体成员
	// conut[x] = count[x] * scale // 数组或 slice 或 map 的元素
	// count[x] *= scale // 避免了在表达式中重
	// v := 1 // 数字变量也可通过++和-- 进行递增和递减
	// v++ // v = v + 1
	// v-- // v = v - 1

	// 多重赋值 允许一个变量一次性被赋值. 在实际更新变量前,右边的所有表达式被推演,当变量同时出现在赋值符的两侧时候 这种形式特别有用.
	// 两个变量交换值时候:
	var x int = 1
	var y int = 2
	x, y = y, x // 右边的先推演
	fmt.Println(x, y)

	z := gcd(10, 50)
	fmt.Println(z)
}
// 取两个整数的最大公约数
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}

	return x
}
// 计算斐波那契数列的第 n 个数
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++{
		x, y = y, x+y
	}
	return x
}