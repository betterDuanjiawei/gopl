package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main()  {
	listener, err := net.Listen("tcp", "localhost:8000") // net.Listener 对象,它在一个网络端口撒花姑娘监听进来的连接.
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept() // 阻塞,直到有连接请求进来.然后返回net.Conn 对象来代表一个连接
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func echo(c net.Conn, shout string, delay time.Duration)  {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn)  {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}