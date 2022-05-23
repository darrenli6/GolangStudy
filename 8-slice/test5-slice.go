package main

import "fmt"

func main() {

	s := []int{1, 2, 3} // len =3 cap=3

	s1 := s[0:2] //1,2

	fmt.Println(s1)
	// 非常容易混淆
	s1[0] = 100
	fmt.Println(s1)
	fmt.Println(s)

	// copy  可以将底层的数组拷贝
	s2 := make([]int, 3)
	// 将s的内容拷贝到s2
	copy(s2, s)
	fmt.Println(s2)

}
