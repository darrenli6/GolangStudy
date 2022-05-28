package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var exitChan = make(chan bool, 1)

//为什么需要context

func f2(ctx context.Context) {

}

func f(ctx context.Context) {

	defer wg.Done()
	go f2(ctx)
LOOP:
	for {
		fmt.Println("hi")
		time.Sleep(time.Second * 1)
		select {
		case <-ctx.Done():
			break LOOP
		default:

		}

	}
}

func main() {
	// 初始化channel
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go f(ctx)

	time.Sleep(time.Second * 5)

	// 通知
	cancel()
	wg.Wait()

	// 如何通知子goroutine 退出

}
