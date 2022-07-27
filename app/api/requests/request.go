package requests

import (
	"fmt"
	"im-services/pkg/model"
)

func IsTableFliedExits(filed string, value string, table string) bool {

	var count int64
	model.DB.Table(table).Where(fmt.Sprintf("%s=?", filed), value).Count(&count)

	if count > 0 {
		return true
	}
	return false
}
