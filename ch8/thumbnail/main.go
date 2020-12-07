package main

import (
	"gopl.io/ch8/thumbnail"
	"log"
	"os"
	"sync"
)

func main()  {

}

func makeThumbnails(filenames []string)  {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

func makeThumbnails2(filenames []string)  {
	for _, f := range filenames {
		go thumbnail.ImageFile(f) // 返回的特快,没有等执行完毕,就返回了
	}
}

func makeThumbnails3(filenames []string)  {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) { // 显式参数,确保当 go语句执行的时候,使用 f的当前值
			thumbnail.ImageFile(f)
			ch <- struct{}{}
		}(f) // 要加 f
	}
	for range filenames {
		<-ch
	}
}

func makeThumbnails4(filenames []string) error {
	errors := make(chan error)
	for _, f := range filenames {
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			errors <- err
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			return err // 不正确,goroutine 泄露, 整个程序卡住或系统内存耗尽
		}
	}
	return nil
}

func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err error
	}

	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}
	for range filenames {
		it := <- ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	return thumbfiles, nil
}

func makeThumbnails6(filenames <-chan string) int64  {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for f := range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			thumb, err  := thumbnail.ImageFile(f)
			if err != nil {
				log.Print(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}

	go func() {
		wg.Wait()
		close(sizes)
	}()
	var total int64
	for size := range sizes { // 主 goroutine把大部分时间花在range循环休眠上了,等待工作者发送值或等待closer()关闭通道
		total += size
	}
	return total
}

