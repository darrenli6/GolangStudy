package main

import "fmt"

type myint int

type Book struct {
	title string
	auth  string
}

func changeBook(book Book) {
	// 传递一个副本
	book.auth = "6666"

}

func changeBook1(book *Book) {
	// 传递一个副本
	book.auth = "6666"

}
func main() {
	/*
		var a myint = 10
		fmt.Println("a = ", a)

	*/

	var book1 Book
	book1.title = "go lang"
	book1.auth = "zhangsan"

	fmt.Printf("%v \n", book1)

	changeBook(book1)

	fmt.Printf("%v\n", book1)

	changeBook1(&book1)

	fmt.Printf("%v\n", book1)

	

}
