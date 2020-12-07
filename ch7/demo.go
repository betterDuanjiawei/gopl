package main

import (
	"bytes"
	//"bytes"
	"fmt"
	"io"
	"os"
	//"time"
)

func main()  {
	//demo1()
	//var x interface{} = []int{1, 2, 3}
	//fmt.Println(x == x) //  comparing uncomparable type []int
	demo3()
}

func demo1() {
	var w io.Writer
	w = os.Stdout // *os.File
	w = new(bytes.Buffer) // *bytes.Buffer
	//w = time.Second
	fmt.Printf("%T\n", w)

	var rwc io.ReadWriteCloser
	rwc = os.Stdout // *os.File 有 Read() Write() Close()方法
	//rwc = new(bytes.Buffer) // *bytes.Buffer缺少 Close 方法
	fmt.Printf("%T\n", rwc)

	w = rwc // ok, io.ReadWriteCloser有 write方法
	//rwc = w // io.Write 缺少 Close 方法

}
/*
func demo2() {
	var any interface{}
	any = true
	any = 12.34
	any = "hello"
	any = map[string]int{"one":1}
	any = new(bytes.Buffer)

	var w io.Writer = new(bytes.Buffer)
	var _ io.Writer = (*bytes.Buffer)(nil) // 将 nil 转换为指针类型的 bytes.Buffer
}

 */

func demo3()  {
	var w io.Writer
	fmt.Printf("%T\n", w) // <nil>

	w = os.Stdout
	fmt.Printf("%T\n", w) //*os.File

	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w) //*bytes.Buffer
}