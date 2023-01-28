package main

import (
	"fmt"
	"sync"
)

/*
同步的goroutine
由于goroutine是异步执行的，
那很有可能出现主程序退出时还有goroutine没有执行完，
此时goroutine也会跟着退出。此时如果想等到所有goroutine任务执行完毕才退出，go提供了sync包和channel来解决同步问题，当然如果你能预测每个goroutine执行的时间，你还可以通过time.Sleep方式等待所有的groutine执行完成以后在退出程序(如上面的列子)。

示例一：使用sync包同步goroutine
sync大致实现方式
WaitGroup
等待一组goroutinue执行完毕.
主程序调用 Add 添加等待的goroutinue数量.
每个goroutinue在执行结束时调用 Done ，此时等待队列数量减1.，主程序通过Wait阻塞，直到等待队列为0.
*/

func main() {

	var go_sync sync.WaitGroup

	for i := 0; i < 10; i++ {
		go_sync.Add(1)
		go cal(i, i+1, &go_sync)
	}

	go_sync.Wait()

}

func cal(i int, i2 int, s *sync.WaitGroup) {

	c := i + i2
	fmt.Printf(" %d + %d =%d \n", i, i2, c)
	defer s.Done()

}
