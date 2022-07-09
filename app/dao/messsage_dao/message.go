/**
  @author:panliang
  @data:2022/7/8
  @note
**/
package messsage_dao

import (
	"im-services/app/models/im_messages"
	"im-services/pkg/model"
)

type MessageDao struct {
}

func (m MessageDao) CreateMessage() {
	var message im_messages.ImMessages
	
	model.DB.Save(&message)
}