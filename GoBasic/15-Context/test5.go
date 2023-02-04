package main

import (
	"context"
	"fmt"
	"time"
)

/*
WithTimeout
此函数类似于 context.WithDeadline。
不同之处在于它将持续时间作为参数输入而不是时间对象。
此函数返回派生 context，
如果调用取消函数或超出超时持续时间，
则会取消该派生 context。
*/

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	go Monitor(ctx)

	time.Sleep(20 * time.Second)
}

func Monitor(ctx context.Context) {

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-time.After(20 * time.Second):
		fmt.Println("stop monitor")
	}
}