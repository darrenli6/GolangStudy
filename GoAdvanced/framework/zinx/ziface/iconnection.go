package ziface

import "net"

type IConnection interface {
	// 启动链接
	Start()

	// 停止链接  结束
	Stop()

	// 获取当前连接绑定的socket conn
	GetTCPConnection() *net.TCPConn

	// 获取当前连接的id
	GetConnID() uint32

	// 获取远程客户端的 TCP状态 Ip port
	RemoteAddr() net.Addr

	// 发送数据 将数据发送给远程的客户端
	Send(data []byte) error
}

type HandleFunc func(*net.TCPConn, []byte, int) error
