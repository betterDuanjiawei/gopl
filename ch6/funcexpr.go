package main

import (
	"fmt"
	"math"
)

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
func main()  {
	p := Point{1, 2}
	q := Point{4, 6}
	distance := Point.Distance
	fmt.Println(distance(p, q))
	fmt.Printf("%T\n", distance)

	scale := (*Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)
	fmt.Printf("%T\n", scale)

	//5
	//func(main.Point, main.Point) float64
	//{2 4}
	//func(*main.Point, float64)

	path := Path{{1, 1}, {2, 2}, {3, 3}}
	path.TranslateBy(Point{1, 1}, true)
	fmt.Println(path) // [{2 2} {3 3} {4 4}]

}

func (p Point) Add(q Point) Point  {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.X}
}

type Path []Point // slice 引用传递,函数直接改变了该值

func (path Path) TranslateBy(offset Point, add bool)  {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}

	for i := range path {
		path[i] = op(path[i], offset)
	}
}