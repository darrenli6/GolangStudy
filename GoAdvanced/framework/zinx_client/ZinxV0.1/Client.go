package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("client start ...")

	// 1. 直接连接远程服务器 得到一个conn连接信息
	conn, err := net.Dial("tcp", "127.0.0.1:10010")
	if err != nil {
		fmt.Println("client start err ,exit")
		return
	}

	for {
		// 连接调用 write
		_, err := conn.Write([]byte("hello zinx v0.1"))
		if err != nil {
			fmt.Println("write conn err ", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error")
			return
		}
		fmt.Printf("server call back %s cnt =%d \n ", buf, cnt)

		time.Sleep(1 * time.Second)

	}

}
