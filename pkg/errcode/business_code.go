package errcode

var (
	ErrorGetTagList = NewError(20010001, "获取标签列表失败")
	ErrorCreateTag  = NewError(20010002, "创建标签失败")
	ErrorUpdateTag  = NewError(20010003, "更新标签失败")
	ErrorDeleteTag  = NewError(20010004, "删除标签失败")
	ErrorCountTag   = NewError(20010005, "统计标签失败")
)
