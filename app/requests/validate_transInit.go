/**
  @author:panliang
  @data:2022/5/21
  @note
**/
package requests

import (
	"fmt"
	"github.com/go-playground/locales/ug_CN"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"im-services/config"
	"im-services/pkg/logger"

	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	uni      *ut.UniversalTranslator
	Validate *validator.Validate
)

// 注册翻译器
func ValidateTransInit() *validator.Validate {
	Validate = validator.New()
	locale := config.Conf.Server.Lang
	switch locale {
	case "zh":
		fmt.Println(locale)
		uni = ut.New(zh.New())
		break
	default:
		uni = ut.New(ug_CN.New())
		break
	}
	trans, _ := uni.GetTranslator(locale)
	err := zhTranslations.RegisterDefaultTranslations(Validate, trans)
	if err != nil {
		logger.Logger.Info(err.Error())
	}
	return Validate
}
