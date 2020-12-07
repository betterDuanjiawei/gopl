package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)
//go run du2/main.go -v /usr /bin /etc
var verbose = flag.Bool("v", false, "show verbose progress messages")
func main()  {
	// 确定初始目录
	flag.Parse()
	roots := flag.Args()
	//os.Args[1:]
	if len(roots) == 0 {
		roots = []string{"."}
	}
	// 遍历文件树
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	// 定期输出结果
	var tick <-chan time.Time // 变量 tick 类型是单向(接收)通道 类型是 time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	var nfiles, nbytes int64
	loop: // 只有 loop:
		for {
			select {
			case size, ok := <-fileSizes:
				if !ok {
					break loop // 标签化的 break语句,将跳出 select 和 for 循环,没有标签的 break, 只能跳出 select,导致循环的下一次迭代. fileSizes 关闭,跳出
				}
				nfiles++
				nbytes += size
			case <-tick :
				printDiskUsage(nfiles, nbytes)
			}
		}
		printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfile, nbytes int64)  {
	fmt.Printf("%d files %.1fGB\n", nfile, float64(nbytes)/1e9)
}

func walkDir(dir string, fileSizes chan<- int64 )  {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size() // 文件,发送一条消息到通道,消息是文件所占的字节数
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir) // ioutil.ReadDir()返回一个os.FileInfo类型的slice.针对单个文件同样的信息可以通过调用 os.Stat()来返回.
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}