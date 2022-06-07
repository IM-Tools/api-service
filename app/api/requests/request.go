/**
  @author:panliang
  @data:2022/6/5
  @note
**/
package requests

import (
	"im-services/pkg/model"
	"regexp"
)

func IsEmail(email string) bool {
	result, _ := regexp.MatchString(`^([\w\.\_\-]{2,10})@(\w{1,}).([a-z]{2,4})$`, email)
	if result {
		return true
	} else {
		return false
	}
}

func IsEmailExits(email string, table string) (bool, string) {
	if !IsEmail(email) {
		return false, "不是一个正确的邮箱"
	}
	var count int64
	model.DB.Table(table).Where("email=?", email).Count(&count)

	if count > 0 {
		return false, "邮箱已经存在了"
	}
	return true, ""
}
