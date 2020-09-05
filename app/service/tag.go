package service

import (
	"github.com/maogou/ginapi/app/model"
	"github.com/maogou/ginapi/pkg/app"
)

//form表示表单映射字段名,binding表示验证规则

//获取tag的参数验证
type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state",default=1 binding:"oneof=0 1"`
}

//创建tag的参数验证
type CreateTagRequest struct {
	Name      string `form:"name" binding:"required,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=2,max=100"`
	State     uint8  `form:"state",default=1 binding:"oneof=0 1"`
}

//更新tag的参数验证
type UpdateTagRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"min=3,max=100"`
	State      uint8  `form:"state" binding:"required,oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

//删除tag的参数验证
type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,get=1"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

//标签总数
func (service Service) CountTag(param *CountTagRequest) (int, error) {
	return service.dao.CountTag(param.Name, param.State)
}

//标签列表
func (service Service) GetTagList(param *TagListRequest, page *app.Page) ([]*model.Tag, error) {
	return service.dao.GetTagList(param.Name, param.State, page.Page, page.PageSize)
}

//新增tag
func (service Service) CreateTag(param *CreateTagRequest) error {
	return service.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

//更新tag
func (service Service) UpdateTag(param *UpdateTagRequest) error {
	return service.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}

//删除tag
func (service Service) DeleteTag(param *DeleteTagRequest) error {
	return service.dao.DeleteTag(param.ID)
}
