package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this *User) Call() {
	fmt.Println("user is called")
	fmt.Printf("%v\n", this)
}
func main() {

	user := User{1, "Aceid", 18}
	DofileAndMethod(user)

}

func DofileAndMethod(input interface{}) {
	// 获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("input tpye is ", inputType)
	// 获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("input value  is ", inputValue)
	// 通过type获取里面的字段
	// 1 通过interface的reflect.type 通过type获取NumFiled 进行遍历
	// 2 得到每个field 数据类型
	// 3 通过filed有一个interface方法等到对应的value

	for i := 0; i < inputType.NumField(); i++ {
		filed := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		// go语言的源码
		fmt.Printf("%s:  %v = %v \n", filed.Name, filed.Type, value)

	}

	// 通过type获取里面的方法调用

	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s : %v \n", m.Name, m.Type)
	}

}
