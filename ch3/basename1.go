package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(basename("a/b/c.go"))
	fmt.Println(basename2("a.b/c/d.go.go"))
}

func basename(s string)  string {
	for i := len(s) - 1; i >= 0 ; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // 如果没找 ,slash取值-1
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}