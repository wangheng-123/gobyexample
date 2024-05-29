//package waitgroups
//
//import (
//	"fmt"
//	"sync"
//)

//func f1(i *int, done1, done2 chan bool) {
//	<-done1
//	fmt.Println("f1:", *i)
//	*i++
//	time.Sleep(time.Second)
//	done2 <- true
//}
//
//func f2(j *int, done1, done2 chan bool) {
//	<-done2
//	tmp := *j
//	str := byte(tmp + 'a' - 1)
//	fmt.Println("f2:", string(str))
//	*j++
//	time.Sleep(time.Second)
//	done1 <- true
//}
//
//func waitgroups() {
//	done1 := make(chan bool, 1)
//	done2 := make(chan bool, 1)
//	i := 1
//	j := 1
//	done1 <- true
//	for {
//		go f1(&i, done1, done2)
//		go f2(&j, done1, done2)
//		//time.Sleep(2 * time.Second)
//	}
//
//	time.Sleep(60 * time.Second)
//}

//package waitgroups
//
//import (
//	"fmt"
//	"time"
//)
//
//func waitgroups() {
//	numCh := make(chan int)
//	letterCh := make(chan rune)
//
//	go func() {
//		defer close(numCh)
//		defer close(letterCh)
//		for i := 1; i <= 26; i++ {
//			numCh <- i
//			letterCh <- 'a' + rune(i-1)
//		}
//	}()
//
//	go func() {
//		for {
//			num, ok := <-numCh
//			if !ok {
//				break
//			}
//			fmt.Printf("%d\n", num)
//
//			letter, ok := <-letterCh
//			if !ok {
//				break
//			}
//			fmt.Printf("%c\n", letter)
//		}
//	}()
//
//	//select {} // 阻塞主 goroutine
//	//<-make(chan struct{})
//	time.Sleep(5 * time.Second)
//}

//func write(numCh chan int, letterCh chan rune) {
//	defer close(numCh)
//	defer close(letterCh)
//
//	for i := 1; i < 27; i++ {
//		numCh <- i
//		letterCh <- 'a' + rune(i-1)
//	}
//}
//
//func read(numCh chan int, letterCh chan rune) {
//	for {
//		if num, ok := <-numCh; ok {
//			fmt.Println(num)
//		} else {
//			break
//		}
//
//		if letter, ok := <-letterCh; ok {
//			fmt.Println(string(letter))
//		} else {
//			break
//		}
//	}
//}
//
//func waitgroups() {
//	numCh, letterCh := make(chan int), make(chan rune)
//	go write(numCh, letterCh)
//	go read(numCh, letterCh)
//	time.Sleep(time.Second)
//}

//package waitgroups
//
//import (
//	"fmt"
//	"sync"
//)
//
//func waitgroups() {
//	var wg sync.WaitGroup
//	wg.Add(2)
//
//	numCh := make(chan int)
//	letterCh := make(chan rune)
//
//	go func() {
//		defer wg.Done()
//		defer close(numCh)
//		defer close(letterCh)
//		for i := 1; i <= 26; i++ {
//			numCh <- i
//			letterCh <- 'a' + rune(i-1)
//		}
//		//close(numCh)
//		//close(letterCh)
//	}()
//
//	go func() {
//		defer wg.Done()
//		for {
//			num, ok := <-numCh
//			if !ok {
//				break
//			}
//			fmt.Printf("%d", num)
//
//			letter, ok := <-letterCh
//			if !ok {
//				break
//			}
//			fmt.Printf("%c", letter)
//		}
//	}()
//
//	wg.Wait()
//}

//package waitgroups
//
//import (
//	"fmt"
//	"time"
//)
//
//func waitgroups() {
//	ch1 := make(chan int, 5)
//	ch2 := make(chan int, 5)
//	go func() {
//		for i := 1; i < 6; i++ {
//			ch1 <- i
//			ch2 <- <-ch1
//		}
//		close(ch1)
//	}()
//
//	go func() {
//		for {
//			time.Sleep(time.Second)
//			msg, ok := <-ch2
//			if ok {
//				fmt.Println(msg)
//			} else {
//				fmt.Println("none")
//				fmt.Println(msg)
//				close(ch2)
//				return
//			}
//		}
//	}()
//	time.Sleep(10 * time.Second)
//}

//package waitgroups
//
//import (
//	"fmt"
//	"time"
//)
//
//func worker(id int, jobs chan int, result chan int) {
//	for j := range jobs {
//		fmt.Println("worker ", id, " processing job ", j)
//		time.Sleep(time.Second)
//		result <- j
//	}
//}
//func waitgroups() {
//	jobs := make(chan int, 100)
//	results := make(chan int, 100)
//
//	for i := 1; i <= 2; i++ {
//		go worker(i, jobs, results)
//	}
//
//	for j := 1; j <= 9; j++ {
//		jobs <- j
//	}
//	close(jobs)
//
//	for {
//		time.Sleep(time.Second)
//		resmsg, ok := <-results
//		jobmsg, ok1 := <-jobs
//		if ok {
//			fmt.Println("res:", resmsg)
//		} else {
//			fmt.Println("no", resmsg)
//		}
//		if ok1 {
//			fmt.Println("jobgs:", jobmsg)
//		} else {
//			fmt.Println("none", jobmsg)
//		}
//	}
//}

package main

func main() {
	ch1 := make(chan int, 1) // 创建一个带有缓冲区大小为1的通道

	ch1 <- 1 // 向带缓冲区的通道发送数据

	//<-ch1 // 向带缓冲区的通道发送数据
}
