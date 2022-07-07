/**
  @author:panliang
  @data:2022/6/5
  @note
**/
package requests

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/thedevsaddam/govalidator"
	"im-services/pkg/model"
	"strings"
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

// 此方法会在初始化时执行
func init() {
	// not_exists:users,email
	govalidator.AddCustomRule("not_exists", func(field string, rule string, message string, value interface{}) error {
		rng := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")

		tableName := rng[0]
		dbFiled := rng[1]
		val := value.(string)

		var count int64
		model.DB.Table(tableName).Where(dbFiled+" = ?", val).Count(&count)

		if count != 0 {

			if message != "" {
				return errors.New(message)
			}

			return fmt.Errorf("%v 已被占用", val)
		}
		return nil
	})

	govalidator.AddCustomRule("exists", func(field string, rule string, message string, value interface{}) error {
		rng := strings.Split(strings.TrimPrefix(rule, "exists:"), ",")

		tableName := rng[0]
		dbFiled := rng[1]
		val := value.(string)

		var count int64
		model.DB.Table(tableName).Where(dbFiled+" = ?", val).Count(&count)

		if count != 0 {

			if message != "" {
				return errors.New(message)
			}
			return nil

		}
		return fmt.Errorf("%v 不存在", val)
	})
}
