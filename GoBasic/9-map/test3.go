package main

import (
	"fmt"
	"sync"
)

func main() {
	var m map[int]int
	m = make(map[int]int)

	var wg sync.WaitGroup

	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			m[i] = i * i
		}()
	}
	wg.Wait()
	for i := 0; i < 1000; i++ {
		fmt.Printf("m[%d] = %d \n", i, m[i])
	}
}
