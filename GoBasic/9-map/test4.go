package main

import (
	"fmt"
	"sync"
)

func main() {

	var m map[int]int
	m = make(map[int]int)

	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m[i] = i * i
			mu.Unlock()

		}(i)
	}

	wg.Wait()

	for i := 0; i < 1000; i++ {
		mu.Lock()
		v, ok := m[i]
		mu.Unlock()
		if ok {
			fmt.Printf("m[%d] = %d \n", i, v)
		} else {
			fmt.Printf("m[%d] not found \n", i)
		}

	}

}

/*
如果在遍历过程中出现了 "not found" 的提示信息，那么可能是因为某些 goroutine 还没有将元素添加到 map 中，就开始了遍历操作，导致某些键值对无法被遍历到。
*/
