package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/maogou/ginapi/app/service"
	"github.com/maogou/ginapi/global"
	"github.com/maogou/ginapi/pkg/app"
	"github.com/maogou/ginapi/pkg/errcode"
)

type JwtAuth struct {
}

func NewJwtAuth() JwtAuth {
	return JwtAuth{}
}

func (jwtAuth JwtAuth) GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)

	if valid == false {
		global.Logger.Errorf(c, "app.BindAndValid errs: %#v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetail(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)

	if err != nil {
		global.Logger.Errorf(c, "app.GenerateToken err: %#v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, expireTime, err := app.GenerateToken(param.AppKey, param.AppSecret)

	if err != nil {
		global.Logger.Errorf(c, "app.GenerateToken err: %#v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	data := map[string]interface{}{"token": token, "expire": expireTime}

	response.ToResponse(data)
	return

}
