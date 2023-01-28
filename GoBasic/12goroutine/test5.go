package main

import (
	"fmt"
	"time"
)

/*
示例二：通过channel实现goroutine之间的同步。

实现方式：通过channel能在多个groutine之间通讯
，当一个goroutine完成时候向channel发送退出信号,
等所有goroutine退出时候，利用for循环channe去channel中
的信号，若取不到数据会阻塞原理，
等待所有goroutine执行完毕，使用该方法有个前提是你已经知道了你启动了多少个goroutine。
*/
const number = 10

func main() {

	exitChan := make(chan bool, 10)

	for i := 0; i < number; i++ {
		go cal1(i, i+1, exitChan)
	}
	for i := 0; i < number; i++ {
		<-exitChan
	}

	close(exitChan)

}

func cal1(i int, i2 int, exitChan chan bool) {
	c := i + i2
	fmt.Printf("%d + %d =%d \n", i, i2, c)
	time.Sleep(time.Second * 2)
	exitChan <- true

}
