package main

import (
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

// 定义一个算数类型
type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.B + args.B
	return nil
}
func main() {

	// 得到一个实例
	arith := new(Arith)
	// 注册到rpc服务
	rpc.Register(arith)
	//挂到http服务上
	rpc.HandleHTTP()

	l, _ := net.Listen("tcp", ":1234")
	http.Serve(l, nil)
}
