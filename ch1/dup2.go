package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if (len(files) == 0) {
		countLine(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file) // file只是一个字符串的文件名,需要打开, 返回 *os.File, error(nil时候打开成功)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2 %s open err: %v", file, err)
				continue //打开文件错误,继续下一个文件
			}
			countLine(f, counts)
			f.Close() // 关闭文件,释放资源
		}
	}

	for line, num := range counts {
		if num > 1 {
			fmt.Printf("%d\t%s\n", num, line)
		}
	}
}

// 公用部分提取出来 流式模式读取输入,按需拆分为行, 可以处理海量输入
func countLine(f *os.File, counts map[string]int)  {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}