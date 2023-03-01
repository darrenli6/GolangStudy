package znet

import (
	"fmt"
	"github.com/darrenli6/GolangStudy/GoAdvanced/framework/zinx/ziface"
	"net"
)

type Connection struct {
	// 当前的连接
	Conn *net.TCPConn

	ConnID uint32

	isClosed bool

	// 退出
	ExitChan chan bool

	Router ziface.IRouter
}

// 初始化的方法
func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		Router:   router,
		isClosed: false,
		ExitChan: make(chan bool, 1),
	}
	return c
}

// 启动链接
func (c *Connection) Start() {

	fmt.Println("Conn start()  connID", c.ConnID)

	// 启动读数据的业务
	// TODO 启动从当前连接 读写数据的业务
	go c.StartReader()

}

// 停止链接  结束
func (c *Connection) Stop() {
	fmt.Println("Conn Stop() connId =", c.ConnID)

	if c.isClosed == true {
		return
	}

	c.isClosed = true
	c.Conn.Close()
	close(c.ExitChan)

}

// 获取当前连接绑定的socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {

	return c.Conn
}

// 获取当前连接的id
func (c *Connection) GetConnID() uint32 {

	return c.ConnID
}

// 获取远程客户端的 TCP状态 Ip port
func (c *Connection) RemoteAddr() net.Addr {

	return c.Conn.RemoteAddr()
}

// 发送数据 将数据发送给远程的客户端
func (c *Connection) Send(data []byte) error {

	return nil
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running ")
	defer fmt.Println("connID = ", c.ConnID, " reader is exit, remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err", err)
			continue
		}

		// 得到当前conn的request
		req := Request{
			conn: c,
			data: buf,
		}

		// 执行注册的路由方法
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
		// 从路由中 找到注册绑定的conn对应的router调用

	}
}
