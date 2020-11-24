package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex // 高并发,为了保证同一时间只有一个 goroutine 在访问变量
var count int
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/counter", counter)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Paht = %s", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprint(w, count)
	mu.Unlock()
}