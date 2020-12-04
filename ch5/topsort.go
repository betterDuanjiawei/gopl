package main

import (
	"fmt"
	"go4.org/sort"
)

var prereqs = map[string][]string{
	"algorithms" : {"data structures"},
	"calculus" : {"linear algebr"},
	"compilers" : {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures" : {"discrete math"},
	"databases" : {"data structures"},
	"discrete math" : {"intro to programming"},
	"formal languages" : {"discrete math"},
	"networks" : {"operating systems"},
	"operating system" : {"data structures", "computer organization"},
	"programming languages" : {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	// 当一个匿名函数需要进行递归,必须线声明一个变量,然后将匿名函数赋给这个变量.如果将这两个变量合成一个声明.函数字面量将不能存在于 visitAll 变量的作用域中,这样也就不能递归调用自己了
	/**
	visitAll := func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	 */
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}