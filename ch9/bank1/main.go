package bank

var deposits = make(chan int) // 发送存款额
var balances = make(chan int) // 接收余额

func Deposit(amount int) {
	deposits <- amount
}

func balance() int {
	return <-balances
}

func teller() {
	var balance int // balance 被限制在 teller goroutine 中
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:

		}
	}
}

func init() {
	go teller() // 启动监控 goroutine
}
