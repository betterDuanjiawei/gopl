package main

import (
	"fmt"
	"gopl.io/ch5/links"
	"os"
)

func main()  {
	worklist := make(chan []string) // 可能有重复的 url列表
	unseenLinks := make(chan string) // 去重后的url列表

	go func() {
		worklist <- os.Args[1:]
	}()

	for i := 0; i < 20; i++  {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				// 为啥这里要加 go呢? 如果不加 go,这里进行发送,但是 worklist 是满的,那么这里就会阻塞
				go func() {
					worklist <- foundLinks // crawl发现的链接通过精心设计的 goroutine发送到任务列表来避免死锁
				}()
			}
		}()
	}

	// 爬取 goroutine 使用同一个通道 unseenLinks进行接收,主 goroutine 负责从任务列表接收到的条目进行去重,然后发送每一条
	// 没有爬去过的 url到 unseenLinks 通道,然后被爬取goroutine接收
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}

}

func crawl(url string) []string  {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		fmt.Print(err)
	}
	return list
}