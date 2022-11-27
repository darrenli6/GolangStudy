package errcode

var (
	Success       = NewError(0, "成功")
	ServerError   = NewError(100000000, "服务内部错误")
	InvalidParams = NewError(10000001, "入参错误")

	ErrorGetTagListFailed = NewError(20000001, "获取标签列表失败")
	ErrorCreateTagFailed  = NewError(20000002, "创建标签失败")
	ErrorUpdateTagFailed  = NewError(20000003, "更新标签失败")
	ErrorDeleteTagFailed  = NewError(20000004, "删除标签失败")
	ErrorCountTagFailed   = NewError(20000005, "统计标签失败")
)
