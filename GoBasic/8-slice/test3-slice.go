package main

import "fmt"

func main() {
	// 1 声明一个slice的切片，并且初始化，默认值是1，2,3

	// slice1 := []int{1, 2, 43, 5}

	//2 没有给出分配空间
	var slice1 []int
	// slice1 = make([]int, 3) //开辟三个空间
	// slice1[0] = 3

	// 3
	// var slice1 []int = make([]int, 3)

	// 4
	// slice1 := make([]int, 3)

	fmt.Printf("len =%d , slice =%v \n", len(slice1), slice1)

	// 判断slice是否为0
	if slice1 == nil {
		fmt.Println("slice is null 切片 ")
	} else {
		fmt.Println("slice是有空间的")
	}

}
