package main

import "fmt"

// func swap(a int, b int) {
// 	var temp int
// 	temp = a
// 	a = b
// 	b = temp
// }

func swap(pa *int, pb *int) {
	var temp int
	temp = *pa // temp ==main::a
	*pa = *pb
	*pb = temp

}

func main() {

	var a int = 10
	var b int = 20

	swap(&a, &b)
	fmt.Println("a =", a, " b=", b)

	var p *int
	p = &a

	fmt.Println(&a) //0xc000014080
	fmt.Println(p)  //0xc000014080

	var pp **int // 二级指针

	pp = &p

	fmt.Println(&p) //0xc00000e030
	fmt.Println(pp) //0xc00000e030

}
