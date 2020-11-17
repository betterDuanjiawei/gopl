package main

import (
	"fmt"
)
//
func main() {
	fmt.Println(5.0/4.0)
	fmt.Println(5/4)
	f := 1e100
	i := int(f)
	fmt.Println(f, i)

	// %[1] 重复使用第一个操作数 %c 输出文字符号 %q 输出带单引号的 ''中间只能有一个字符
	ascii := 'a'
	unicode := '国'
	newline := '\n'

	fmt.Printf("%d %[1]c %[1]q \n", ascii)
	fmt.Printf("%d %[1]c %[1]q \n", unicode)
	fmt.Printf("%d %[1]q\n", newline)

}