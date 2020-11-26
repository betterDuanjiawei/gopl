package main

import (
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
)

func main() {
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)

	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)

	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii) //97 a 'a'
	fmt.Printf("%d %[1]c %[1]q\n", unicode) //22269 国 '国'
	fmt.Printf("%d %[1]c %[1]q\n", newline)
	//10
	//'\n'

	// float
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d, ex = %8.3f, %[2]g, %[2]e\n", x, math.Exp(float64(x)))
	}

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // 0 -0 +Inf(正无穷大) -Inf(负无穷大) NaN(not a number, 如0/0, Sqrt(-1))
	fmt.Println(math.IsNaN(z), math.IsNaN(z/z), math.NaN(), math.NaN() == math.NaN(), math.NaN() == z/z, math.NaN() != z/z) // false true NaN false false true
	//var test string = "0" // cannot convert test (type string) to type int
	var test float64  =  0.0
	fmt.Println(int(test))

	// complex
	var c complex128 = complex(1, 2)
	var d complex128 = complex(3, 4)
	fmt.Println(c, d, c*d, real(c*d), imag(c*d)) // (1+2i) (3+4i) (-5+10i) -5 10
	fmt.Println(1i * 1i) // (-1+0i) i2 = -1

	// string
	s := "hello 世界"
	fmt.Println(len(s)) // 12
	fmt.Println(s[:5]) // hello
	fmt.Println(s[7:]) // ��界
	fmt.Println(s[0]) // 104
	fmt.Println(s[6:9]) // 世
	//fmt.Println(s[0:len(s)+1]) // 宕机:下标越界 runtime error: slice bounds out of range [:13] with length 12
	t := s
	s = "你好 world"
	fmt.Println(t, s) // hello 世界 你好 world
	//s[0] = 'h' // cannot assign to s[0]
	//fmt.Println(s)

	const GoUsage = `Go is a tool for manaing Go source code.

	Usage:
		go command [argumens]
	...`
	fmt.Println(GoUsage)
	//Go is a tool for manaing Go source code.
	//
	//	Usage:
	//go command [argumens]
	//...
	//strings.HasPrefix()
	str := "hello, 世界"
	fmt.Println(len(str))
	fmt.Println(utf8.RuneCountInString(str)) // 统计文字符号个数

	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(s[i:]) // r 文字符号本身 size 按 utf-8编码所占用字节数.
		fmt.Printf("%d\t%c\n", i, r)
		i += size+1
	}

	// []rune
	str2 := "段佳伟"
	fmt.Printf("% x\n", str2)
	fmt.Println(str2)
	r2 := []rune(str2)
	fmt.Printf("%x\n", r2)
	fmt.Println(r2)
	fmt.Println(string(r2))
	fmt.Println(string(65))
	fmt.Println(string(0x4eac))
	fmt.Println(string(1234567))
//	e6 ae b5 e4 bd b3 e4 bc 9f
//	段佳伟
//	[6bb5 4f73 4f1f]
//	[27573 20339 20255]
//	段佳伟
//	A
//	京
//	�
	strings.HasPrefix()

}