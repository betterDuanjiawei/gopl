package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	// "io"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("get %s failed, err: %v", url, err)
			os.Exit(1)
		}
		content, err := ioutil.ReadAll(resp.Body) // 读取整个响应结果
		resp.Body.Close() // 关闭连接资源
		if err != nil {
			fmt.Println(err)
			continue
		}
		// content.close()
		fmt.Printf("%s", content)
	}
}