package main

import (
	"fmt"
	"gopl.io/ch5/links"
	"log"
	"os"
)
//go run crawl1/findlinks.go http://gopl.io/
func main()  {
	worklist := make(chan []string)

	go func() {
		worklist <- os.Args[1:]
	}()

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) { // 循环变量捕获问题,显式传参
					worklist <- crawl(link)
				}(link)
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