package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Commencing countdown.")
	// time.Tick 返回一个通道,它定期发送事件,像一个节拍器一样.每个事件的值是一个时间戳
	tick := time.Tick(2*time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	launch()
}
func launch() {
	fmt.Println("Lift off!")
}