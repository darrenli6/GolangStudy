package main

import (
	"fmt"
	"strconv"
)

func main() {

	var v int = 100

	var pt1 *int = &v

	var pt2 **int = &pt1

	fmt.Println("变量v的值 = ", v)
	fmt.Println("变量v的地址值 = ", &v)

	fmt.Println("变量pt1的值 = ", pt1)
	fmt.Println("变量pt1的地址值 = ", &pt1)
	fmt.Println("变量pt2的值 = ", pt2)

	// 解引用
	fmt.Println("pt2地址值是 ", *pt2)

	fmt.Println("*(pt2地址)值是 ", **pt2)

	demo1()

}

func demo1() {

	var v int = 100

	var pt1 *int = &v

	var pt2 **int = &pt1

	fmt.Println("v = ", v)

	*pt1 = 200

	fmt.Println("更新pt1之后 v的值", v)

	**pt2 = 300
	fmt.Println("更新pt2之后 v的值", v)

	b := isPalindromeInt(11311)
	fmt.Println(b)

}

func isPalindromeInt(x int) bool {
	// 先转成字符
	str := strconv.Itoa(x)
	strBytes := []rune(str)

	left := 0
	right := len(strBytes) - 1
	for left < right {
		if strBytes[left] != strBytes[right] {
			return false
		}
		left++
		right--
	}
	return true
}
