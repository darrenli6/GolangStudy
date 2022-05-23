package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("read a book")
}

func (this *Book) WriteBook() {
	fmt.Println("write a book")
}

func main() {
	// pair <type:book,value:book{}地址 >
	b := &Book{}

	//r : pair<type:,value:>
	var r Reader
	r = b
	r.ReadBook()

	var w Writer
	// r:pair<type:book,value:book{}地址>
	w = r.(Writer) // 此处 为什么断言会成功，因为wr 具体类型type是一致的
	w.WriteBook()
}
