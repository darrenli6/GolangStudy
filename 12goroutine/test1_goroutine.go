package main

import (
	"fmt"
	"time"
)

// 从goroutine
func newTack() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine i = %d\n", i)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	// 创建一个go协程
	go newTack()

	fmt.Printf("go routine exit\n")
	/*
		i := 0
		// main go routine
		for {
			i++
			fmt.Printf("main Goroutine i = %d\n", i)
			time.Sleep(time.Second * 1)
		}

	*/
}
