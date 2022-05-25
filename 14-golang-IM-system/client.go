package main

import (
	"flag"
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

var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器Ip地址")
	flag.IntVar(&serverPort, "port", 8888, "服务端端口")

}

func main() {

	// 服务器解析
	flag.Parse()

	client := NewClient(serverIp, serverPort)

	if client == nil {
		fmt.Println(">>> 链接服务器失败")
		return
	}

	fmt.Println(">>> 链接服务器成功")

	// 启动客户端的业务
	select {}
}
