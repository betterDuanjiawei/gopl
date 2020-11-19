package main

import (
	"fmt"
	"sort"
)
// map 拥有键值对得无序集合
func main() {
	// var graph = make(map[string]map[string]bool) // 变量 graph的键类型是 string类型,值类型是 map 类型 map[string]bool
	ages := map[string]int{
		"djw" : 29,
		"lxq" : 28,
	}

	// var names []string
	names := make([]string, 0, len(ages)) // 指定长度更加高效,创建一个长度为0的空的,容量为 map长度的 slice

	for name, _ := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Println(ages[name])
	}

	var map1 map[string]string
	fmt.Println(map1==nil)
	fmt.Println(len(map1)==0)
	// fmt.Println(cap(map1)==0) // map 有长度但是没有容量 invalid argument map1 (type map[string]string) for cap
	map1 = map[string]string{} // 初始化了 值是 map[]
	fmt.Println(map1, map1==nil) // map[] false
	map1["djw"] = "djw" // panic: assignment to entry in nil map
	fmt.Println(map1) // map[djw:djw]
	
}

// 两个 map做比较
func equal(x, y map[string]int) bool {
	if (len(x) != len(y)) { // 先比较长度,再比较内容
		return false
	}
	for xk, xv := range x {
		if yv, ok := y[xk]; !ok || xv != yv{ // 先看是否由这个元素 再看是否相等
			return false
		}
	}

	return true
}