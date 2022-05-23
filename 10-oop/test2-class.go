package main

import "fmt"

// 如果类名首字母大写，表示其他包也能访问
//go 语言大小写是有区分的
type Hero struct {
	// 如果说类的属性首字母大写，表示该属性是对外能够访问，
	//否则只能类的内部能访问
	Name  string
	Ad    int
	level int
}

/*
func (this Hero) Show() {
	fmt.Println("hero name = ", this.Name)
	fmt.Println("hero ad = ", this.Ad)
	fmt.Println("hero level = ", this.Level)
}
func (this Hero) GetName() string {
	fmt.Println("Name =", this.Name)
	return this.Name
}
func (this Hero) setName(nweName string) {
	// 调用该对象的一个副本拷贝
	this.Name = nweName
}

*/

func (this *Hero) Show() {
	fmt.Println("hero name = ", this.Name)
	fmt.Println("hero ad = ", this.Ad)
	fmt.Println("hero level = ", this.level)
}
func (this *Hero) GetName() string {
	fmt.Println("Name =", this.Name)
	return this.Name
}
func (this *Hero) setName(nweName string) {
	// 调用该对象的引用
	this.Name = nweName
}

func main() {

	// 初始化

	hero := Hero{Name: "zhangsan", Ad: 100, level: 1}
	hero.Show()

	hero.setName("lisi")

	hero.Show()
}
