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
	flag       int
}

func NewClient(serverIp string, serverPort int) *Client {

	// 创建

	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
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

func (client *Client) menu() bool {
	var flag int
	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {

		client.flag = flag
		return true

	} else {
		fmt.Println("请输入合法的数字>>>>>")
		return false
	}

}

func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {

		}

		// 根据不同模式处理不同的业务
		switch client.flag {
		case 1:
			//公聊
			fmt.Println("公聊模式选择")
			break
		case 2:
			fmt.Println("私聊模式选择")
			break
		case 3:
			fmt.Println("更新用户名")
			break
		}
	}
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
	client.Run()
}
