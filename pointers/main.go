package main

import "fmt"

// 我们将通过两个函数：`zeroval` 和 `zeroptr` 来比较指针和值类型的不同。
// `zeroval` 有一个 `int` 型参数，所以使用值传递。
// `zeroval` 将从调用它的那个函数中得到一个 `ival` 形参的拷贝
func zeroval(ival int) {
	fmt.Println("&ival:", &ival)
	ival = 0
}

// `zeroptr` 有一和上面不同的 `*int` 参数，意味着它用了一个 `int` 指针。
// 函数体内的 `*iptr` 接着_解引用_这个指针，从它内存地址得到这个地址对应的当前值。
// 对一个解引用的指针赋值将会改变这个指针引用的真实地址的值
func zeroptr(iptr *int) {
	fmt.Println("iptr:", iptr)
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)
	fmt.Println("&i:", &i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// 通过 `&i` 语法来取得 `i` 的内存地址，即指向 `i` 的指针
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// 指针也是可以被打印的
	fmt.Println(&i)
}
