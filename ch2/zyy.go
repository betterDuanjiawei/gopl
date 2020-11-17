package main

import (
	"fmt"
	"os"
	"log"
)

var cwd string

func initt() {
	cwd, err := os.Getwd() // 获取当前工作目录 不是赋值给 cwd
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	log.Printf("%s", cwd)
	// fmt.Println(cwd)
}

func inittt() {
	var err error
	cwd, err = os.Getwd() // 实现了给全局变量 cwd 赋值
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	log.Printf("%s", cwd)
	// fmt.Println(cwd)
}

func main() {
	// x := 8;
	// y := 9;
	if  x := 8; x == 0 { // x 和 y 的作用域只在循环体内,所以将 x := 8;放在上一行,不仅仅是写法简单紧凑,而且变量的作用域也不同
		fmt.Println(x)
	} else if  y := 9; y == x { // x 在这里是可以作用到得
		fmt.Println(y)
	} else {
		fmt.Println("default")
	}
	// fmt.Println(x, y)
/**
词法块
	if f, err := os.Open(name); err != nil { // f 未使用
		return err
	}
	f.Stat() // f 变量不存在
	f.Close()

	// 最好的写法 在 if块中处理错误然后返回,这样成功执行的路径不会变得支离破碎
	f, err := os.Open(name);
	if (err != nil) {
		return err
	}
	f.Stat()
	f.Close()

	if f,err := os.Open(name); err != nil {
		return err
	} else {
		f.Stat()
		f.Close()
	}
*/
	initt()
	fmt.Println(cwd)
	inittt()
	fmt.Println(cwd)
}