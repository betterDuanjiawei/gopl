package main

import (
	"io"
	"log"
	"net"
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
		handleConn(conn)
	}
}

func handleConn(c net.Conn)  {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		// net.Conn 满足 io.Writer 接口.
		// time.Time.Format 方法提供了格式化日期和时间信息的方式.它的参数是一个模板.指示如何格式化一个参考时间,具体如 Mon Jan 2 03:04:05PM 2006 UTC-0700这样的形式
		if err != nil {
			return
		}
		time.Sleep(1* time.Second)
	}
}