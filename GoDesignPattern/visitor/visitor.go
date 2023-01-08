package main

// 先定义一个需要访问的数据结构
type Info struct {
	Namespace   string
	Name        string
	OtherThings string
}

type VisitorFunc func(*Info, error) error

// 将访问数据的方法抽象为接口
type Visitor interface {
	Visit(visitorFunc VisitorFunc) error
}

// info实现这个接口
func (i *Info) Visit(visitorFunc VisitorFunc) error {
	return visitorFunc(i, nil)
}
