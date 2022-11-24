package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)

	fmt.Println("len(c)", len(c), ", cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("sub goroutine exit ")
        // 如果=4 这样就会阻塞
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("sub goroutine 正在运行,ele= ", i, " len(c)", len(c), "  cap(c)=", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		numer := <-c
		fmt.Println("number = ", numer)

	}
	fmt.Println("main goroutine exit ")

}
