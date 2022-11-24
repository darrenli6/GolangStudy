package main

import "fmt"

func main() {

	// len = 3  cap=5

	var numbers = make([]int, 3, 5)

	fmt.Printf("len =%d ,cap=%d  slice=%v \n ", len(numbers), cap(numbers), numbers)

	// 向number追加一个元素
	numbers = append(numbers, 1)

	fmt.Printf("len =%d ,cap=%d  slice=%v \n ", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 1)
	numbers = append(numbers, 3)
	numbers = append(numbers, 3)
	numbers = append(numbers, 3)
	numbers = append(numbers, 3)

	fmt.Printf("len =%d ,cap=%d  slice=%v \n ", len(numbers), cap(numbers), numbers)

	// len 与cap不同
	// len是指左右指针的距离，容量是指左指针到底层数组末尾的距离。

}
