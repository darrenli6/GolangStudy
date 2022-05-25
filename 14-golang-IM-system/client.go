package main

import (
	"fmt"
	"net"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
}

func NewClient(serverIp string, serverPort int) *Client {

	// 创建

	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
	}

	//链接 server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net Dial error", err)
		return nil
	}

	client.conn = conn
	return client

}

func main() {
	client := NewClient("127.0.0.1", 8888)

	if client == nil {
		fmt.Println(">>> 链接服务器失败")
		return
	}

	fmt.Println(">>> 链接服务器成功")

	// 启动客户端的业务
	select {}
}
