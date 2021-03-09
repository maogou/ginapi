package common

import (
	"github.com/gin-gonic/gin"
	"github.com/maogou/ginapi/pkg/app"
	"github.com/maogou/ginapi/pkg/errcode"
)

type Intercept struct {
}

func NewIntercept() Intercept {
	return Intercept{}
}

//拦截404
func (i Intercept) Intercept404(c *gin.Context) {
	response := app.NewResponse(c)

	notFound := errcode.NotFound
	response.ToErrorResponse(notFound)
	return
}

//拦截405
func (i Intercept) Intercept405(c *gin.Context) {
	response := app.NewResponse(c)

	methodNotAllow := errcode.MethodNotAllow
	response.ToErrorResponse(methodNotAllow)
	return
}
