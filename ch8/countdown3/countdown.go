package main

import (
	"fmt"
	"time"
)

func main() {
	demo2()
	// 创建终止通道
	/*abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown. pleas return to abort")
	tick := time.Tick(1*time.Second)
	for countdown := 0; countdown < 10; countdown++ {
		fmt.Println(countdown)
		select {
		case <-tick:
			//
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()*/
}
func launch() {
	fmt.Println("Lift off!")
}

func demo()  {
	ticker := time.NewTicker(1*time.Second)
	<-ticker.C
	ticker.Stop()
}
func demo2()  {
	abort := make(chan struct{})
	select {
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	default:
		//fmt.Println("default")
	}
	fmt.Println("end") // 跑一次的结果是:end,如果想一直监听,需要循环
}
