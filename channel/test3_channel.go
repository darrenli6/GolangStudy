package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {

		for i := 0; i < 5; i++ {
			c <- i

		}

		close(c)
	}()

	 // range 迭代
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("main finishing ")
}
