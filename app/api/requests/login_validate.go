/**
  @author:panliang
  @data:2022/6/3
  @note
**/
package requests

import "github.com/thedevsaddam/govalidator"

type LoginForm struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required" `
}

func ValidateLogin(data LoginForm) map[string][]string {
	rules := govalidator.MapData{
		"email":    []string{"required", "min:2", "max:20", "email"},
		"password": []string{"required", "min:6", "max:20"},
	}

	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 2",
			"max:Email 长度需小于 20",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
		"password": []string{
			"required:密码为必填项",
			"min:长度需大于 6",
		},
	}
	opts := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}

	errs := govalidator.New(opts).ValidateStruct()

	return errs
}
