package main

import "fmt"

// 本质是一个指针
type AminalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

//具体的类
type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("cat sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "cat"
}

//具体的类
type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("dog sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(a AminalIF) {
	a.Sleep()
	fmt.Println("colort ", a.GetColor())
}
func main() {

	/*
		var animal AminalIF
		animal = &Cat{"Green"}

		animal.Sleep()

		animal = &Dog{"yello"}
		animal.Sleep()
	*/

	cat := Cat{"grey"}
	dog := Dog{"yellow"}

	showAnimal(&cat)
	showAnimal(&dog)

}
