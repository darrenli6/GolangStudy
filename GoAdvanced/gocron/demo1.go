package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
)

func task() {
	fmt.Println("任务执行")
}

func main() {

	// 创建一个新的定时器
	s := gocron.NewScheduler()

	// 设置任务
	s.Every(2).Seconds().Do(task)

	// 启动定时器
	<-s.Start()
}
