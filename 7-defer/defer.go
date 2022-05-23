package main

import "fmt"

func main() {

	// defer是栈的格式
	defer fmt.Println("main defer1")
	defer fmt.Println("main defer2")

	fmt.Println("main : hello go 1")
	fmt.Println("main : hello go 2")
}
