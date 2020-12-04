package main

import (
	"fmt"
	"log"
	"time"
)

func main()  {
	bigSlowOperation()
	double(4)
	fmt.Println(triple(4))
}

func bigSlowOperation()  {
	log.Printf("enter %s", "yes")
	time.Sleep(2 * time.Second)
	defer trace("bigSlowOperation")()
	//defer func() func(){
	//	start := time.Now()
	//	msg := "bigSlowOperation"
	//	log.Printf("enter %s", msg)
	//	return func() {
	//		log.Printf("exit %s (%s)", msg, time.Since(start))
	//	}
	//}()
	time.Sleep(10 * time.Second)

}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func double(x int) (result int) {
	defer func() {
		fmt.Printf("double(%d) = %d\n", x, result)
	}()
	return x + x
}

func triple(x int) (result int) {
	defer func() { result += x}()
	return double(x)
}