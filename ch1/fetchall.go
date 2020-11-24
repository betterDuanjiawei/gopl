package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main()  {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go getContent(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("总耗时: %f", time.Since(start).Seconds())
}

func getContent(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer resp.Body.Close()
	size, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	ch <- fmt.Sprintf("%s 获取大小为: %d, 获取耗时为: %f", url, size, time.Since(start).Seconds())

}