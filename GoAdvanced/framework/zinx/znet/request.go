package znet

import "github.com/darrenli6/GolangStudy/GoAdvanced/framework/zinx/ziface"

type Request struct {
	// 已经和客户端建立好链接
	conn ziface.IConnection

	// 客户端请求的数据
	data []byte
}

// 得到当前链接

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

// 得到数据
func (r *Request) GetData() []byte {
	return r.data
}
