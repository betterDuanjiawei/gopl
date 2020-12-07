/**
并发回声服务器
不仅可以同时处理多个来自客户端的请求
还可以通过go 来并发处理一个客户端的多次输入操作
需要考虑 net.Conn 并发调用的安全性
*/
package oldcode

import (
	"net"
	"log"
	"bufio"
	"time"
	"strings"
	"fmt"
)

func main() {
	listenner, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listenner.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn) // 同时处理多个客户端的请求
	}
}

func handleConn(c net.Conn)  {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1 * time.Second) // 同时处理一个客户端的多次输入
	}
	c.Close()
}

func echo(c net.Conn, shout string, delay time.Duration)  {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	time.Sleep(delay)
}
