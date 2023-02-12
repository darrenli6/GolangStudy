package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

var printPid sync.Once

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:18888")

	if err != nil {
		log.Fatalln(err)
		return
	}
	defer listener.Close()

	for {
		// 无论循环多少次 Do方法执行执行
		printPid.Do(func() {
			fmt.Println("当前的进程是", os.Getpid())
		})

		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("监听端口错误", err)

		}

		log.Println(conn.RemoteAddr(), " link successfully")
	}

}

/*
sync.Once结构体内部包含了一个互斥锁和布尔值，
互斥锁保证布尔值和数据的安全
，而布尔值用来记录初始化是否完成，

　　这样设计就能保证初始化操作的时候是并发安全的
,并且初始化操作也不会被执行多次
*/
