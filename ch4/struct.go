package main

import (
	"fmt"
	"time"
)
// 结构体 零个或多个任意类型的命名变量组合在一起的聚合类型,
type Employee struct { // 定义结构体 成员变量的顺序对结构体同一性很重要,如果 Name 和 Address 相反,那么我们就在定义一个不同的结构体类型
	ID int
	Name, Address string
	DoB time.Time
	Position string
	Salary int
	ManagerID int
}
//  结构体 s不能定义定义一个有用相同结构体 s的成员变量,也就是一个聚合类型不可以包含它自己,对数组也适用,但是 s中可以定义一个 s的指针类型,即*s.
var dilbert Employee // 结构体变量
func main() {
	// 结构体的成员都可以通过点号方式来解答 dilbert.Name,可以给结构体的成员赋值

	// 结构体嵌套和匿名成员
	//var w Wheel
	//w.Circle.Center.X = 8
	//w.Circle.Center.Y = 8
	//w.Circle.Radius = 5
	//w.Spokes = 20

	//var w Wheel
	//w.X = 8 // w.Circel.Point.X = 8
	//w.Y = 8
	//w.Radius = 5
	//w.Spokes = 20

	var w Wheel
	w = Wheel{Circle{Point{8, 8}, 5}, 20}

	w2 := Wheel{
			Circle: Circle{
			Point{
					X: 8,
					Y: 8,
				},
				5, // ,必须的
		},
		Spokes: 20, // ,必须的
	}

	fmt.Printf("%v\n", w) // {{{8 8} 5} 20}
	fmt.Printf("%#v\n", w2) // main.Wheel{Circle:main.Circle{Point:main.Point{X:8, Y:8}, Radius:5}, Spokes:20}

}

//type Point struct {
//	X,Y int
//}
//type Circle struct {
//	Center Point
//	Radius int
//}
//type Wheel struct {
//	Circle Circle
//	Spokes int
//}

type Point struct {
	X,Y int
}
type Circle struct {
	Point
	Radius int
}
type Wheel struct {
	Circle
	Spokes int
}


