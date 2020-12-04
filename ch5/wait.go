package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main()  {

	//if err := WaitForServer(strings.Join(os.Args[1:], "")); err != nil {
	//	fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
	//	os.Exit(1)
	//}
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tires := 0; time.Now().Before(deadline); tires++ {
		log.SetPrefix("wait...")
		log.SetFlags(0)
		log.Println(tires, time.Second, time.Second << uint(tires))
		time.Sleep(time.Second << uint(tires)) // 指数退避策略
		//2020/12/02 12:25:52 0 1s 1s
		//2020/12/02 12:25:53 1 1s 2s
		//2020/12/02 12:25:55 2 1s 4s
		//2020/12/02 12:25:59 3 1s 8s
		//2020/12/02 12:26:07 4 1s 16s

	}
}

func WaitForServer(url string)  error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tires := 0; time.Now().Before(deadline); tires++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tires)) // 指数退避策略
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}