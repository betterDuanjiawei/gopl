package main

import (
	"fmt"
	"os"
	"time"
)

func main()  {
	demo()
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // 读取单个字节
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown. pleas return to abort")
	select {
	// time.After 在立即返回一个通道,然后启动一个新的goroutine在间隔指定时间之后,发送一个值到它上面
	case <- time.After(10*time.Second):
		//不执行任何操作
	case <- abort:
		fmt.Println("Launch aborted!")
		return
	}
	launch()
}
func launch() {
	fmt.Println("Lift off!")
}

func demo()  {
	// ch 的缓冲区大小是1,他要么是空,要么是满,因此只有在其他的一种情况下执行.要么 i 是偶数时候发送,要么 i是奇数时候接收
	ch := make(chan int, 4)
	for i := 0; i < 10; i++ {
		select {
		case ch <- i:
			//
		case x := <-ch:
			fmt.Println(x)
		}

	}
}