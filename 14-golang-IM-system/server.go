package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

// 创建一个server接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}
	return server
}

func (this *Server) Handle(conn net.Conn) {
	fmt.Println("连接建立成功")
}

// 启动服务器的接口
func (this *Server) Start() {
	// socket listen

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	// close
	defer listener.Close()

	for {
		// accept

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener acccept error ", err)
			continue
		}

		// do handle
		go this.Handle(conn)

	}

	// close listen socket

}
