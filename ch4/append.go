package main

import "fmt"

func main()  {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

	var z []int
	z = append(z, 1)
	z = append(z, 2, 3)
	z = append(z, 4, 5, 6)
	z = append(z, z...) // 追加 z中的所有元素 参数后的...表示如何将一个slice转换为参数列表
	fmt.Println(z)

	var k []int
	k = appendInt(k, 1)
	k = append(k, 2, 3, 4)
	k = appendMInt(k, 2, 3, 4, 5, 6, 7, 8)
	fmt.Printf("%v\t%[1]T\t%d\t%d\n", k, len(k), cap(k))
}

func appendInt(x []int, y int) []int  {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		// 为了达到分摊线性复杂性,容量扩展一倍
		if zcap < 2 * len(x){
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // 内置 copy 函数
	}
	z[len(x)] = y

	return z
}

func appendMInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen

		for n := 1; zcap < n * len(x) ; n++ {
			zcap = n*len(x)
		}
		//if zcap < 2 * len(x) {
		//	zcap = 2*len(x)
		//}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)

	return z
}