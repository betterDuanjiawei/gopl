package main

import (
	"fmt"
)

func main() {
	const freezingF, boilingF = 32.0, 212.0 // 常量是 =  变量是 :=
	fmt.Printf("%gF = %gC\n", freezingF, ftoc(freezingF))
	fmt.Printf("%gF = %gC\n", boilingF, ftoc(boilingF))
}

func ftoc(f float64) float64 { // 如果直接返回不需要定义返回变量值 ,如果写了c float64 下面就必须有函数体, c = (f- 32) * 5 /9  return c
	return (f - 32) * 5 / 9
}