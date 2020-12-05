package main

import (
	"fmt"
	"math"
	"time"
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
	q := Point{4 ,6}
	distanceFromP := p.Distance // 方法变量 后面没有()
	fmt.Println(distanceFromP(q))

	var origin Point
	fmt.Println(distanceFromP(origin))

	scaleP := p.ScaleBy
	scaleP(2)
	fmt.Println(p)
	scaleP(3)
	fmt.Println(p)
	scaleP(10)
	fmt.Println(p)

	r := new(Rocket)
	// time.AfterFunc()在指定的延迟之后调用一个函数值
	time.AfterFunc(10 * time.Second, func() {
		r.Launch()
	})
	// 用方法变量的简单写法
	time.AfterFunc(10 * time.Second, r.Launch)

}

type Rocket struct{}
func (r *Rocket) Launch(){}
