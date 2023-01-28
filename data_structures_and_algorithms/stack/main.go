package main

import (
	"fmt"
	"sync"
)

type ArrayStack struct {
	Array []string
	Size  int
	Lock  sync.Mutex
}

func (s *ArrayStack) ArrPush(v string) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	s.Array = append(s.Array, v)
	s.Size++
}

func (s *ArrayStack) ArrPop() string {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	if s.Size == 0 {
		return ""
	}
	val := s.Array[s.Size-1]
	s.Array = s.Array[:s.Size-1]
	s.Size--
	return val
}

// 3. 获取栈顶元素
func (s *ArrayStack) GetArrTopVal() string {
	if s.Size == 0 {
		return ""
	}
	return s.Array[s.Size-1]
}

// 4. 获取栈的大小
func (s *ArrayStack) GetArrStackSize() int {
	return s.Size
}

// 5. 判断栈是否为空
func (s *ArrayStack) ArrStackIsEmpty() bool {
	return s.Size == 0
}
func main() {
	arrayStack := new(ArrayStack)
	arrayStack.ArrPush("cat")
	arrayStack.ArrPush("dog")
	arrayStack.ArrPush("hen")
	fmt.Println("arrayStack:", arrayStack.Array)
	fmt.Println("size:", arrayStack.GetArrStackSize())
	fmt.Println("pop:", arrayStack.ArrPop())
	fmt.Println("pop:", arrayStack.ArrPop())
	fmt.Println("arrayStack:", arrayStack.Array)
	fmt.Println("size:", arrayStack.GetArrStackSize())
	arrayStack.ArrPush("drag")
	fmt.Println("arrayStack:", arrayStack.Array)
	fmt.Println("pop:", arrayStack.ArrPop())
	fmt.Println("arrayStack:", arrayStack.Array)
}
