/**
只 close1,后续 square中会继续 从 natural 中发送,所以会导致程序崩溃
只 close2,主 goroutine 中会继续<-squares 从中发送,也会导致程序崩溃
closet1 和 close2都 close,
*/
package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func ()  {
		for x := 1; x <= 10; x++ {
			naturals <- x
		}
		close(naturals) // 1
	}()

	go func ()  {
		// for {
		// 	x := <-naturals
		// 	
		// squares <- x * x
		// }
		// 变种接收
		for {
			x, ok := <-naturals
			// fmt.Printf("%d\t", x)
			if !ok {
				break
			}
			squares <- x * x
		}
		close(squares) // 2
	}()
	for {
		// 当关闭了通道之后,任何的发送操作都会导致程序崩溃
		fmt.Println(<-squares)
	}
	
}