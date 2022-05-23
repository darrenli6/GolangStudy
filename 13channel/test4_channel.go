package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 1, 1

	for {
		select {
		//写入
		case c <- x:
			x = y
			y = x + y
			// 读
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func main() {

	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			// 读
			fmt.Println(<-c)
		}
		// 写
		quit <- 0
	}()
	fibonacci(c, quit)
}
