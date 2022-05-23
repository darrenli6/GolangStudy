package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("eat ... ")
}

func (this *Human) Walt() {
	fmt.Println("Walt ... ")
}

type SuperMan struct {
	Human // superman类 继承了Human

	level int
}

func (this *SuperMan) Eat() {
	fmt.Println("superman  eat ...")
}

func (this *SuperMan) Fly() {
	fmt.Println(" fly ... ")
}



func main() {
	h := Human{"zhangsan", "male"}
	h.Eat()
	h.Walt()

	// 定义一个子类对象
	// s := SuperMan{Human{"lisi", "female"}, 99}
	var s SuperMan
	s.name = "lisi"
	s.sex = "femal"
	s.level = 11

	s.Walt()
	s.Eat()
	s.Fly()

}
