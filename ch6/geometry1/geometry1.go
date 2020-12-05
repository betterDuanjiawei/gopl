package geometry1

import (
	"fmt"
	"math"
)

func main()  {
	p := Point{1, 2}
	q := Point{4, 6}
	// 两个没有冲突
	fmt.Println(Distance(p, q)) // 包级别的函数 geometry1.Distance()
	fmt.Println(p.Distance(q)) // Point 类型的方法 Point.Distance()

	// perim := []Point{} // 错误写法
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim. Distance())
}

type Point struct {
	X, Y float64
}

// 普通函数
func Distance(p, q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}
// Point 类型的方法
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y -p.Y)
}

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func PathDistance(path Path) float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}