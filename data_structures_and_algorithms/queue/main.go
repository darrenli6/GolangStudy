package main

import (
	"fmt"
	"sync"
)

type ArrQue struct {
	Qs   []string
	Size int
	Lock sync.Mutex
}

func (t *ArrQue) ArrAdd(data string) {
	t.Lock.Lock()
	defer t.Lock.Unlock()

	t.Qs = append(t.Qs, data)
	t.Size++
}

func (t *ArrQue) ArrPop() string {
	t.Lock.Lock()
	defer t.Lock.Unlock()
	if t.Size == 0 {
		return ""
	}
	v := t.Qs[0]
	t.Qs = t.Qs[1:]
	t.Size -= 1
	return v
}

func main() {
	tArr := new(ArrQue)
	tArr.ArrAdd("1")
	tArr.ArrAdd("2")
	r := tArr.ArrPop()
	fmt.Println(r)
}
