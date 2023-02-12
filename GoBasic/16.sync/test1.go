package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	go hello(&wg)

	wg.Wait()

	fmt.Println("main over ")

}

//需要注意sync.WaitGroup是一个结构体，传递的时候要传递指针。

func hello(s *sync.WaitGroup) {
	defer s.Done()
	fmt.Println("helo")
}
