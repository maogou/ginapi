package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/maogou/ginapi/pkg/app"
	"github.com/maogou/ginapi/pkg/errcode"
)

//token中间件
func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)

		if queryToken, exist := c.GetQuery("token"); exist {
			token = queryToken
		} else {
			token = c.GetHeader("token")
		}

		if token == "" {
			ecode = errcode.TokenInvalidParams
		} else {
			_, err := app.ParseToken(token)

			if err != nil {
				//断言错误
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			}
		}

		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}

		//执行下一个中间件
		c.Next()
	}
}
