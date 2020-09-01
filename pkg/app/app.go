package app

import (
	"github.com/gin-gonic/gin"
	"github.com/maogou/ginapi/pkg/errcode"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
}

type Page struct {
	//页码
	Page int `json:"page"`
	//每页数量
	PageSize int `json:"page_size"`
	//总行数
	TotalRows int `json:"total_rows"`
}

//实例化Response
func NewResponse(ctx *gin.Context) *Response  {
	return &Response{
		Ctx: ctx,
	}
}

//响应正常json
func (r *Response) ToResponse(data interface{})  {
	if data == nil {
		data = gin.H{}
	}

	r.Ctx.JSON(http.StatusOK,data)
}

//响应分页json
func (r *Response) ToResponseList(list interface{},totalRows int)  {
	r.Ctx.JSON(http.StatusOK,gin.H{
		"list":list,
		"pager":Page{
			Page:GetPage(r.Ctx),
			PageSize: GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

//错误响应
func (r *Response) ToErrorResponse(err *errcode.Error)  {
	response := gin.H{
		"code":err.Code(),
		"msg":err.Msg(),
	}

	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}

	r.Ctx.JSON(err.StatusCode(),response)
}
