package oldcode

import (
	"log"
)

func main() {
	
}

// 处理缩略图的操作是串行的, 其实是各个操作是无关联的独立的,可以用并行处理来提高效率
func makThumbnails(filenames []string)  {
	for _, f := range filenames {
		if _, err := imageFile(f); err != nil {
			log.Println(err)
		}
	}
}

// 每一个文件的操作都是并行的,即使和第一个版传递的是只有一个元素的 slice, 2版也比1版快,是因为 go imageFile(f) 之后函数就直接返回主 goroutine 了,
func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go imageFile(f)
	}
}
// chan 是关键词,不能随便使用, 通过通道来知道是否 go 执行完成
func makeThumbnails3(filenames []string)  {
	ch := make(chan struct{})
	for _, f := range  filenames {
		go func (f string)  {
			imageFile(f)
			ch <- struct{}{}
		}(f) // 这里的匿名函数 显示传递 f,可以保证 go语句执行的时候,使用的 f 是正确的
		// go func(){
		// 	imageFile(f) // 错误,f会被更新,都是最后一个元素
		// }()
	}

	for range filenames {
		<-ch
	}
	
}

// error 类型关键词 goroutine泄露可能导致 整个程序卡住或者系统内存耗尽
func makeThumbnail4(filenames []string)  {
	errors := make(chan error) // 无缓冲通道
	for _, f := range filenames {
		// errors := make(chan error)
		go func (f string) error  {
			_, err := imageFile(f)
				errors <- err
		}(f)
	}
	for range filenames {
		if err := <-errors; err != nil {
			return err // 这里不正确,goroutine 泄露,后面的 error 没有接收者了
		}
	}

	return nil
}

// 使用的是有缓冲通道,不会发生 goroutine 泄露,但是只有有一个错误还是会返回
func makeThumnail5(filenames []string) (thumbfiles []string, err error)  {
	type item struct {
		thumbfile string
		err error
	}

	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func (f string)  {
			var it item
			it.thumbfile, it.err = imageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}
	return thumbfiles, nil
}

// 安全的计数器类型 sync.WaitGroup
func makeThumbnail6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for f := range filenames {
		wg.Add(1)
		// worker
		go func (f string)  {
			defer wg.Done() // 等价于 Add(-1) //保证发送错误时候计数器可以递减
			thumb, err := imageFile(f)
			if err != nil {
				log.Println(err)
				return // 结束 go 的匿名函数
			}
			info, _ := os.Stat(thumb) // 获取文件的属性,返回描述文件f 的 FileInfo 类型值
			size <- info.Size() // 获取文件大小
		}(f)
	}
	// closer
	go func ()  {
		wg.Wait() // 等待在关闭之前
		close(sizes)
	}()
	var total int64
	for size := range sizes {
		total += size
	}

	return total
}

func imageFile(infile string) (string, error)  {
	
}