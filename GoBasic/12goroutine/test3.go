package main

import (
	"fmt"
	"runtime"
	"time"
)

func cal_runtime() {
	// 获取主机的逻辑CPU个数
	num := runtime.NumCPU()
	fmt.Printf(" 获取主机的逻辑CPU个数 ： %d \n", num)
	i := runtime.GOMAXPROCS(num)
	fmt.Printf("同时执行的最大CPU数 %d ", i)
}

func main() {

	//cal_runtime()

	// 对异常进行捕捉

	arr := make([]int, 4)
	for i := 0; i < 10; i++ {
		go addedle(arr, i)
	}
	time.Sleep(time.Second * 2)

}

func addedle(arr []int, i int) {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("add ele fail ")
		}
	}()

	arr[i] = i
	fmt.Println(arr)
}
