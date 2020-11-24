package main

import (
	"fmt"
	"log"
	"net/http"
)

// 更加详细的输出请求信息
func main()  {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "URL:%s, Method:%s, 协议:%s\n", r.URL.Path, r.Method, r.Proto)
	for k,v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}

	for k,v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}