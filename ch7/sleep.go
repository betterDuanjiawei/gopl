package main

import (
	"flag"
	"fmt"
	"time"
)
/**
go run sleep.go -period 2
invalid value "2" for flag -period: parse error

go run sleep.go -period 2s ok
 */
var period = flag.Duration("period", 1* time.Second, "sleep period")
func main()  {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}