package main

import (
	// "fmt"
)

func main() {
	var f func(int) int
	f(3) // 宕机:调用空函数 函数类型的零值是 nil
// 	panic: runtime error: invalid memory address or nil pointer dereference
// [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x1056f47]

// goroutine 1 [running]:
// main.main()
//         /Users/v_duanjiawei/go/src/gopl.io/ch5/func.go:9 +0x27
// exit status 2
	recover() // 会终止当前的宕机状态并且返回宕机的值,如果 recover在其他任何情况下运行,则它没有任何效果而且返回 nil p := recover()
	panic() // 无论是 Go 语言底层抛出 panic，还是我们在代码中显式抛出 panic，处理机制都是一样的：当遇到 panic 时，Go 语言会中断当前协程中（main 函数）后续代码的执行，然后执行在中断代码之前定义的 defer 语句（按照先入后出的顺序），最后程序退出并输出 panic 错误信息，以及出现错误的堆栈跟踪信息，在这里就是：
	// 遇到第一个 panic 程序就中断了，不再往后执行了，管你后面天王还是老子

}