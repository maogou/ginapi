package routes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/maogou/ginapi/app/handler/api/v1"
	"github.com/maogou/ginapi/app/middleware"

	_ "github.com/maogou/ginapi/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	//使用多语言翻译(验证器)中间件
	router.Use(middleware.Translations())

	//配置swagger文档
	//url := ginSwagger.URL("http://127.0.0.1:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	article := v1.Article{}
	tag := v1.Tag{}
	auth := v1.JwtAuth{}

	//不需要鉴权的路由
	router.POST("/api/v1/token", auth.GetAuth)

	//需要鉴权的路由
	apiV1 := router.Group("/api/v1").Use(middleware.Jwt())
	{
		// 创建标签
		apiV1.POST("/tags", tag.Create)
		// 删除指定标签
		apiV1.DELETE("/tags/:id", tag.Delete)
		// 更新指定标签
		apiV1.PUT("/tags/:id", tag.Update)
		// 获取标签列表
		apiV1.GET("/tags", tag.List)

		// 创建文章
		apiV1.POST("/articles", article.Create)
		// 删除指定文章
		apiV1.DELETE("/articles/:id", article.Delete)
		// 更新指定文章
		apiV1.PUT("/articles/:id", article.Update)
		// 获取指定文章
		apiV1.GET("/articles/:id", article.Get)
		// 获取文章列表
		apiV1.GET("/articles", article.List)
	}

	return router
}
