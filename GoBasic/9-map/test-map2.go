package main

import "fmt"

func main() {

	cityMap := make(map[string]string)

	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "Newyork"

	//遍历
	for key, value := range cityMap {
		fmt.Println("key= ", key, "  value=", value)
	}

	// 删除

	delete(cityMap, "China")

	//修改

	cityMap["USA"] = "huashengdun"

	fmt.Println("-------")

	for key, value := range cityMap {
		fmt.Println("key= ", key, "  value=", value)
	}

}
