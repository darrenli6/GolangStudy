package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

// 结构文件的序列化和反序列化

type Address struct {
	City    string
	Country string
}

// 序列化数据的存放的路径
var filePath string

// 将数据序列号后写到文件中
func encode() {
	pa := Address{"beijing", "China"}

	// 打开文件
	file, _ := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)

	defer file.Close()

	//encode后写到这个文件中
	enc := gob.NewEncoder(file)
	enc.Encode(pa)
}

func decode() *Address {
	file, _ := os.Open(filePath)
	defer file.Close()

	var pa Address
	// decode操作
	dec := gob.NewDecoder(file)
	dec.Decode(&pa)
	return &pa
}

func main() {

	filePath = "./address.gob"
	encode()
	pa := decode()
	fmt.Println(*pa)
}
