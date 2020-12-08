package main
var (
	sema = make(chan struct{}, 1) // 用来保护 balance 的二进制信号量
	balance int
)
func Desposit(amount int)  {
	sema <- struct{}{} //获取令牌
	balance += amount
	<-sema // 释放令牌
}
func Balance() int {
	sema <- struct{}{} //获取令牌
	b := balance
	<-sema // 释放令牌
	return  b
}