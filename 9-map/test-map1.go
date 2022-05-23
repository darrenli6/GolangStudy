package main

import (
	"fmt"
)

func main() {
	// 声明map类型 key是string  value 是string
	var myMap map[string]string
	if myMap == nil {
		fmt.Println("mymap1 是一个空map")
	}
	// 开辟10 个内存空间
	myMap = make(map[string]string, 10)

	myMap["one"] = "java"
	myMap["two"] = "C++"

	fmt.Println(myMap)

	// 第二种声明方式

	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "python"
	fmt.Println(myMap2)
	// 第三种声明方式

	myMap3 := map[string]string{
		"one": "php",
		"two": "java",
	}
	fmt.Println(myMap3)

}
