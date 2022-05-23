package main // 文件名与包名没有关系

import (
	"fmt"
	"time"
)

// 左花括号与方法一行
func main() {
	// golang中加“;” 与不加都可以 ， 建议不加
	fmt.Println("hello go")

	time.Sleep(1 * time.Second)
}
