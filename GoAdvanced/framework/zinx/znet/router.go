package znet

import "github.com/darrenli6/GolangStudy/GoAdvanced/framework/zinx/ziface"

/*
实现Router时候 先嵌入这个BaseRouter基类
然后根据需要对这个基类的方法进行重写就好了
*/
type BaseRouter struct {
}

// 这里之所以BaseRouter为null 有的router不希望有PreHandle

// 在处理conn业务之前的狗子方法Hook
func (br *BaseRouter) PreHandle(request ziface.IRequest) {

}

// 在处理conn业务的狗子方法Hook
func (br *BaseRouter) Handle(request ziface.IRequest) {

}

// 在处理conn业务之后的狗子方法Hook
func (br *BaseRouter) PostHandle(request ziface.IRequest) {

}