/**
  @author:panliang
  @data:2022/7/3
  @note
**/
package dao

import "im-services/pkg/model"

type DataDao struct {
}

func (dao *DataDao) IsUserExits(table string, filed string, value interface{}) bool {

	if model.DB.Table(table).Where(filed+"=?", value).RowsAffected == 0 {
		return false
	}
	return false
}
