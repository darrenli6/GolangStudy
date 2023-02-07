package main

import "github.com/darrenli6/GolangStudy/GoAdvanced/framework/zinx/znet"

func main() {
	// 创建一个server的句柄
	s := znet.NewServe("v0.1")
	// 启动
	s.Serve()
}
