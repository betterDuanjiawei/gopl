package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n",  n, fibN)
}

func spinner(delay time.Duration)  {
	for {
		// for _, r := range `-\|/` { //  目前理解应该是不转义纯字符串的意思
		for _, r := range "=\\|/" { // 如果要用 "" 就必须转义\\
			fmt.Printf("\r%c", r) // 输出单个字符
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}