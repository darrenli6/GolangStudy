package ziface

/*
路由的抽象接口
路由里的数据都是IRequest

*/

type IRouter interface {
	// 在处理conn业务之前的狗子方法Hook
	PreHandle(request IRequest)
	// 在处理conn业务的狗子方法Hook
	Handle(request IRequest)

	// 在处理conn业务之后的狗子方法Hook
	PostHandle(request IRequest)
}
