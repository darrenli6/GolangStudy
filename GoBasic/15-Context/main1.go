package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var exitChan = make(chan bool, 1)

//为什么需要context

func f() {

	defer wg.Done()
LOOP:
	for {
		fmt.Println("hi")
		time.Sleep(time.Second * 1)
		select {
		case <-exitChan:
			break LOOP
		default:

		}

	}
}
func main() {
	// 初始化channel

	wg.Add(1)
	go f()

	time.Sleep(time.Second * 5)
	exitChan <- true

	wg.Wait()

	// 如何通知子goroutine 退出

}
