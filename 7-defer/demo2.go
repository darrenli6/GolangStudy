package main

import "fmt"

/*
知识点二  defer和return 谁先谁后
return 先执行 defer后执行
*/

func deferFunc() int {

	fmt.Println("defer func")
	return 0
}

func returnFunc() int {
	fmt.Println("return func")
	return 0
}

func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	returnAndDefer()
}
