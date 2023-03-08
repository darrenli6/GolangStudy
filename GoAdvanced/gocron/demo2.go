package main

import (
	"errors"
	"github.com/jasonlvhit/gocron"
)

func task1() error {
	err := errors.New("error")
	// 处理逻辑
	if err != nil {
		return err
	}
	return nil
}

func main() {

	// 创建一个新的定时器
	s := gocron.NewScheduler()

	// 设置任务，并指定错误恢复和重试的次数
	s.Every(2).Seconds().Do(task1)

	// 启动定时器
	<-s.Start()
}
