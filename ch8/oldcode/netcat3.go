package oldcode

import (
	"net"
	"log"
	"io"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func ()  {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{} // 无任何信息,单纯的同步
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // 会阻塞等待后台 goroutine 完成
	log.Println("main goroutine done")
}

func mustCopy(dst io.Writer, src io.Reader)  {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}