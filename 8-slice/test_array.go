package main

import "fmt"

func main() {
	// 表示一个固定长度的数组
	var myArray1 [10]int
	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}

	myArray2 := [10]int{0, 1, 2, 3}
	for k, v := range myArray2 {
		fmt.Println("k = ", k, "v ", v)
	}
	//乐行
	fmt.Printf("myarray 1  type  is %T\n", myArray1)
	fmt.Printf("myarray 2  type  is %T\n", myArray2)

	myArray3 := [4]int{3, 45, 5, 34}
	fmt.Printf("myarray 3  type  is %T\n", myArray3)

	printArray(myArray3)

	for k, v := range myArray3 {
		fmt.Println("k = ", k, "  v=", v)
	}

}

func printArray(myArr [4]int) {
	for k, v := range myArr {
		fmt.Println("k = ", k, "  v=", v)
	}
	// 没有修改，值拷贝
	myArr[0] = 111
}
