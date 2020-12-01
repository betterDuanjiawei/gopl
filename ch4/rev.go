package main

import "fmt"

func main()  {
	a := [...]int{1, 2, 3, 4, 5}
	reverse(a[:]) // 这里必须是a[:] 不应该是 a,因为 a是数组, 函数的参数是 slice
	fmt.Println(a)
	// 将一个 slice 左移n个元素的简单方式是连续调用 reverse函数3次.

	s := []int{0, 1, 2, 3, 4, 5}
}

func reverse(s []int)  {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}