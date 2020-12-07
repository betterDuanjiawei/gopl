package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)
//  ./fetch http://www.w3.org/TR/2006/REC-xml11-20060816 | go run xmlselect.go div div h2
func main()  {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect : %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if containAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

// 注意这个函数和巧妙
func containAll(x, y []string) bool {
	for len(y) <= len(x) { //for 循环
		if len(y) == 0 {
			return true
		}
		// 有啥用呢?
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}