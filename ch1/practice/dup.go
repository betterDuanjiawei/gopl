package main

import (
	"fmt"
	"os"
	"bufio"
)
// 输出出现重复行的文件的名称
func main() {
	// 存放 line 出现的文件名和在该文件中出现的次数
	 dtlCounts := make(map[string]map[string]int)
	//  m := make(map[string]int)
	 filenames := os.Args[1:]
	 if len(filenames) == 0 {

	 } else {
		 for _, filename := range filenames {
			 f, err := os.Open(filename)
			 if err != nil {
				 fmt.Printf("open %s err: %v", filename, err)
			 }
			countLine(f, filename, dtlCounts)
		 }
	 }
	 for line, maps := range dtlCounts {
		 for fn, num := range maps {
			 if num > 1 {
				 fmt.Printf("%s在%s文件中 出现了%d次\n", line, fn, num)
			 }
		 }
	 }
	//  fmt.Println(dtlCounts)
}

func countLine(f *os.File, filename string, dtlCounts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		// 报错解决 panic: assignment to entry in nil map,需要初始化 map
		if dtlCounts[line] == nil {
			dtlCounts[line] = make(map[string]int)
		}
		dtlCounts[line][filename]++
	}
}