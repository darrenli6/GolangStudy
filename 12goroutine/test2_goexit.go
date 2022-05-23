package main

import (
	"fmt"
	"time"
)

func main() {

	// 用go创建一个一参数为空 返回值为空的函数
	/*
		go func() {
			defer fmt.Println("A.defer")

			func() {
				defer fmt.Println("B.defer")
				runtime.Goexit()
				fmt.Println("B")
			}()

			fmt.Println("A")

		}()
	*/

	go func(a int, b int) bool {

		fmt.Println("a = ", a, "   b = ", b)
		return true
	}(10, 20)

	for {
		time.Sleep(1 * time.Second)
	}
}
