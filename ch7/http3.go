package main

import (
	"fmt"
	"log"
	"net/http"
)

func main()  {
	db := database{"shote" : 5, "sock" : 50}
	mux := http.NewServeMux()
	//mux.Handle("/list", http.HandlerFunc(db.list)) // 将 db.list 函数转化为 http.HandleFunc类型
	//mux.Handle("/price", http.HandlerFunc(db.price))
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type dollars float32

func (d dollars) String() string  {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request)  {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request)  {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %s\n", item)
		return //要记得加
	}
	fmt.Fprintf(w, "%s : %s\n", item, price)
}