package main

import (
	"fmt"
	"os"
)

// 假设我们想要创建一个文件，向它进行写操作，然后在结束时关闭它。
// 这里展示了如何通过 `defer` 来做到这一切
func main() {
	// 在 `createFile` 后得到一个文件对象，
	// 我们使用 defer 通过 `closeFile` 来关闭这个文件。
	// 这会在封闭函数（`http-clients`）结束时执行，就是 `writeFile` 结束后
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
