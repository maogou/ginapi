package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/maogou/ginapi/global"
	"github.com/maogou/ginapi/pkg/app"
	"github.com/maogou/ginapi/pkg/errcode"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Errorf(c, "panic recover err: %#v", err)
				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()

		c.Next()
	}

}
