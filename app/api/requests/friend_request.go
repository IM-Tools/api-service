/**
  @author:panliang
  @data:2022/7/3
  @note
**/
package requests

import (
	"github.com/thedevsaddam/govalidator"
)

type UpdateFriendRequest struct {
	ID     string `json:"id" validate:"required"`
	Status int    `json:"status" validate:"required,gte=1,lte=2"`
}
type CreateFriendRequest struct {
	ToId        string `json:"to_id" validate:"required"`
	Information string `json:"information" validate:"required"`
}

func ValidateCreateFriend(data CreateFriendRequest) map[string][]string {

	// 自定义验证规则
	rules := govalidator.MapData{
		"to_id": []string{"required", "exists:im_users,id"},
	}

	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"to_id": []string{
			"required:手为必填项，参数名称 ",
		},
	}

	// 配置初始化
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}

	// 开始验证
	return govalidator.New(opts).ValidateStruct()
}
