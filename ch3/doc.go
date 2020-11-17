/**
go数据类型 4大类
1. 基本数据类型 数字(number) 字符串(string) 布尔值(boolean)
2. 聚合类型 数组(array) 结构体(struct) 通过组合各种简单类型得到更加复杂的数据类型
3. 引用类型 切片(slice) 指针(pointer) 字典(map) 函数(function) 通道(channel) 共同点:全部间接指向程序变量或状态,于是操作所引用的数据效果就是会遍及该数据的全部引用	
4. 接口类型


数值类型: 包括几种大小不同的整数, 浮点数 复数,各个类型的值有自己的大小,对正负号支持也各异
int8 int16 int32 int64
uint8 uint16 uint31 uint64
int 目前是使用最广泛的数值类型 32位或者64位
uint
rune 是 int32 的同义词,通常用于指明一个值是Unicode 码点
byte 是 uint8 的同义词,强调一个值是原始数据,而非量值
uintptr 大小不明确,但是足以完整存放指针, 仅用于底层编程,unsafe 包

int uint uintptr 都有别于其他大小明确的类型,int 和 int32是不同类型,即使 int大小天然就是32位,显示转换,才可以使用
int8 -128~127 -2(n-1) ~ 2(n-1) -1 
uint8 255 

go 二元操作符 运算 比较 逻辑, 按优先级 降序排序
* / % << (左移)>> (右移) &(位运算AND) &^(位清空 AND NOT)
+ - |(位运算 OR) ^(位运算 XOR 异或)
== != <= >= < >
&&
||

算术运算符可以用于整数 浮点数 复数,而取模运算符%只能用于整数.
go 取模余数的正负号总是于被除数一致, -5%3 = -2 -5%-3 = -2
除法运算/ 取决于操作数是否都为整型,整数相除,商会舍弃小数部分, 5.0/4.0 = 1.25 5/4 = 1

尽管go 具备无符号整形,而且某些值也不可能为负,但是我们往往还采用有符号的整数,如数组的长度

类型转换 T(x)将 x的值转换为类型 T,缩减大小的整形转换和浮点数和整数的转换会引起精度丢失
浮点转整数 会舍弃小数部分,趋零截尾
*/