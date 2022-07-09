/**
  @author:panliang
  @data:2022/6/5
  @note
**/
package requests

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"im-services/pkg/model"
)

func ValidateInit() *validator.Validate {
	Validate := validator.New()
	err := Validate.RegisterValidation("checkTableFiled", CheckTableFiled)
	if err != nil {
		fmt.Println("注册失败！")
	}
	return Validate
}

func CheckTableFiled(f validator.FieldLevel) bool { // FieldLevel contains all the information and helper functions to validate a field
	fmt.Println("字段参数", f.Field(), f.StructFieldName())

	//model.DB.Table(data[0]).Where(fmt.Sprintf("%s=?", f.Field()))

	return false
}

// 判断字段是否在表中存在
func IsTableFliedExits(filed string, value string, table string) bool {

	var count int64
	model.DB.Table(table).Where(fmt.Sprintf("%s=?", filed), value).Count(&count)

	if count > 0 {
		return true
	}
	return false
}
