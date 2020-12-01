# 4.复合数据类型
* 数组和结构体都是聚合类型,它们的值由内存中的一组变量组成.数组的元素具有相同的类型,而结构体种的元素类型则可以不同.
* 数组和结构体都是长度固定的.反之 slice 和 map 是动态数据结构,它们的长度在元素添加到结构中时可以动态增长
## 4.1 数组
* 数组是具有固定长度且拥有零个或者多个相同数据类型元素的序列.由于数组长度固定,在 go中很少使用. slice 的长度可以增加和缩短,很多场合使用的更多
* 数组的每个元素是通过索引来访问的.索引从0~len(s) - 1
* go 内置的 len()可以返回数组中的元素个数
* 数组可以使用range 循环
* 一个新数组中的元素初始值为元素类型的零值,对于数字来说就是0,也可以使用数组字面量来初始化一个数组: `var q [3]int = [3]int{1, 2, 3}`
* 在数组字面量中,如果省略号...,出现在数组长度的位置,那么数组的长度由初始化数组的元素个数决定. `q := [...]int{1, 2, 4, 5, 7, 8}`
* 数组的长度是数组类型的一部分.所以[3]int和[4]int 是不同的数组类型; 数组的长度必须是常量表达式,也就是说,这个表达式的值在程序编译时就可以确定
* 数组 slice map 结构体的字面语法都是相似的.
* 没有指定值的索引位置的元素默认被赋予数组元素类型的零值.如`r := [...]int{99: -1}` 定义了一个拥有100个元素的数组 r,除了最后一个元素值时-1外,该数组中的其他元素值都是0
* 如果一个数组的元素类型是可比较的.那么这个数组也是可以比较的.用==操作符来比较两个数组.
* crypto/sha256 Sum256() 使用 SHA256加密散列算法生成一个摘要, 摘要信息是256位,即[32]byte
* go的数组是值传递
* 可以显式的传递一个数组的指针给函数,这样在函数内部的对数组的任何修改都回反映到原始数组上面

## 4.2 slice
* slice表示一个拥有相同类型元素的可变长度序列.[]T 元素类型是 T,它是没有长度的数组类型
* 一个底层数组可以对应多个 slice
* x[m:n] 求字符串的子串操作和对字节 slice([]byte) 做 slice 操作有相似性
* 因为 slice包含了指向数组元素的指针,所以将一个 slice 传递给函数的时候,可以在函数内部修改底层数组的元素.,创建一个数组的 slice 等于为数组创建了一个别名.
* slice []int{1, 2, 3}和数组的字面量 [3]int{1, 2, 3}很像
* 和数组不同的是,slice 无法做比较,因此不能用==来测试两个 slice是否拥有相同的元素.
* bytes.Equal() 来比较两个字节 slice []byte.
* slice 的元素不是直接的. 如果底层数组元素改变.同一个 slice在不同的时间会拥有不同的元素.
* 由于散列表 go map 类型,仅对元素的键做浅拷贝.这就要求散列表李的键在整个生命周期内必须保持不变. 所以 slice 不能作为 map 的键
* 不允许直接比较 slice. slice 唯一允许比较的就是和 nil 做比较
* slice的零值是 nil,值为 nil的slice没有对应的底层数组.值为 nil 的 slice长度和容量都是0,但是也有非nil 的slice长度和容量是零.[]int{} make([]int, 3)[3:]
* 对于任何类型它的值可以是 nil,那么这个类型的 nil 值可以用一种转换表达式:[]int(nil)
* 检查一个 slice是否为空,用 len(s) == 0, 而不是 s == nil,因为 s != nil的情况下,slice也可能为空
```
var s []int
s = nil // len(s) == 0, s == nil
s = []int(nil) // len(s) == 0, s == nil
s = []int{} // len(s) == 0, s != nil
```
* make([]T, len, cap)  make创建了一个无名数组,并返回了它的一个slice.
* append() 内置函数 append()用来将元素追加到 slice后面 因为不清楚 append()到底有没有重新分配内存. 通常我们将append的调用结果再次赋值给传入的append 函数的 slice `runes = append(runes, i)`
* copy(dst, src) 第一个参数是目标 slice,第二个参数是源 slice.copy 将源 slice的元素复制到目标 slice中. 返回值:实际复制的个数
* 不仅是调用 append 函数的情况下需要更新 slice变量.对应任何函数,只要有可能改变slice的长度或者容量,抑或是使得 slice指向不同的底层数组.都需要更新slice变量.
* 为了正确使用 slice.虽然底层数组的元素是间接引用的.但是 slice的指针 长度和容量不是.要更新一个 slice的指针 长度或容量必须使用如上的显式赋值.
* 从这个角度看.slice并不是纯引用类型.而是下面这种聚合类型
```
type IntSlice struct{
    ptr *int
    len, cap int
}
```
* append() 可以给 slice 同时添加多个元素,甚至添加另一个 slice里的所有元素
```
	var z []int
	z = append(z, 1)
	z = append(z, 2, 3)
	z = append(z, 4, 5, 6)
	z = append(z, z...) // 追加 z中的所有元素 参数后的...表示如何将一个slice转换为参数列表
	fmt.Println(z) // [1 2 3 4 5 6 1 2 3 4 5 6]
```
* appendInt(x []int, y ...int) []int //y ...int 表示该函数接受可变长度参数列表

* slice可以用来实现栈,
```
stack = append(stack, v)
top := stack[len(stack)-1] // 栈顶是最后一个元素
stack = stack[:len(stack)-1] // 弹出最后一个元素来缩减栈
```
* 从 slice的中间移除一个元素,并保留剩余元素的顺序 copy来将高位索引的元素向前移动来覆盖被移除元素所在位置
```
func move(slice []int, i int) []int {
    copy(slice[i:], slice[i+1:])
    return slice[:len(slice)-1]
}
```
* 如果不需要维持 slice的剩余元素的顺序,可以简单的将slice的最后一个元素替换要去除的元素
```
func move(slice []int, i int) []int {
    slice[i] = slice[len(slice) - 1]
    return slice[:len(slice)-1]
}
```
 * r, size := utf8.DecodeLastRune(b[:i]) // 取最后一个 unicode码点

## 4.3 map
* 散列表:拥有键值对元素的无序集合,键是唯一的,键对应的值可以根据键来添加 更新 移除. 常量时间操作
* map是散列表的引用,map 类型 map[k]v,k 和 v是字典的键值对应的数据类型
* map 的键必须拥有相同的数据类型, 值也同样拥有相同的数据类型,但是键的类型和值的类型不一定相同
* 键的类型 k,必须是可以通过==比较的数据类型,所以 map可以检测一个键是否存在; 最好不要比较浮点类型的键,NaN可以是浮点型值
* 新的空的 map: make(map[string]int) | map[string]int{}
* map的元素访问:下标的方式 ages["djw"]
* delete(map, k) 移除 map中的元素,即使键不存在于 map中也是安全的
* 查找给定键对应的值,如果不存在返回值类型的零值
* x+=y x++ 对 map 同样适用
* map的元素是没有地址的 用&map[k] 方式获取是错误的,原因是 map 是可以增长的,会导致已有的元素重新散列到新的存储位置,是之前获取到的地址无效
* 循环 for range 关键字
* map中的元素迭代顺序是不固定的
*  var map1 map[string]int  只声明了没有初始化,是零值 map的零值是 nil,也就是说没有引用任何散列表
* 空 map 的长度 len()是0,map由长度,但是没有容量 cap() invalid argument map1 (type map[string]string) for cap
* 大多数map的操作都可以安全的在map 的零值 nil 上操作,包括查找元素 删除元素 获取 map的个数,执行 range循环,这和空 map 的行为一致,但是向零值map中设置元素会报错;设置之前必须初始化 map
```
var map1 map[int]int //只声明了,没有初始化,由零值是 nil
map = map[int][int]{} // 初始化的一个空的 map map[] 不是 nil
```
* 元素的值是数值类型,去判断元素在不在 map 时候,因为有可能恰好该元素就是0,或者没有该元素,返回值的零值0;所以需要这样判断
```
// 第二个值是 bool,用来报告该元素是否存在, 一般叫做 ok
if age, ok := ages["djw"]; !ok {
    // djw不在 map的键中
}
```
* map和 slice一样不可比较,唯一合法的比较就是和 nil作比较, 为了判断两个 map 是否拥有相同的键和值,必须写一个循环

```
ages := make(map[string]int)
// map 字面量新建一个带有初始化键值对元素的 map
ages2 := map[string]int{
    "djw": 29,
    "lxq": 28,
}

ages3 := make(map[string]int)
ages3["djw"] = 29
ages3["lxq"] = 28

```

* 如果需要按照某种顺序来遍历 map中的元素,我们必须显式来给键排序,如果 map 的键是字符串,可以用 sort包中的 Strings 函数来进行键的排序.
```
//var s []string
	s := make([]string, 0, len(m)) 
	for k := range m {
		s = append(s, k)
	}
	sort.Strings(s)
	for _, v := range s {
		fmt.Println(m[v])
	}
```
指定一个 slice 的长度会更加高效
* map的零值类型是nil,长度是0, 也就是说没有引用任何类型的散列表
```
	var test1 map[string]int
	fmt.Println(test1==nil) // true
	fmt.Println(len(test1)==0) // true
```
大多数 map 操作都可以安全地在 map的零值nilh 上执行,包括查找元素,删除元素,获取 map元素个数(len),执行 range 循环,因为这个和空 map的行为一致,但是向零值 map 中设置元素会导致错误:
```
	test1["age"] = 1
	fmt.Println(test1) // panic: assignment to entry in nil map
```
* 设置元素之前,必须初始化 map
* 通过下标的方式访问 map 中的元素总是会有值.如果键在 map中,你会得到键对应的值;如果键不在 map中,你将得到 map值类型的零值. 如果你想知道一个元素是否在 map 中,如果元素类型是数值类型,你需要能辨别一个不存在的元素或者恰好这个元素的值是0,
可以:
```

if age, ok := ages["bob"]; !ok {
// ok 是布尔值,用来报告该元素是否存在.
}
```
* 和 slice 一样,map不可比较,唯一合法的比较就是和 nil 比较.为了判断两个 map 是否拥有相同的键和值,必须写一个循环
```
// 值为 int类型的 map的比较
func mEqual(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
    // 注意我们使用!ok 来区分"元素不存在"和"元素存在但值是零"的情况,如果简单写成 xv != y[k],那么 mEqual(map[string]int{"A":0}, map[string]int{"B":42})
	for k, xv := range x {
		if yv,ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	
	return true
}
}
```
* map 的键是唯一的,所以可以用来实现集合
* 使用 map 记录 add 函数被调用的次数, 使用 fmt.Sprintf("%q")来将一个一个字符串 slice转换为一个适合做 map 键的字符串. 同样的方法适用于任何不可直接比较的键类型,不仅仅局限于 slice.
```
var m = make[string]int

func k(list []string) string {
    return fmt.Sprintf("%q", list)
}

func add(list []string){
    m[k(list)]++
}
func count(list []string){
    return m[k(list)]
}
```
* xx.ReadRune()
```
in := bufio.NewReader(os.Stdin)
r, n, err := in.ReadRune() // 返回 rune(解码的字符串 unicode.ReplacementChar 不合法的 utf-8字符,长度是1), nbytes(字节长度), error(错误  io.EOF 文件结束, 其他错误)
```
* map 的值类型可以是复合类型,如 map和 slice, 值:map[string]bool, addEdge: 延迟初始化 map 的方法,当 map中的每个键第一次出现的时候初始化 hasEdge:在 map中值不存在的情况下,也可以直接使用
```
var graph = make(map[string]map[string]bool)
func main()  {
	addEdge("a", "first")
	fmt.Println(hasEdge("a", "first"))
	fmt.Println(hasEdge("a", "seconds"))
}

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
```

## 4.4结构体
* 结构体是将零个或者多个任意类型的命名变量组合在一起的聚合数据类型.每个变量都叫做结构体的成员.
* 结构体的成员通过.号来访问.结构体的成员都是变量,可以给结构体的成员赋值,或者获取成员变量的地址,通过指针来访问它.
* .号同样可以用在结构体指针上.可以用.号来访问结构体指针
* 结构体的成员变量通常一行写一个,变量名称在类型前面,但是相同类型的连续成员变量可以写在一行上.
```
type Employee struct {
    ID int,
    Name, Address string,
    DoB time.Time
    Position string
    salary int
    ManagerID int
}
```
* 成员变量的顺序对于结构体的同一性很重要,如果是互换了 Name和 Address 的顺序,那么我们就是在定义不同类型的结构体.一般来将我们只组合相关的成员变量
* 一个结构体可以同时包含可导出和不可导出的成员变量
* 命名结构体类型不可以定义一个用于相同结构体类型 s的成员变量,也就是一个聚合类型不可以包含它自己(对数组也适用),但是 S中可以包含一个 S 的指针类型,即*S,我们可以用它来创建一些递归数据结构:链表和树.
```
二叉树
package main

import "fmt"

type tree struct {
	value int
	left, right *tree
}
func main()  {
	values := []int{1, 3, 2, 5, -1, 18, 100, -8}
	Sort(values)
	//fmt.Println(Sort(values))
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	fmt.Println(appendValues(values[:0], root))
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t // return &tree{value: value}
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}

	return t
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}
```
* 结构体的零值由结构体成员的零值组成.
* 没有任何成员变量的结构体称为空结构体,写作:struct{},它没有任何长度,也不携带任何信息.有时候用它来替代被当做集合使用的 map中 的布尔值,来强调只有键是有用的.但是这种写法节约的内存很少,而且语法复杂,所以尽量避免这样使用
```
seen := make(map[string]struct{})

if _, ok := seen[s]; !ok {
    seen[s] = struct{}{}
}
```
* 结构体字面量:通过设置结构体的成员变量来设置,两种方式:不指定成员名称按顺序赋值,和指定成员名字按需赋值
```
type Point struct{x,y,z int}
p := Point{1, 2, 3} //正确的顺序,每个都赋值,不推荐使用,开发和阅读空难,可维护性差
p2 := Point{x: 1, z:3} // 指定部分或全部成员变量的名称和值来初始化结构体变量.因为指定了成员变量的名字,所以他们的顺序无所谓的
```

