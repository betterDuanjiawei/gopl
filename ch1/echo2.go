package main

import (
	"fmt"
	"os"
)

func main() {
	s, step := "", "" // 一行多变量赋值初始化
	for _, arg := range os.Args[1:] {
		s += step + arg
		step = " "
	}
	fmt.Println(s)
}