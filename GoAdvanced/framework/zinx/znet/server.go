package znet

import (
	"fmt"
	"github.com/darrenli6/GolangStudy/GoAdvanced/framework/zinx/ziface"
	"net"
)

// IServer的接口实现  定义一个server的服务模块
type Server struct {
	// 服务器的名称
	Name string
	// 版本
	IpVersion string

	// 服务器绑定的ip地址
	Ip string

	// 服务器监听的端口
	Port int
	// 当前server添加一个router
	Router ziface.IRouter
}

func (s *Server) Start() {
	fmt.Printf("[start] Server Listenner at IP :%s ,Port %d , is staring \n", s.Ip, s.Port)

	go func() {
		// 1 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IpVersion, fmt.Sprintf("%s:%d", s.Ip, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error", err)
			return
		}
		// 2 监听服务器的地址
		listener, err := net.ListenTCP(s.IpVersion, addr)
		if err != nil {
			fmt.Println("listen ", s.IpVersion, " err", err)
			return
		}

		fmt.Println("start zinx server succ ", s.Name, " success Listenning")

		var cid uint32
		cid = 0

		// 3 阻塞等待客户端链接，处理客户端链接业务 读写
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("accept err", err)
				continue
			}
			// 已经与客户端链接，做一些的业务，做一个基本的最大512字节长度的回显业务

			dealConn := NewConnection(conn, cid, s.Router)
			cid++
			//启动当前的业务

			go dealConn.Start()

		}
	}()

}

func (s *Server) Stop() {
	// TODO 将一些服务器的资源 状态 或者将一些开启的连接信息，进行停止或者回收

}

func (s *Server) Serve() {
	// 启动server服务功能
	s.Start()

	// TODO 做一些启动服务器之后额外业务

	// 阻塞状态
	select {}

}
func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("Add Router Succ!!")

}

func NewServe(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IpVersion: "tcp4",
		Ip:        "0.0.0.0",
		Port:      10010,
		Router:    nil,
	}

	return s
}
