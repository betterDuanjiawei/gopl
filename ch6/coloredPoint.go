package main

import (
	"fmt"
	"image/color"
	"math"
)

func main()  {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2
	fmt.Println(cp.Y)

	// 方法调用
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1},red}
	var q = ColoredPoint{Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point)) // 5
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point)) // 10

	demo2()
}


type Point struct {
	X, Y float64
}
func (p *Point)ScaleBy(factor float64)  {
	p.X *= factor
	p.Y *= factor
}

// Point 类型的方法
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y -p.Y)
}

// 内嵌使我们更加简单的定义了 ColoredPoint类型,它包含 Point类型的所有字段和更多的自由字段,如果需要可以直接使用ColoredPoint内的所有字段而不需要提Point
type ColoredPoint struct {
	Point // 提供 X,Y
	Color color.RGBA
}

type ColoredPoint2 struct {
	*Point
	Color color.RGBA
}

func demo2() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	p := ColoredPoint2{&Point{1, 1}, red}
	q := ColoredPoint2{&Point{5, 4}, blue}
	fmt.Println(p.Distance(*(q.Point)))
	fmt.Println(p.Distance(*q.Point)) // 和上面的一样,先 q.Point 再取* 表示 Point类型变量
	// 和这种是不一样的写法和含义 (*pptr).Distance() 前面是一个变量类型,*是用来转换的
	p.Point = q.Point // 都变成{5, 4}了
	p.ScaleBy(2)
	fmt.Println(*p.Point, *q.Point) // {10 8} {10 8}
}

type ColoredPoint3 struct {
	Point
	color.RGBA
}