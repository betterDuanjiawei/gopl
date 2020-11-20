package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main()  {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		log.Print(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		log.Printf("fetch %s statuscode is %s", url, resp.Status)

		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		resp.Body.Close()
	}
}

