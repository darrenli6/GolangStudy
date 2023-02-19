package main

import "fmt"

type geo interface {
	area() int
}

type rect struct {
	with, height int
}

func (r rect) area() int {
	return r.with * r.height
}

type circle struct {
	radio int
}

func (c circle) area() int {
	return 3 * c.radio
}

func cal(geo2 geo) {
	fmt.Println(geo2.area())
}

func main() {
	r := rect{with: 2, height: 3}
	c := circle{radio: 4}
	cal(r)
	cal(c)
}
