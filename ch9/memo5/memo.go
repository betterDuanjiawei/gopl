package memo

// 存了调用 Func 的结果
type Memo struct {
	requests chan request
}
type request struct {
	key string
	response chan<- result
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
	memo := &Memo{requests:make(chan request)}
	go memo.server(f)
	return memo
}
// 并发安全,但是每次调用f 时候都上锁,因此 get把我们希望并行的i/o操作串行化了.
func (memo *Memo) Get(key string) (value interface{}, err error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) server(f Func)  {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
}

func (memo *Memo) Close()  {
	close(memo.requests)
}

func (e *entry) call(f Func, key string)  {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result)  {
	<-e.ready
	response <- e.res
}




