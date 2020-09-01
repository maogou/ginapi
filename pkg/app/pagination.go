package app

import (
	"github.com/gin-gonic/gin"
	"github.com/maogou/ginapi/global"
	"github.com/maogou/ginapi/pkg/transform"
)

//获取当前的页码
func GetPage(c *gin.Context) int  {
	page := transform.StrTo(c.Query("page")).MustInt()

	if page <= 0 {
		return 1
	}

	return page
}

//获取limit
func GetPageSize(c *gin.Context) int  {
	pageSize := transform.StrTo(c.Query("page_size")).MustInt()

	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}

	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}

	return pageSize
}

//获取offset
func GetPageOffset(page,pageSize int) int  {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}

	return result
}
