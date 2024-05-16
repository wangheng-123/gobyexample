package main

import "fmt"

/*
当使用通道作为函数的参数时，你可以指定这个通道是不是 只用来发送或者接收值。这个特性提升了程序的类型安全性
*/

// `ping` 函数定义了一个只允许发送数据的通道。尝试使用这个通
// 道来接收数据将会得到一个编译时错误
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// `pong` 函数定义了一个只允许接收来自通道（ping）的数据
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	//msg->pings->pongs
	pings := make(chan string, 1) //写数据channel
	pongs := make(chan string, 1) //读数据channel
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
