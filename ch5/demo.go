package main

import (
	"fmt"
	"os"
)

func main()  {
	f := square
	fmt.Println(f(3))
	f = negative
	fmt.Println(f(5))
	//f = product // cannot use product (type func(int, int) int) as type func(int) int in assignment

	var fc func(int) int
	fc(3) // panic: runtime error: invalid memory address or nil pointer dereference
	if fc != nil {
		fc(3)
	}
}

func square(n int) int {
	return n*n
}

func negative(n int) int  {
	return -n
}

func product(m, n int) int {
	return m*n
}

//捕获迭代变量
var rmdirs []func()
func bhbl() {
	for _, d := range tempDirs() {
		dir := d // 在循环体内将循环变量赋给一个新的局部变量 dir.
		os.MkdirAll(dir, 0755)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir)
		})
	}

	for _, dir := range tempDirs() {
		os.MkdirAll(dir, 0755)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir) // 不正确,匿名函数可以获取到外部的变量 dir变量的实际取值是最后一次迭代时的值并且所有的 os.RemoveAll 调用最终都试图删除同一个目录
		})
	}
	//
	for _, rmdir := range rmdirs {
		rmdir()
	}
}
