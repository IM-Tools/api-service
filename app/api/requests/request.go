/**
  @author:panliang
  @data:2022/6/5
  @note
**/
package requests

import (
	"fmt"
	"im-services/pkg/model"
)

// 判断字段是否在表中存在
func IsTableFliedExits(filed string, value string, table string) bool {

	var count int64
	model.DB.Table(table).Where(fmt.Sprintf("%s=?", filed), value).Count(&count)

	if count > 0 {
		return true
	}
	return false
}
