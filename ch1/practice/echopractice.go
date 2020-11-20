package main

import (
	"fmt"
	"os"
	"time"
	"strings"
)

func main() {
	fmt.Println(os.Args[0])

	for k, v := range os.Args[1:] {
		fmt.Println(k)
		fmt.Println(v)
	}

	var slice1 []string
	for i := 1; i <= 1000; i++ {
		slice1 = append(slice1, string(i))
		// 不能写成 slice1[i] = i
	}
	fmt.Println(slice1)
	s, step := "", ""
	start1 := time.Now()
	for _, v := range slice1 {
		s += step + v
		step = " "
	}
	// end1 := time.Now()
	// optime := end1.Sub(start1)
	optime1 := time.Since(start1).Seconds()
	fmt.Println(s)
	fmt.Printf("操作耗时%v", optime1)

	var s2 string
	start2 := time.Now()
	s2 = strings.Join(slice1, " ")
	optime2 := time.Since(start2).Seconds()
	fmt.Println(s2)
	fmt.Printf("操作耗时%v", optime2)
}