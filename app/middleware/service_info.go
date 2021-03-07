package middleware

import "github.com/gin-gonic/gin"

//服务信息-微服务可能有多个服务
func ServiceInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_name", "gin-api")
		c.Set("app_version", "1.0.0")
	}
}
