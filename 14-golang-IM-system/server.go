package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

// 监听message广播消息channel的goroutine 一旦有消息就发送给全部在线的user
func (this *Server) ListenMessage() {
	for {

		msg := <-this.Message

		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

type Server struct {
	Ip   string
	Port int

	// 在线用户的列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	// 消息广播的channel
	Message chan string
}

// 创建一个server接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

func (this *Server) Broadcast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg

	this.Message <- sendMsg
}

func (this *Server) Handle(conn net.Conn) {
	fmt.Println("连接建立成功")

	user := NewUser(conn, this)

	user.Online()

	// 监听用户是否是活跃的channel
	isLive := make(chan bool)

	// 接受客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("conn errror ", err)
				return
			}
			//提取用户的消息 取出 \n
			msg := string(buf[:n-1])
			// 将消息进行转播
			user.DoMessage(msg)

			//代表活跃
			isLive <- true

		}
	}()

	// 当前handler 阻塞
	// 定时器
	for {
		select {
		case <-isLive:
			// 重置定时器
			// 不做任何事情 为了激活select 更新下面定时器

		case <-time.After(time.Second * 100):
			// 超时
			// 将当前的客户端强制关闭
			user.SendMsg("你被踢了")

			// 销毁资源
			close(user.C)
			//关闭连接
			conn.Close()

			// 退出当前的handle
			return

		}
	}

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

	// 启动监听
	go this.ListenMessage()

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
