package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request)  {
	path := r.URL.Path
	fmt.Fprintf(w, "URL.Path = %s\n", path)
}