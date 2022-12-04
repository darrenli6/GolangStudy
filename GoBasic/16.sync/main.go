package main

import (
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	queue []string
	cond  *sync.Cond
}

func (q *Queue) Enqueue(s string) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.queue = append(q.queue, s)
	fmt.Printf("put string %s  notidy all \n", s)
	q.cond.Broadcast()
}

func (q *Queue) Dequeue() string {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	if len(q.queue) == 0 {
		fmt.Printf("please wait \n")
		q.cond.Wait()
	}
	result := q.queue[0]
	fmt.Printf("get %s ", result)
	q.queue = q.queue[1:]
	return result
}

func main() {
	q := Queue{
		queue: []string{},
		cond:  sync.NewCond(&sync.Mutex{}),
	}

	// 生产者
	go func() {
		for {
			q.Enqueue("hello")
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		q.Dequeue()
		time.Sleep(time.Second)
	}

}
