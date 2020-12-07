package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
//go run du4/main.go -v / /usr /bin /etc
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
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes) // 这里是& 指针
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()
	// 当检测到输入的时候,取消遍历
	go func ()  {
		os.Stdin.Read(make([]byte, 1))
		close(done)
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
		case <-done:
			// 耗尽fileSize 通道,丢弃他的所有值,直到通道关闭,这么做是为了保证所有的walkDir调用可以执行完,不会卡在向 fileSizes通道发送消息上
			for range fileSizes {
				// 不执行任何操作
			}
			return
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

func walkDir(dir string, n *sync.WaitGroup,fileSizes chan<- int64 )  {
	defer n.Done()
	// 轮询取消状态,如果设置状态,什么都不做,立马返回
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1) // 这里调用了也不能少
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes) // 这里也有 go
		} else {
			fileSizes <- entry.Size() // 文件,发送一条消息到通道,消息是文件所占的字节数
		}
	}
}
var sema = make(chan struct{}, 20)
// 高峰期会有很多很多 goroutine,
func dirents(dir string) []os.FileInfo {
	//sema <- struct{}{}
	select {
	case sema <- struct{}{}: // 获取令牌
	case <-done:
		return nil
	}
	defer func() { // 释放令牌
		<-sema
	}()
	// defer <-sema 错误写法
	entries, err := ioutil.ReadDir(dir) // ioutil.ReadDir()返回一个os.FileInfo类型的slice.针对单个文件同样的信息可以通过调用 os.Stat()来返回.
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}