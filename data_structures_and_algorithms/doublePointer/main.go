package main

import "fmt"

func main() {

	var v int = 100

	var pt1 *int = &v

	var pt2 **int = &pt1

	fmt.Println("变量v的值 = ",v)
	fmt.Println("变量v的地址值 = ",&v)

	fmt.Println("变量pt1的值 = ",pt1)
	fmt.Println("变量pt1的地址值 = ",&pt1)
	fmt.Println("变量pt2的值 = ",pt2)



}
