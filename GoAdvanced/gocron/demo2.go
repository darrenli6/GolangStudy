package main

import (
	"fmt"
	"github.com/avast/retry-go"
	"github.com/go-co-op/gocron"
	"time"
)

func main() {
	// 创建一个新的调度器
	s := gocron.NewScheduler(time.Local)

	// 设置任务
	task := func() error {
		// 这里是您的任务代码
		//return fmt.Errorf("发生了一个错误")
		fmt.Println("执行任务")
		return nil
	}

	// 设置任务执行的时间间隔
	_, err := s.Every(2).Seconds().Do(func() {
		err := retry.Do(task, retry.Attempts(3), retry.Delay(time.Second))
		if err != nil {
			fmt.Println(err)
		}

	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// 启动调度器
	s.StartAsync()

	// 等待调度器退出
	s.StartBlocking()

}
