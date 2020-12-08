package memotest

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

func HttpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

type M interface {
	Get(key string) (interface{}, error)
}
func Sequential(t *testing.T, m M)  {
	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
}

func incomingURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			"https://www.taobao.com",
			"https://www.jd.com",
			"https://www.baidu.com",
			"https://www.qq.com",
			"https://www.taobao.com",
			"https://www.jd.com",
			"https://www.baidu.com",
			"https://www.qq.com",
		} {
			ch <- url
		}
		close(ch) // 这里发送完就关闭了
	}()
	return ch
}

func Concurrent(t *testing.T, m M)  {
	var n sync.WaitGroup
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
}