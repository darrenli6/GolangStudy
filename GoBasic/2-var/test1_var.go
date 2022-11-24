package main

import "fmt"

var gA int = 100
var gB int = 200

// 全局 不支持
// gC:=11

func main() {
	// 声明一个变量，默认的值是0
	var a int
	fmt.Println("a =", a)
	fmt.Printf("type of a = %T\n", a)

	// 方法二 声明一个变量 初始化一个值
	var b int = 100
	fmt.Println("b =", b)
	fmt.Printf("type of b =%T \n", b)

	var bb string = "abcd"
	fmt.Printf("bb = %s , type pf bb = %T \n", bb, bb)

	// 方法三 初始化的时候 可以省去数据类型 通过自动匹配类型
	var c = 100
	fmt.Println("c =", c)
	fmt.Printf("type of c =%T \n", c)

	var cc = "abcdg"
	fmt.Printf("cc = %s , type pf cc = %T \n", cc, cc)

	// 方法四   省去var关键字  直接自动匹配 常用的方法
	e := 100
	fmt.Println("e= ", e)
	fmt.Printf("type of e = %T \n", e)

	f := "abcd"
	fmt.Println(" f = ", f)

	// ==
	fmt.Println("ga=", gA, "  gb = ", gB)

	// 声明多个变量
	var xx, yy int = 100, 29

	fmt.Println("xx =", xx, " yy =", yy)

	var kk, ll = 100, " darren"
	fmt.Println("ll =", ll, " kk =", kk)
	// 多行的多变量声明
	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println("vv = ", vv, ", jj =", jj)

}
