# 6 方法
* oop 面向对象编程
* 对象就是简单的一个值或变量,并且拥有方法,而方法是某种特定类型的函数.面向对象编程就是使用方法来描述每一个数据结构的属性和操作.于是,使用者不需要了解对象本身对的实现

## 6.1 方法声明
* 方法的声明和普通函数的声明类似,只是在函数名字前面多加个一个参数.这个参数把这个方法绑定到这个参数对应的类型上
```
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
```
* 附加的参数 p称为方法的接收者.
* go语言中,接收者不再使用特殊名(this self),而是我们自己选择接受者的名字,就想其他的参数变量一样,由于接收者会频繁的使用到,因此最好能够选择简短而且在整个方法中名称始终保持一致的名字,最常用的方法就是取类型名称的首字母,如Point 的 p
* 调用方法的时候,接收者在方法名前面.这样就和声明保持一致
```
    p := Point{1, 2}
	q := Point{4, 6}
	// 两个没有冲突
	fmt.Println(Distance(p, q)) // 包级别的函数 geometry1.Distance()
	fmt.Println(p.Distance(q)) // Point 类型的方法 Point.Distance()
```
* 表达式p.Distance()称作选择子.因为它为 接收者 p 选择了合适的方法.选择子也用于选择结构类型中的某些字段值,就像 p.X中的字段值.方法和字段值来自同一命名空间,所以 X 名的方法会和 X的字段值冲突,编译器会报错
* 因为每一个类型都有他自己的一个命名空间,所以我们可以在不同的类型中使用相同的名字作为方法名
* Path 是一个命名的 slice 类型,而非 Point这样的结构体类型.但是我们依旧可以给它定义方法.go 和许多其他面向对象的语言不同.它可以将方法绑定到任意类型上.可以很方便的为简单的类型(int strings slice map func())定义附加行为
* 同一个包下的任何类型都可以声明方法.只要它的类型既不是指针类型也不是接口类型
```
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
```
* 编译器会根据方法名和接收者来决定调用哪一个方法
* 类型所拥有的方法名是唯一的.但是不同类型可以拥有相同的方法名,而且没必要用附加字段来修饰方法名(PathDistance)
* 命名比函数更加简短.在包的外部进行调用的时候,方法能够使用更加简短的名字而且可以忽略包名
```
	perim := geometry1.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())
	fmt.Println(geometry1.PathDistance(perim))
```

## 6.2 指针接收者的方法
* 由于主调函数会复制每一个实参变量,如果函数需要更新一个变量,或者如果一个实参太大而我们希望避免复制整个实参,因此我们必须使用指针来传递变量的地址.这也同样适用于更新接收者,我们将它绑定到接收者 *Point
```
// 这个方法的名字是: (*Point).ScaleBy 圆括号是必须的.如果没有圆括号,则表达式会被解析为 *(Point.ScaleBy)
func (p *Point) ScaleBy(factor float64){
    p.X *= factor
    p.Y *= factor
}
```
* 在真实的程序中,习惯上遵循如果 Point的任何一个方法使用指针接收者,那么所有的 Point 方法都应该使用指针接收者,即使有的方法并不需要.
* 命名类型Point和指向他们的指针*Point是唯一可以出现在接收者声明处的类型.而且为了防止混淆,不允许本身是指针的类型进行方法声明
```
type P *int

//  invalid receiver type P (P is a pointer type)
func (p P) test()  {
	fmt.Println("test")
}
```
* 如果接收者 q是 Point类型的变量,但是方法要求一个*Point的接收者,我们可以简写 q.ScaleBy()
```
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
```
* 但是如果实参接收者是*Point类型,那么用 Point.Distance的方式调用 Point类型的方法是合法的.因为我们有办法从地址获取 Point的值,只要解引用指向接收者的指针值即可.编译器会自动隐式的插入一个*操作符
```
pptr.Distance()
(*pptr).Distance()
```
* 可以成立的三种形式:
```
// 1.实参接受者和形参接收者是同一个类型.比如都是 T类型或*T类型
Point{1, 2}.Distance(4)
pptr.ScaleBy(2)
// 2.实参接收者是 T 类型的变量,而形参接收者是*T类型.编译器会隐式的获取变量的地址
p.ScaleBy(2) //隐式的转换为 (&p).ScaleBy(2)
// 3.实参接收者是*T类型而形参接受者是 T 类型,编译器会隐式地解引用接收者.获得实际的取值
pptr.Distance(2) // (*pptr).Distance()
```
### nil 是一个合法的接收者
* 就像一些函数允许nil 指针做为实参,方法的接收者也是如此,尤其是 nil是类型中有意义的零值,map slice 类型.
```
/**
package url
type Values map[string][]string

func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}
 */
// make slice字面量 m[key]
func main()  {
	m := url.Values{"lang" : {"en"}} // 直接构造
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("q"))
	fmt.Println(m.Get("item"))
	fmt.Println(m["item"])

	m = nil
	fmt.Println(m.Get("item")) // Values(nil).Get["item"]
	m.Add("item", "3") // panic: assignment to entry in nil map 宕机 更新一个空 map
}
```

## 6.3 通过结构体内嵌组成类型
* 内嵌使我们更加简单的定义了 ColoredPoint类型,它包含 Point类型的所有字段和更多的自由字段,如果需要可以直接使用ColoredPoint内的所有字段而不需要提Point
* 同理这也适用于 Point类型的方法.我们能够通过类型为 ColoredPointd的接收者调用内嵌类型Point的方法.即使在 ColoredPoint类型没有声明过这个方法的前提下
* Point 类型的方法被嵌入到ColoredPoint类型中.以这种方式.内嵌允许构成复杂的类型,该类型由许多个字段构成,每个字段提供了一些方法
```
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
```
* 不是基于类的面向对象的父类和子类的关系.ColoredPoint 类型并不是 Point类型,但是它包含了一个 Point,并且它的另外2个方法 Distance()和 ScaleBy()来自 Point
```
p.Distance(q) // 编译错误,不能将 q (ColoredPoint)类型转换为Point 类型.
// 内嵌的字段会告诉编译器生成额外的包装方法来调用Point 声明的方法.
func (p ColoredPoint) Distance(q Point) int {
    return p.Point.Distance(q) // 接收者的值是 p.Point 而不是 p
}

func (p *ColoredPoint) ScaleBy(factor float64){
    p.Point.ScaleBy(factor)
}
```
* 匿名字段类型可以是个指向命名类型的指针,这个时候字段和方法间接的来着所指向的对象
```
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
```
* 结构体类型可以拥有多个匿名字段.同时拥有Point的所有方法和RGBA的方法,以及其他任何在 ColoredPoint3中直接声明的方法 先从自己声明的开始寻找,一层层寻找.当同一查找级别中拥有同名方法的时候,编译器会报告选择子不明确的错误
```
type ColoredPoint3 struct {
	Point
	color.RGBA
}
```
* 方法只能在命名的类型(Point)和指向他们的指针(*Point)中声明,但内嵌帮助我们能够在未命名的结构体类型中声明方法
```
var (
	mu sync.Mutex
	mapping = make(map[string]string)
)

func Lookup(key string) string {
	mu.Lock()
	v := mapping[key]
	mu.Unlock()
	return v
}
// 新的变量名更加贴切,而且 sync.Mutex 是内嵌的,它的 Lock和 Unlock方法也包含进结构体了,允许我们直接使用cache变量进行加锁
var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping:make(map[string]string), //如果要换行,那么这里必须要有,
}

func LookUp2(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}
```

## 6.4 方法变量与表达式
* 方法变量使得函数只需要提供实参而不需要提供接收者就能够调用 p.Distance 赋予一个变量 这个变量就是方法变量
```
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
```
* 与方法变量相关的是方法表达式.和调用一个普通的函数不同的是,在调用方法的时候,必须提供接收者,并且按照选择子的语法进行调用.而方法表达式写成 T.f或(*T).f,其中 T是类型,是一种函数变量,把原来的方法接收者替换成函数的第一个形参,因此它可以像平常函数一样调用
```
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

}
```
* 如果你需要用一个值来代表多个方法中的一个,而方法都属于同一个类型,方法变量可以帮助你调用这个值所对应的方法来处理不同的接收者
```

    path := Path{{1, 1}, {2, 2}, {3, 3}}
	path.TranslateBy(Point{1, 1}, true)
	fmt.Println(path) // [{2 2} {3 3} {4 4}]

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
```

## 6.5 位向量 没看懂,需要后续继续重点学习


## 6.6 封装
* 如果变量或方法是不能通过对象访问到的,这称为封装的变量或方法.封装(数据隐藏)是面向对象编程中重要的一方面
* go 只有一种方式控制命名的可见性:定义的时候,首字母大写的标识符是可以从包导出,而首字母没大写的则导不出.同样的机制也适用于结构体内的字段和类型的方法,结论就是: 要封装一个对象,必须使用结构体
* go中封装的是单元是包而不是类型.无论是在函数体内的代码 还是方法内的代码,结构体类型内的字段对于同一包内的所有代码都是可以见的
* 封装的优点:
    1. 使用方不能直接修改对象的变量,所以不需要更多的语句来检查变量的值
    2. 隐藏实现细节可以防止使用方依赖的属性发生改变
```
type Buffer struct{
    buf []byte
    initial [64]byte
}
func (b *Buffer) Grow(n int) {
    if b.buf = nil {
        b.buf = b.initial[:0]
    }
    if len(b.buf) + n > cap(b.buf) {
        buf := make([]byte, b.Len(), 2*cap(b.buf) + n)
        copy(buf, b.buf) // 复制数据到 buf里
        b.buf = buf // 替换 b.buf 的容量
    }
}
```
    3. 防止使用者肆意地改变对象内部的变量,因为对象的变量只能被同一个包内的函数修改,所以包的作者能够保证所有的函数可以维护对象内部的资源
```
Counter 类型允许使用者递增或者重置计数器,但是不能随意设置修改当前计数器的值
type Counter struct {n int}
func (c *Counter) N() int {return c.n}
func (c *Counter) Increment() {c.n++}
func (c *Counter) Reset() {c.n = 0}
```
* 仅仅用来获得或者修改内部变量的函数成为getter和 setter.一般在命名 getter的时候,通常将 get前缀省略,这个简洁的命名习惯也适用于其他冗余的前缀上
```
package Log
type Logger struct {
    flags int
    prefix string
}
func (l *Logger) Flags() int
func (l *Logger) SetFlags(flags int)
func (l *Logger) Prefix() string
func (l *Logger) SetPrefix(prefix string)
```

* go 也运行导出的字段.但是一旦导出就要考虑各种问题了
* 封装并不总是必须的.time.Duration对外暴露了int64的整形用于获得微妙数,这使得我们能够对其进行通常的数学运算和比较操作,甚至定义常数
```
const day = 24 * time.Hour
fmt.Println(day.Seconds())
```

