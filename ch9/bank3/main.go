package main

import "sync"

func main()  {

}

var (
	mu  sync.Mutex
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance += amount
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}

// 不是原子操作, 3个串行操作,每个操作都申请释放了互斥锁,但是对于整个序列没有上锁
func Withdraw(amount int) bool {
	Deposit(-amount)
	if Balance() < 0 {
		Deposit(amount)
		return false
	}
	return true
}

func Withdraw2(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	Deposit(-amount) // 会导致函数里再次取获取锁,死锁,Withdraw2会一直卡住
	if Balance() < 0 {
		Deposit(amount)
		return false
	}
	return true
}

func Withdraw3(amount int)  bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}
	return true
}

func deposit(amount int)  {
	balance += amount
}

func Deposit(amount int) {
	mu.Lock()
	deposit(amount)
	mu.Unlock()
}

