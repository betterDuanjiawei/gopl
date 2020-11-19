# 4.复合数据类型
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

// 新
```