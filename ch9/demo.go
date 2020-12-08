package main

import (
	"go4.org/testing/functest"
	"image"
)

func main()  {

}

func demo1()  {
	var x []int
	go func() {
		x = make([]int, 10)
	}()

	go func() {
		x = make([]int, 1000000)
	}()
	// 雷区:未定义行为
	// x可能是 nil,一个长度为10的 slice 或一个长度为100000的 slice. slice:指针 长度 容量, 指针时第一个 make 调用而长度来自第二个 make 调用,x是嵌合体,
	// 它名义上长度为100000,但是底层的数组只有10个元素,在这种情况下,尝试存储到第99999个元素,会伤及一段很遥远的内存,其恶果无法预测看,问题也很难调试和定位.雷区:未定义行为
	x[99999] = 1 // 未定义行为,可能造成内存异常
}

func demo2()  {


}
// 下面有延迟初始化,对于每一个键,在第一次访问时才触发加载. 并行调用,Icon 的结果
var icons = make(map[string]image.Image)
func Icon(name string) image.Image  {
	icon, ok := icons[name]
	if !ok {
		icon = loadIcon(name)
		icons[name] = icon
	}
	return icon
}

func loadIcon(name string) image.Image  {

}
// 如果在创建其他 goroutine之前就用完整的数据来初始化 map,并且不再修改.那么无论多少 goroutine也可以安全地并发调用 Icon,因为每个 goroutine都只读取这个 map
var icons = map[string]image.Image{
	"spades.png" : loadIcon("spades.png"),
	"hearts.png" : loadIcon("hearts.png"),
	"diamonds.png" : loadIcon("diamonds.png"),
	"clubs.png" : loadIcon("clubs.png"),
}

func Icon(name string) image.Image {
	return icons[name]
}

type Cake struct {
	state string
}

func baker(cooked chan<- *Cake)  {
	for {
		cake := new(Cake)
		cake.state = "cooked"
		cooked <- cake // baker 不再访问cake变量
	}
}

func icer(iced chan<- *Cake, cooked <-chan *Cake)  {
	for cake := range cooked {
		cake.state = "iced"
		iced <- cake // icer 不再访问 cake变量
	}
}

var icons map[string]image.Image

func loadIcons()  {
	icons = map[string]image.Image{
		"spades.png" : Icon("spades.png"),
		"hearts.png" : Icon("hearts.png"),
		"diamonds.png" : Icon("diamonds.png"),
		"clubs.png" : Icon("clubs.png"),
	}
}
// 并发不安全
func Icon(name string) image.Image {
	if icons[name] == nil {
		loadIcons() // 一次性初始化
	}
	return icons[name]
}

