package main

import (
	"fmt"
	"time"
)

// 这里是 worker，我们将并发执行多个 worker。
// worker 将从 `jobs` 通道接收任务，并且通过 `results` 发送对应的结果。
// 我们将让每个任务间隔 1s 来模仿一个耗时的任务
func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("worker ", id, " processing job ", j)
		time.Sleep(time.Second)
		result <- j
	}
}

func main() {
	// 为了使用 worker 线程池并且收集他们的结果，我们需要 2 个通道
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 这里启动了 3 个 worker，初始是阻塞的，因为还没有传递任务
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	// 这里我们发送 9 个 `jobs`，然后 `close` 这些通道
	// 来表示这些就是所有的任务了
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	// 最后，我们收集所有这些任务的返回值
	for a := 1; a < 10; a++ {
		if a == 6 {
			close(results)
		}
		msg, ok := <-results
		fmt.Println(ok)
		fmt.Println("msg:", msg)
	}
	close(results)

	//ch := make(chan int)
	//close(ch)
	//fmt.Println(<-ch)
}

//在main函数里面向无缓冲的channel发送数据或者接受数据都会导致deadlock！- 解决方法是使用另外一个goroutine对这个channel收或者发数据
//在main函数里面向零值的而且有缓存的channel收数据会导致deadlock！- 解决方法是使用另外一个goroutine对这个channel发数据或者在main
//里面给这个channel先发数据
//为避免出现从channel中收数据时发生deadlock，最好使用select-case或者for i遍历去取channel中对应数量的数据
//从一个已经关闭的通道中读取数据，不论该通道是有缓冲的还是无缓冲的，都会读取到该通道类型的零值
//如果一个channel和另外一个channel关联了，关闭一个channel不会导致另外一个channel关闭
