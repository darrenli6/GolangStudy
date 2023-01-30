package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

/*
context包可以提供一个请求从API请求边界到各goroutine的请求域数据传递、取消信号及截至时间等能力。
*/

func main() {
	http.HandleFunc("/", SayHello1)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func SayHello1(writer http.ResponseWriter, request *http.Request) {
	//每个请求对应的Handler，常会启动额外的的goroutine进行数据查询或PRC调用
	fmt.Println(&request)

	// 没有对声明周期进行管理 context进行了管理
	go func() {
		for range time.Tick(time.Second) {

			select {
			case <-request.Context().Done():
				fmt.Println("requist is done")
				return
			default:
				fmt.Println("current request is in progress ")

			}

		}
	}()
	time.Sleep(2 * time.Second)

	writer.Write([]byte("hi new request come"))

}
