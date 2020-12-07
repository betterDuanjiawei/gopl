package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main()  {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File) //&{0xc000124000} *os.File
	fmt.Printf("%v\t%[1]T\n", f)
	//c := w.(*bytes.Buffer)
	//fmt.Printf("%v\t%[1]T\n", c) // interface conversion: io.Writer is *os.File, not *bytes.Buffer

	demo2()
	demo3()
	demo4()
}

func demo2() {
	var w io.Writer
	w = os.Stdout
	rw := w.(io.ReadWriter)
	fmt.Printf("%T\t%T\n", w, rw) // *os.File        *os.File
}

func demo3()  {
	var w io.Writer = os.Stdout
	if f, ok := w.(*os.File); ok {
		fmt.Printf("%T\n", f) // *os.File
	}
	if b, ok := w.(*bytes.Buffer); !ok {
		fmt.Printf("%v%[1]T\n", b) // <nil>*bytes.Buffer
	}
	//os.IsExist()
	//os.IsNotExist()
	//os.IsPermission()
}

func demo4() {
	_, err := os.Open("no such file")
	fmt.Println(err)
	fmt.Printf("%#v\n", err)
	//open no such file: no such file or directory
	//&os.PathError{Op:"open", Path:"no such file", Err:0x2}
}

func sqlQuote(x interface{}) string {
	switch x := x.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", x)
	case bool:
		if x {
			return "TRUE"
		}
		return "FALSE"
	case string:
		return sqlQuoteString(x)
	default:
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}