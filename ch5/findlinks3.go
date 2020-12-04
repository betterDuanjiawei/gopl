package main

import (
	"fmt"
	"gopl.io/ch5/links"
	"log"
	"os"
)

//go run findlinks3.go http://golang.org
// go run findlinks3.go https://www.hao123.com/

// breadthFirst 对每一个 worklist 元素调用 f,并将返回的内容添加到 worklist 中,每一个元素,最多调用一次 f
func main()  {
	breadthFirst(crawl, os.Args[1:])
}
// 在爬虫中,项节点都是 url.
func breadthFirst(f func(item string) []string, worklist []string)  {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}
func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}