package memo

import "sync"

// 存了调用 Func 的结果
type Memo struct {
	f Func
	mu sync.Mutex // 在 cache前面,保护 cache
	cache map[string]result
}

// 用于记忆的函数类型
type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err error
}

func New(f Func) *Memo {
	return &Memo{f:f, cache:make(map[string]result)}
}
// 并发安全,但是每次调用f 时候都上锁,因此 get把我们希望并行的i/o操作串行化了.
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	//defer
	res, ok := memo.cache[key]
	memo.mu.Unlock()
	if !ok {
		res.value, res.err = memo.f(key) // 在两个临界区之间,可能会有多个 goroutine 计算f(key)并且更新map
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	return res.value, res.err
}