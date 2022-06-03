/**
  @author:panliang
  @data:2022/5/21
  @note
**/
package requests

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// 数据翻译
func validateTransInit(validate *validator.Validate) ut.Translator {

	uni := ut.New(zh.New())
	// 获取翻译器
	trans, _ := uni.GetTranslator("zh")
	// 注册翻译器
	err := zhTranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println(err)
	}
	return trans
}
