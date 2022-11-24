package main

import "fmt"

// const 来定义枚举
const (
	// 可以在const() 添加一个关键字 iota,每行的iota就会累加0 ，第一行的类iota默认值是0
	BEIJING = iota * 10
	SHANGHAI
	SHENZHEN
)

const (
	a, b = iota + 1, iota + 2 // iota =0 a=1 b =2
	c, d                      // iota=1 c=2 b=3

	e, f // iota=2 e=3 f=4

	g, h = iota * 2, iota * 3 // iota=3  g=6 h =9

	i, k
)

func main() {

	// 常量 只读属性

	const length int = 10
	fmt.Println("length = ", length)

	// length=1 常量是不允许修改的

	fmt.Println("北京 = ", BEIJING)
	fmt.Println("shanghai = ", SHANGHAI)
	fmt.Println("SHENZHEN = ", SHENZHEN)

	fmt.Println("h = ", h)
	//iota 只能配合const使用
	// var a int =iota

}
