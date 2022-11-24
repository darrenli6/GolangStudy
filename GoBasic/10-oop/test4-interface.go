package main

import "fmt"

// 万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("my func is called")
	fmt.Println(arg)

	// interface{} 如何区分 此时引用的数据类型是什么呢 ？
	// 给 interface{} 提供了 “类型断言” 机制

	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string")
	} else {
		fmt.Println("arg is  string value=", value)
	}

}

type Book struct {
	auth string
}

func main() {
	book := Book{"GoLang"}

	myFunc(book)
	myFunc(100)
	myFunc("abk")

}
