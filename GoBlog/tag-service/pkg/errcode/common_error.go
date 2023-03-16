package errcode

var (
	Success = NewError(0, "成功")
	Fail    = NewError(1000000, "内部错误")
)
