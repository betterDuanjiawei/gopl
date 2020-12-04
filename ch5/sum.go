package main

import "fmt"

func main()  {
	fmt.Println(sum())
	fmt.Println(sum(3))
	fmt.Println(sum(1, 2, 3, 4, 5))

	values := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(values...)) //当实参已经存在于一个 slice 中的时候在最后一个参数后面放一个省略号...,
}

func sum(vals ...int) int  {
	total := 0
	for _, val := range vals { // 在函数体内 vals 是一个 int类型slice
		total += val
	}
	return total
}