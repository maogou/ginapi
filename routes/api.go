package routes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/maogou/ginapi/app/handler/api/v1"
)

func NewRouter() *gin.Engine  {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	article := v1.Article{}
	tag := v1.Tag{}

	v1 := router.Group("/api/v1")
	{
		// 创建标签
		v1.POST("/tags",tag.Create)
		// 删除指定标签
		v1.DELETE("/tags/:id",tag.Delete)
		// 更新指定标签
		v1.PUT("/tags/:id",tag.Update)
		// 获取标签列表
		v1.GET("/tags",tag.List)

		// 创建文章
		v1.POST("/articles",article.Create)
		// 删除指定文章
		v1.DELETE("/articles/:id",article.Delete)
		// 更新指定文章
		v1.PUT("/articles/:id",article.Update)
		// 获取指定文章
		v1.GET("/articles/:id",article.Get)
		// 获取文章列表
		v1.GET("/articles",article.List)
	}

	return router
}
