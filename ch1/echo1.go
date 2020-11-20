package main

import (
	"fmt"
	"os"
)

func main() {
	var s, step string
	for i := 1; i < len(os.Args); i++ {
		s += step + os.Args[i] // 这里 step 和 参数的位置 顺序要注意,颠倒了效果就不一样了
		step = " "
	}
	// for _, arg := range os.Args[1:len(os.Args)] {
		
	// 	// if k == len(os.Args) - 1 {
	// 	// 	step = ""
	// 	// }
	// 	// s += arg + step
	// 	s += step + arg
	// 	step = " "
	// }
	fmt.Println(s+"go")
	fmt.Println(os.Args[0]) // go run echo1.go 的结果/var/folders/gd/6znpcr_57pb7k9s59h0395_w0000gp/T/go-build366445065/b001/exe/echo1
		// ./echo1  的结果 ./echo1	
}