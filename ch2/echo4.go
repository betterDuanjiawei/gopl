// echo4输出其命令行参数
package main

import (
	"flag"
	"fmt"
	"strings"
)
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}

	p := new(int)
	q := new(int)
	fmt.Println(p == q) // false

	x := new(struct{})
	y := new([0]int)
	fmt.Println(x == y) // invalid operation: x == y (mismatched types *struct {} and *[0]int)

}