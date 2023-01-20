package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
)

// 准备编码的数据
type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

func encode(data interface{}) *bytes.Buffer {
	// buffer 类型实现io.write的接口
	var buf bytes.Buffer
	// 得到编码器
	enc := gob.NewEncoder(&buf)
	// 用encode来编码data
	enc.Encode(data)

	return &buf
}
func decode(data interface{}) *Q {
	d := data.(io.Reader)
	// 获取解码器
	dec := gob.NewDecoder(d)
	var q Q
	dec.Decode(&q)
	return &q
}

func main() {

	data := P{2, 3, 4, "darren"}
	buf := encode(data)
	var q *Q
	q = decode(buf)
	fmt.Printf("%q: {%d,%d}\n", q.Name, *q.X, *q.Y)
}
