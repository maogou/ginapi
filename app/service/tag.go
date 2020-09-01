package service

//form表示表单映射字段名,binding表示验证规则

//获取tag的参数验证
type CountTagRequest struct {
	State uint8 `form:"state",default=1 binding:"oneof=0 1"`
}

//创建tag的参数验证
type CreateTagRequest struct {
	Name string `form:"name" binding:"max=100"`
	State uint8 `form:"state",default=1 binding:"oneof=0 1"`
}

//更新tag的参数验证
type UpdateTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
	Name string `form:"name" binding:"min=3,max=100"`
	State uint8 `form:"state" binding:"required,oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

//删除tag的参数验证
type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,get=1"`
}
