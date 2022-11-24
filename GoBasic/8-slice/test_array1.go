package main

import "fmt"

func printArray1(myArr []int) {
	// 切片是引用传递
	for _, v := range myArr {
		fmt.Println("---- ", v)
	}
	myArr[0] = 1000
}

func main() {

	myArr := []int{1, 2, 4, 6}

	fmt.Printf("myAarr is %T \n ", myArr)

	printArray1(myArr)
	fmt.Println("---")
	for _, v := range myArr {
		fmt.Println("---- ", v)
	}
}
