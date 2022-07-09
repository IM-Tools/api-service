/**
  @author:panliang
  @data:2022/7/9
  @note
**/
package middleware

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	Uni      *ut.UniversalTranslator
	Validate *validator.Validate
	//local    = config.Conf.Server.Lang
	local = "zh"
)

func ValidateLocal() gin.HandlerFunc {
	return func(cxt *gin.Context) {
		tans, _ := Uni.GetTranslator(local)
		zh_translations.RegisterDefaultTranslations(Validate, tans)
		cxt.Next()
	}
}
