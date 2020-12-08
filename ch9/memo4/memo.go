package memo

import "sync"

// 存了调用 Func 的结果
type Memo struct {
	f Func
	mu sync.Mutex // 在 cache前面,保护 cache
	cache map[string]*entry
}

type entry struct {
	res result
	ready chan struct{} // res 准备好后会关闭
}

// 用于记忆的函数类型
type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err error
}

func New(f Func) *Memo {
	return &Memo{f:f, cache:make(map[string]*entry)}
}
// 并发安全,但是每次调用f 时候都上锁,因此 get把我们希望并行的i/o操作串行化了.
func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	//defer
	e := memo.cache[key]
	if e == nil {
		// 对key第一次访问,这个 goroutine负责计算数据和广播数据
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)
		close(e.ready) // 广播数据已经准备完毕的消息,其他从该通道读取数据的操作将不会被阻塞,得到零值
	} else {
		// 因为e是指针数据,所以关闭读的操作可以放在上面
		memo.mu.Unlock()
		<-e.ready // 等待数据准备完毕
	}
	return e.res.value, e.res.err
}


