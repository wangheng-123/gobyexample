package main

import "fmt"

func main() {
	// `var` 声名1个或者多个变量
	var a = "initial"
	fmt.Println(a)

	//可以一次声明多个变量
	var b, c int = 1, 2
	fmt.Println(b, c)

	//go 将自动推断已经初始化的变量类型
	var d = true
	fmt.Println(d)

	//声明后却没有给出对应的初始值时，变量会初始化为零值
	// 例如 int 的零值是 0
	var e int
	fmt.Println(e)

	//:= 语法是声明并初始化变量的简写
	f := "short" // 等价于 var f string = "short"
	fmt.Println(f)
}
