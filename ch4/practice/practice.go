package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main()  {
	x := [...]int{1, 2, 3, 4, 5}
	//x = reverse(&x)
	//fmt.Println(x)
	fmt.Println(2%5)
	fmt.Println(5.0/4.0, 5/4, 5.0/4, 5/4.0) // 1.25 1 1.25 1.25

	fmt.Println(rotate(x[:], 2))

	s := []string{"a", "b", "b", "c", "d", "c", "c", "e", "a"}
	fmt.Println(distinct(s))
	fmt.Println(distinct2(s))

	// 4.6
	b := []byte("北京\t欢迎\n您")
	fmt.Printf("%s\n%v\n", b, b)
	ex4_6(b)
	fmt.Printf("%s\n%v\n", b, b)

	// 4.7
	b = []byte("北京欢迎您welcome")
	fmt.Printf("%s\n", b)
	fmt.Printf("%s\n", mbreverse(b))

	//
	//charcount()
	//charcount1()


	// 4.9
	words := wordfreq()
	for k, v := range  words {
		fmt.Printf("%s\t%d\n", k, v)
	}
}
// 4.3 重写reverse函数，使用数组指针作为参数而不是slice
func reverse(p *[5]int) [5]int {
	var v = *p
	for i, j := 0, len(*p)-1; i < j; i, j = i+1, j-1 {
		v[i], v[j] = v[j], v[i]
	}

	return v
}

// 4.4 编写函数rotate，实现一次遍历就可以完成元素旋转
func rotate(s []int, n int) (t []int) {
	if n <= 0 || n >= len(s) {
		return s
	}
	t = make([]int, len(s))
	for i, v := range s {
		t[(i+n)%len(s)] = v
	}
	return
}
// 4.5 编写一个就地处理函数，用于去除[]string slice中相邻重复字符串元素
func distinct(s []string) []string {
	news := make([]string, len(s))
	for k, v := range s {
		if k < len(s) - 1 && v == s[k+1] {
			continue
		}
		news[k] = s[k]
	}

	return news
}

func distinct2(s []string) []string {
	dist := s[0:1]
	for i := 1; i < len(s); i++ {
		//if s[i] != dist[i-1] {
		//	dist = append(dist, s[i])
		//}
		if s[i] != dist[len(dist)-1]{ // 这里不能是 i-1,因为 i和 dist的元素个数是不同步的
			//for s[i] != dist[len(dist)-1]{
			dist = append(dist, s[i])
		}
	}

	return dist
}

// 4.6 编写一个就地处理函数，用于将一个UTF-8编码的字节slice中所有相邻的Unicode空白字符（查看unicode.IsSpace）缩减为一个ASCII空白字符。
func ex4_6(b []byte) []byte {
	var i int
	for i, l := 0, 0; l < len(b); {
		r, size := utf8.DecodeRune(b[i:])
		//fmt.Printf("%v\t%s", r, b[i:])
		l += size
		if unicode.IsSpace(r) {
			if i > 0 && b[i-1] == byte(32) {
				copy(b[i:], b[i+size:])
			} else {
				b[i] = byte(32)
				copy(b[i+1:], b[i+size:])
			}
		} else {
			i += size
		}
	}
	return b[0:i]
}
// 4.7 修改函数reverse，来反转一个UTF-8编码字符串中的字符元素，传入参数是该字符串对应的字节slice类型（[]byte）。是否可以做到不重新分配内存就实现该功能。
func mbreverse(b []byte) []byte {
	var res []byte
	for i := len(b); i > 0; {
		r, size := utf8.DecodeLastRune(b[:i])
		fmt.Println(r, size)
		res = append(res, []byte(string(r))...)
		i -= size
	}

	return res
}

// 4.8
func charcount() {
	counts := make(map[rune]int)
	typecount := make(map[string]int)
	invalid := 0
	var utflen [utf8.UTFMax+1]int
	//f, err := ioutil.ReadFile("test.txt")
	//if err != nil {
	//	fmt.Printf("open err: %v\n", err)
	//	os.Exit(1)
	//}
	in := bufio.NewReader(os.Stdin)
	for {
		r, nbyte, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("read error:%v\n", err)
			os.Exit(1)
		}
		if unicode.IsLetter(r) {
			typecount["letter"]++
		} else if unicode.IsNumber(r) {
			typecount["number"]++
		} else {
			typecount["other"]++
			if r == unicode.ReplacementChar && nbyte == 1 {
				invalid++
				continue
			}
		}
		counts[r]++
		utflen[nbyte]++
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("typecount\tcount\n")
	for t, n := range typecount {
		fmt.Printf("%s\t%d\n", t, n)
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

// 4.8参考的别人的版本
func charcount1() {
	seen := make(map[rune]bool)
	var isLetter, isNumber, isOther, isInvalid int
	in := bufio.NewReader(os.Stdin)
	for {
		r, nbyte, err := in.ReadRune()
		if  err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("error:%v", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && nbyte == 1{
			isInvalid++
			continue
		}
		if unicode.IsLetter(r) && !seen[r] {
			seen[r] = true
			isLetter++
			continue
		}
		if unicode.IsNumber(r) && !seen[r] {
			seen[r] = true
			isNumber++
			continue
		}
		if !seen[r] {
			seen[r] = true
			isOther++
		}
	}

	fmt.Fprintf(os.Stdout, "letterCount:%d\tnumberCount:%d\totherCount:%d\tinvalidCount:%d\n", isLetter, isNumber, isOther, isInvalid)
}

// 4.9
func wordfreq() map[string]int{
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	words := make(map[string]int)
	for scanner.Scan() {
		words[scanner.Text()]++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	return words
}
