package main

import (
	"io"
	"net/http"
	"os"
	"path"
)

func main()  {

}

func fetch(url string) (filename string, n int64, err error)  {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body) // io.Copy 的错误更加重要点, 在写文件系统中,尤其是 NFS,写错误往往不是立即返回,而是推迟到文件关闭的时候.如果无法检查关闭操作的结果,就会导致一系列的数据丢失.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}

	return  local, n, err
}