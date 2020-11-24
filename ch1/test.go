package main

import "fmt"

func main()  {
	fmt.Println(Signum(1000))
}
// y := x++ 错误写法
func Signum(x int) int {
	switch {
	case x > 0:
		x++
		return x
		//return +1
	default:
		return x
	case x < 0:
		x--
		return x
		//return -1
	}
}