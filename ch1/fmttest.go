package main

import (
	"fmt"
)

func main()  {
	fmt.Print("hello world")
	fmt.Println("你好 世界")
	s := fmt.Sprintf("冬天的%s火锅", "第一场")
	print(s)
	// s1 := fmt.Sprint("嘻嘻", "哈嘿", 100, "哈哈", true)
	s1 := fmt.Sprint("abc", "def", 100, true,  "ghy")
	println(s1)
}