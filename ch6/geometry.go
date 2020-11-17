package main

import (
	"math"
	"fmt"
)

type Point struct{
	X, Y float64
}



// 普通函数
func Distance(p, q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

// Point 类型的方法
func (p Point)Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func main(){
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}

	fmt.Println(perim.Distance())
}

type Path []Point // 放在下面和上面没啥区别
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		// if (i > 0) { 有()也不报错
		if i > 0 { 
			sum += path[i-1].Distance(path[i])
		}
	}

	return sum
}