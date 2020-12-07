package main

import (
	"fmt"
	"log"
	"net/http"
)

func main()  {
	db := database{"shote" : 5, "sock" : 50}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

type dollars float32

func (d dollars) String() string  {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars
// http.ResponseWriter 也是一个接口,它扩充了 io.Writer,加了发送 HTTP 响应头的方法,用 http.Error这个工具也可以达到同样的目的
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	// 基于 URl的路径部分:req.URL.Path
	switch req.URL.Path {
	case "/list" :
		for item, price := range  db {
			fmt.Fprintf(w, "%s:%s\n", item, price)
		}
	case "/price" :
		item := req.URL.Query().Get("item") // req.URL.Query()解析为一个 map(multimap), url.Values
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound) //返回一个 HTTP错误 404
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}