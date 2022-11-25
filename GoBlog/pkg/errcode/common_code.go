package errcode

var (
	Success       = NewError(0, "成功")
	ServerError   = NewError(100000000, "服务内部错误")
	InvalidParams = NewError(10000001, "入参错误")
)
