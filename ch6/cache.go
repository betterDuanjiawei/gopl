package main

import "sync"

func main()  {

}

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