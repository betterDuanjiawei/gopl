package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main(){
	counts := make(map[rune]int)
	invalid := 0
	var utflen [utf8.UTFMax+1]int // 为什么这里是 utf8.UTFMax+1呢?因为数组的索引是从0开始的,而 utf8的字节长度是1-4,所以需要+1

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // 返回解码的字符, utf-8编码的中的字节长度和错误值. 错误有可能是文件结束 io.EOF,如果输入的是不合法的 utf-8字符,那么返回的字符是 unicode.ReplacementChar 并且长度是1
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("len\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("%d invalid UTF-8 characters\n", invalid)
	}

	fmt.Print(utflen)
}