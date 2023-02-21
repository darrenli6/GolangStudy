package main

import "fmt"

func adder() func() int {
	i := 10
	return func() int {
		return i + 1
	}
}

func adder1() func(x int) int {
	i := 10
	return func(x int) int {
		i = i + x
		return i
	}
}

func main() {
	fn := adder()
	fmt.Println(fn()) //11
	fmt.Println(fn()) //11

	fn1 := adder1()
	fmt.Println(fn1(10)) //20
	fmt.Println(fn1(10)) //30
}
