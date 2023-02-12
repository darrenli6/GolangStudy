package main

import (
	"fmt"
	"strconv"
	"sync"
)

//　　go语言中内置的map不是并发安全的

var m = make(map[string]int)

func get(key string) int {
	return m[key]
}
func set(key string, value int) {
	m[key] = value
}

func main1() {
	// 内置的map 是线程不安全的
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			key := strconv.Itoa(i)
			set(key, i)
			fmt.Printf("key:%v, value:%v", key, get(key))
		}(i)
	}

	wg.Wait()
}

func main() {

	var (
		m  = sync.Map{}
		wg = sync.WaitGroup{}
	)

	for i := 0; i < 20; i++ {
		wg.Add(1)

		go func(i int) {
			key := strconv.Itoa(i)

			m.Store(key, i)
			v, ok := m.Load(key)
			if ok {
				fmt.Printf("key:%v, value:%v", key, v)
			}
			defer wg.Done()

		}(i)
	}

	wg.Wait()
}
