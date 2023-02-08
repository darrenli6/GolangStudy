package main

import "github.com/darrenli6/GolangStudy/GoAdvanced/framework/zinx/znet"

func main() {

	s := znet.NewServe("zinx 0.2")
	s.Serve()
}
