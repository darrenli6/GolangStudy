package ziface

// 定义一个服务器接口
type IServer interface {
	// 启动服务
	Start()
	// 停止服务
	Stop()
	// 运行服务
	Serve()

	// 路由功能 给当前的的服务注册一个路由功能
	AddRouter(router IRouter)
}
