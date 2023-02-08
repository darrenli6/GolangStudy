package ziface

/*
Irequest 接口
实际上客户端请求的连接信息，和请求的数据包装到request里面
*/

type IRequest interface {
	//得到当前链接
	GetConnection() IConnection

	// 得到数据
	GetData() []byte
}
