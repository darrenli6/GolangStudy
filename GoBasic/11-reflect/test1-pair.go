package main

import "fmt"

func main() {

	var a string
	//pair <statictype:string,value:"abcd">
	a = "abcd"

	// pair<type:string,value:"abcd">
	var allType interface{}

	allType = a

	str, _ := allType.(string)

	fmt.Println(str)

}
