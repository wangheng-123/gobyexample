package main

import (
	"encoding/base64"
	"fmt"
)

// 这个语法引入了 `encoding/base64` 包并使用名称 `b64` 代替默认的 `base64`。
// 这样可以节省点空间
func main() {
	// 这是将要编解码的字符串
	data := "abc123!?$*&()'-=@~"

	// Go 同时支持标准的和 URL 兼容的 base64 格式。
	// 编码需要使用 `[]byte` 类型的参数，所以要将字符串转成此类型
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// 解码可能会返回错误，如果不确定输入信息格式是否正确，
	// 那么，你就需要进行错误检查了
	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(sDec)
	fmt.Println()

	// 使用 URL 兼容的 base64 格式进行编解码
	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
