package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Printf("%% %q %q %q", `a`, 'a', "a")  // % "a" 'a' "a"
	// var counts1 map[string]int // nil
	counts := make(map[string]int) // map[]
	// fmt.Printf("%T, %T",counts1, counts)
	input := bufio.NewScanner(os.Stdin)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}
}