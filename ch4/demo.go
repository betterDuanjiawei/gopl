package main

import (
	"fmt"
	"sort"
)

func main()  {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a) - 1])
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for _, v := range a {
		fmt.Println(v)
	}

	var r [3]int = [3]int{1, 2}
	fmt.Println(r)

	m := map[string]int{"a": 1, "c":2, "d": 5,  "b": 8}
	ssort(m)

	// 零值
	var test1 map[string]int
	fmt.Println(test1==nil) // true
	fmt.Println(len(test1)==0) // true
	test1["age"] = 1
	fmt.Println(test1) // panic: assignment to entry in nil map
}

// slice的比较
func equal(x, y []string)  bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}

	return true
}

// 值为 int类型的 map的比较
func mEqual(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv,ok := y[k]; !ok || yv != xv {
			return false
		}
	}

	return true
}

func ssort(m map[string]int) {
	//var s []string
	s := make([]string, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	sort.Strings(s)
	for _, v := range s {
		fmt.Println(m[v])
	}
}

