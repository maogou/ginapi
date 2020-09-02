package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

//验证库的多语言翻译中间件
func Translations() gin.HandlerFunc  {
	return func(c *gin.Context) {
		uni := ut.New(en.New(),zh.New())
		//获取header中的locale值
		locale := c.GetHeader("locale")
		trans,_ := uni.GetTranslator(locale)
		v,ok := binding.Validator.Engine().(*validator.Validate)

		if ok {
			switch locale {
			case "en":
				//注册en的验证器和en语言
				_ = en_translations.RegisterDefaultTranslations(v,trans)
			default:
				_ = zh_translations.RegisterDefaultTranslations(v,trans)
			}

			//将trans设置到请求上下文中
			c.Set("trans",trans)
		}

		c.Next()
	}
}
