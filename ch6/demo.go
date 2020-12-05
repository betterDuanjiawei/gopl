package main

import (
	"fmt"
	"gopl.io/ch6/geometry1"
)

func main()  {
	demo1()
	demo2()
}

func demo1()  {
	perim := geometry1.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())
	fmt.Println(geometry1.PathDistance(perim))
}
/**
type P *int

//  invalid receiver type P (P is a pointer type)
func (p P) test()  {
	fmt.Println("test")
}
 */
type Point struct {
	X, Y float64
}

func (p *Point)ScaleBy(factor float64)  {
	p.X *= factor
	p.Y *= factor
}

func demo2()  {
	// 通过*Point调用*Point.ScaleBy
	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r)

	// 第二种写法
	p := Point{1, 2}
	pptr := &p
	pptr.ScaleBy(2)
	fmt.Println(p)

	q := Point{1, 2}
	(&q).ScaleBy(2)
	fmt.Println(q)

	// 如果接收者 q是 Point类型的变量,但是方法要求一个*Point的接收者,我们可以简写 q.ScaleBy()
	q.ScaleBy(3)
	fmt.Println(q)
	// 实际上编译器会对变量进行%q的隐式转换.只有变量才允许这么做,包括结构体的字段.像p.X 和数组或 slice的元素,比如perim[0]
	// 不能对一个不能取地址的 Point接收者参数调用*Point 方法,因为无法获取临时变量的地址
	Point{1,2 }.ScaleBy(4) // 编译错误:不能获取 Point类型字面量的地址  cannot call pointer method on Point literal 无法用字面量类型调用指针方法
}
// 当定义一个类型允许 nil作为接收者的时候,应该在文档中显式地标明.
// IntList 整形链表 *IntList的类型 nil代表空列表
type IntList struct {
	Value int
	Tail *IntList
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}

	return list.Value + list.Tail.Sum()
}