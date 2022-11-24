package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100

	return c
}

// 返回多个返回值 匿名的
func foo2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	return 666, 777

}

//返回多个返回值 有形参名称的
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("----foo30---")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	//r1 r2 属于foo3的形参，默认值是0 防止野指针
	//r1 r2 作用域是foo3 整个函数体内
	fmt.Println("r1 = ", r1)
	fmt.Println("r2 = ", r2)
	// 有名称的返回变量
	r1 = 1000
	r2 = 2000

	return
}

func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("----foo4---")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	// 有名称的返回变量
	r1 = 1000
	r2 = 2000

	return
}

func main() {

	c := foo1("aaaa", 55)

	fmt.Println("c = ", c)

	r1, r2 := foo2("aaaa", 999)

	fmt.Println("r1 = ", r1, "  r2=", r2)

	r1, r2 = foo3("foo3", 999)

	fmt.Println("r1 = ", r1, "  r2=", r2)

	r1, r2 = foo4("foo4", 999)

	fmt.Println("r1 = ", r1, "  r2=", r2)

}
