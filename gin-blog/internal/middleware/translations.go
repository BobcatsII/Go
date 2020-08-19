package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	//locales 多语言包，该库需要配合 universal-translator 使用
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	//universal-translator 通用翻译器
	"github.com/go-playground/universal-translator"
	//下面是 validator 翻译器
	validator "github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		uni := ut.New(en.New(), zh.New(), zh_Hant_TW.New())
		//通过GetHeader获取约定的header参数locale，判别是en还是zh。
		locale := c.GetHeader("locale")
		trans, _ := uni.GetTranslator(locale)
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			switch locale {
			//调用RegisterDefaultTranslations方法，将验证器和对应语言类型的 Translator 注册进来
			//同时将	Translator 存储到全局上下文，以便后续翻译使用。
			case "zh":
				_ = zh_translations.RegisterDefaultTranslations(v, trans)
				break
			case "en":
				_ = en_translations.RegisterDefaultTranslations(v, trans)
				break
			default:
				_ = zh_translations.RegisterDefaultTranslations(v, trans)
				break
			}
			c.Set("trans", trans)
		}

		c.Next()
	}
}
