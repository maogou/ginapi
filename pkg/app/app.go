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

type SuccessList struct {
	List  interface{} `json:"list"`
	Pager Page        `json:"pager"`
}

//实例化Response
func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		Ctx: ctx,
	}
}

//响应修改-创建-删除等操作的json
func (r *Response) ToResponseMsg(msg string) {
	success := errcode.Success
	response := gin.H{
		"code": success.Code(),
		"msg":  msg,
	}

	r.Ctx.JSON(http.StatusOK, response)
}

//响应单条正常json
func (r *Response) ToResponse(data interface{}) {
	success := errcode.Success
	if data == nil {
		data = gin.H{}
	}
	response := gin.H{
		"code": success.Code(),
		"msg":  success.Msg(),
		"data": data,
	}

	r.Ctx.JSON(http.StatusOK, response)
}

//响应多条分页json
func (r *Response) ToResponseList(list interface{}, totalRows int) {
	success := errcode.Success
	response := gin.H{
		"code": success.Code(),
		"msg":  success.Msg(),
		"data": SuccessList{
			List: list,
			Pager: Page{
				Page:      GetPage(r.Ctx),
				PageSize:  GetPageSize(r.Ctx),
				TotalRows: totalRows,
			},
		},
	}

	r.Ctx.JSON(http.StatusOK, response)
}

//错误响应
func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{
		"code": err.Code(),
		"msg":  err.Msg(),
	}

	details := err.Details()
	if len(details) > 0 {
		response["error"] = details
	}

	r.Ctx.JSON(err.StatusCode(), response)
}
