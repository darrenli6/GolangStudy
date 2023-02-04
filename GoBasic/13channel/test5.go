package main

import "fmt"

func main() {

	taskch := make(chan int, 20) // 任务
	resch := make(chan int, 20)  //结果
	exitch := make(chan bool, 5) // 退出

	go func() {
		for i := 0; i < 10; i++ {
			taskch <- i
		}
		close(taskch)
	}()

	for i := 0; i < 5; i++ {
		go Task(taskch, resch, exitch)
	}

	go func() {
		for i := 0; i < 5; i++ {
			<-exitch
		}

		close(resch)
		close(exitch)
	}()

	for res := range resch {
		fmt.Println("task res ", res)
	}
}

func Task(taskch chan int, resch chan int, exitch chan bool) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("do task eerror", err)
			return
		}
	}()

	for t := range taskch {
		fmt.Println("do task", t)
		resch <- t
	}
	exitch <- true
}
