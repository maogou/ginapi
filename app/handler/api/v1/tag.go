package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/maogou/ginapi/app/service"
	"github.com/maogou/ginapi/global"
	"github.com/maogou/ginapi/pkg/app"
	"github.com/maogou/ginapi/pkg/errcode"
	"github.com/maogou/ginapi/pkg/transform"
)

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

// @Summary 获取多个标签
// @Produce  json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	//参数绑定验证
	param := service.TagListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)

	if !valid {
		global.Logger.Errorf(c, "tag List app.BindAndValid err:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetail(errs.Errors()...))
		return
	}

	//实例化service
	svc := service.New(c.Request.Context())
	page := app.Page{
		Page:     app.GetPage(c),
		PageSize: app.GetPageSize(c),
	}

	//获取总条数
	totalRows, err := svc.CountTag(&service.CountTagRequest{
		Name:  param.Name,
		State: param.State,
	})

	if err != nil {
		global.Logger.Errorf(c, "svc.CountTag err : %v", err)
		response.ToErrorResponse(errcode.ErrorCountTag)
		return
	}

	//获取tag列表
	tags, err := svc.GetTagList(&param, &page)

	if err != nil {
		global.Logger.Errorf(c, "svc.GetTagList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagList)
		return
	}

	//响应结果
	response.ToResponseList(tags, totalRows)
	return
}

// @Summary 新增标签
// @Produce  json
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string false "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	param := service.CreateTagRequest{}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)

	if !valid {
		global.Logger.Errorf(c, "tag Create app.BindAndValid err: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetail(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())

	err := svc.CreateTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CreateTag err:%v", err)
		response.ToErrorResponse(errcode.ErrorCreateTag)
		return
	}

	response.ToResponseMsg("创建标签成功")
	return
}

// @Summary 更新标签
// @Produce  json
// @Param id path int true "标签ID"
// @Param name body string false "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {
	//更新操作需要手动获取id构建form/post请求
	param := service.UpdateTagRequest{ID: transform.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "tag Update app.BindAndValid err:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetail(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.UpdateTag err:%v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTag)
		return
	}

	response.ToResponseMsg("更新标签成功")
	return
}

// @Summary 删除标签
// @Produce  json
// @Param id path int true "标签ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
	//删除操作需要手动获取id构建form/post请求
	param := service.DeleteTagRequest{ID: transform.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)


	valid, errs := app.BindAndValid(c, &param)

	if !valid {
		global.Logger.Errorf(c, "tag Delete app.BindAndValid err:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetail(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())

	err := svc.DeleteTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.DeleteTag err:%v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTag)
		return
	}

	response.ToResponseMsg("删除标签成功")
	return
}
