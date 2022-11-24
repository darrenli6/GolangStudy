package main

import "fmt"

func main() {

	//定义一个 channel
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutne结束")

		fmt.Println("goroutine 正在运行")

		c <- 666 //将666 发送给c  阻塞
	}()

	// 从c接受数据  发生阻塞
	num := <-c

	fmt.Println("numer =", num)

	fmt.Println("main goroutine 正在运行")

}
