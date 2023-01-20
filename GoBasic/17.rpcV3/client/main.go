package main

import (
	"fmt"
	"net/rpc"
)

type Args struct {
	A, B int
}

func main() {

	// 连接服务端 创造一个client
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	args := &Args{7, 8}
	var reply int

	// 通过call方法调用Arith类型的Multiply方法
	client.Call("Arith.Multiply", args, &reply)
	// 得到计算结果
	fmt.Printf("Arith : %d + %d = %d \n", args.A, args.B, reply)
}