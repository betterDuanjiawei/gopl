package main

import (
	"fmt"
	"syscall"
)

func main()  {
	var err error = syscall.Errno(15)
	fmt.Println(err.Error())
	fmt.Println(err)
}