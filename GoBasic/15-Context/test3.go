package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()
	go HelloHandler(ctx, 500*time.Millisecond)

	select {
	case <-ctx.Done():
		fmt.Println("hello handle", ctx.Err())
	}

}

func HelloHandler(ctx context.Context, duration time.Duration) {

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-time.After(duration):
		fmt.Println("process requiest with", duration)
	}
}
