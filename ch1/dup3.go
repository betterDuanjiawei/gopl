package main

import (
	"fmt"
	"os"
	// "io"
	"io/ioutil"
	"strings"
)
// io/ioutil ReadFile 读取整个命名文件的内容 strings.Split 将一个字符串通过分隔符 分隔为一个 slice
func main() {
	counts := make(map[string]int)
	filenames := os.Args[1:]
	for _, filename := range filenames {
		file, err := ioutil.ReadFile(filename) // 读取出来的时[]byte 类型
		if err != nil {
			fmt.Printf("read %s failed, err: %v", filename, err)
			continue
		}
		fileSlice := strings.Split(string(file), "\n")
		for _, line := range fileSlice {
			counts[line]++
		}
	}

	for line, num := range counts {
		if num > 1 {
			fmt.Printf("%d\t%s\n", num, line)
		}
	}
}