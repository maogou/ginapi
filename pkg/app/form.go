package app

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

//实现了error接口-golang规定只要实现了接口的方法就实现了该接口
func (v ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

//实现了error接口
func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

//绑定并验证翻译错误信息
func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors

	err := c.ShouldBind(v)

	if err != nil {
		v := c.Value("trans")
		trans, _ := v.(ut.Translator)
		verres, ok := err.(validator.ValidationErrors)
		if !ok {
			return false, nil
		}

		for key, value := range verres.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}

		return false, errs
	}

	return true, nil

}
