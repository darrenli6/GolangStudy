package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", SayHello)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func SayHello(writer http.ResponseWriter, request *http.Request) {
	//每个请求对应的Handler，常会启动额外的的goroutine进行数据查询或PRC调用
	fmt.Println(&request)

	// 没有对声明周期进行管理 context进行了管理
	go func() {
		for range time.Tick(time.Second) {
			fmt.Println("current request is in progress ")
		}
	}()
	time.Sleep(2 * time.Second)

	writer.Write([]byte("hi new request come"))

}
