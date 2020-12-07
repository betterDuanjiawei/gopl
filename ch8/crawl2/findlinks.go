package main

import (
	"fmt"
	"gopl.io/ch5/links"
	"log"
	"os"
)
//  go run crawl2/findlinks.go http://gopl.io/
func main()  {
	worklist := make(chan []string)
	// 计数器 n跟踪发送到任务列表的个数
	var n int
	n++
	go func() {
		worklist <- os.Args[1:]
	}()
	seen := make(map[string]bool)
	for ; n >0; n-- { // 如果第一个条件为空,那么就是;
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}

		}
	}
}

var token = make(chan struct{}, 20)
func crawl(url string) []string {
	fmt.Println(url)
	token <- struct{}{} // 获取令牌
	list, err := links.Extract(url)
	<-token
	if err != nil {
		log.Print(err)
	}
	return list
}