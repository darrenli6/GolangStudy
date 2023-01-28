package main

import (
	"fmt"
	"sync"
)

/*
goroutine之间的通讯
goroutine本质上是协程，可以理解为不受内核调度，而受go调度器管理的线程。goroutine之间可以通过channel进行通信或者说是数据共享，当然你也可以使用全局变量来进行数据共享。

示例：使用channel模拟消费者和生产者模式
*/

func main() {

	datachan := make(chan int, 100) // 通讯信息的管道

	var wg sync.WaitGroup

	// 生产者
	for i := 0; i < 10; i++ {
		go Productor(datachan, i, &wg)
		wg.Add(1)
	}

	// 消费者
	for i := 0; i < 10; i++ {
		go Consumer(datachan, &wg)
		wg.Add(1)
	}

	wg.Wait()

}

func Productor(datachan chan int, i int, s *sync.WaitGroup) {
	datachan <- i
	fmt.Println("生产数据 ", i)
	s.Done()
}

func Consumer(datachan chan int, s *sync.WaitGroup) {

	a := <-datachan
	fmt.Println("消费数据 ", a)
	s.Done()
}
