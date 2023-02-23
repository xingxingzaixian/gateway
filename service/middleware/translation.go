package middleware

import (
	"gateway/public"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"regexp"
)

// 设置Translation
func TranslationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//参照：https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
		//设置支持语言
		en := en.New()
		zh := zh.New()

		//设置国际化翻译器
		uni := ut.New(zh, zh, en)
		val := validator.New()

		//根据参数取翻译器实例
		locale := c.DefaultQuery("locale", "zh")
		trans, _ := uni.GetTranslator(locale)

		//翻译器注册到validator
		switch locale {
		case "en":
			en_translations.RegisterDefaultTranslations(val, trans)
			val.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fld.Tag.Get("en_comment")
			})
			break
		default:
			zh_translations.RegisterDefaultTranslations(val, trans)
			val.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fld.Tag.Get("comment")
			})

			//自定义验证方法
			//https://github.com/go-playground/validator/blob/v9/_examples/custom-validation/main.go
			val.RegisterValidation("valid_service_name", func(fl validator.FieldLevel) bool {
				matched, _ := regexp.Match(`^[a-zA-Z0-9_]{6,128}$`, []byte(fl.Field().String()))
				return matched
			})

			//自定义翻译器
			//https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
			val.RegisterTranslation("valid_service_name", trans, func(ut ut.Translator) error {
				return ut.Add("valid_service_name", "{0} 不符合输入格式", true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("valid_service_name", fe.Field())
				return t
			})
		}

		c.Set(public.TranslatorKey, trans)
		c.Set(public.ValidatorKey, val)
		c.Next()
	}
}
