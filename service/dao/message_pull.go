/**
  @author:panliang
  @data:2022/6/8
  @note
**/
package dao

import (
	"im-services/app/models/offline_message"
	"im-services/pkg/model"
)

// 拉取离线私聊消息
func (offline *OfflineMessageDao) PullPrivateOfflineMessage(id int64) []offline_message.ImOfflineMessages {

	var list []offline_message.ImOfflineMessages

	// 拉去最近半个月内的消息记录
	//timeStamp := helpers.GetDayTime(-15)

	model.DB.Table("im_offline_messages").
		Where("status=0 and receive_id=? and send_time>", id, 163607878).
		Find(&list)

	return list
}

// 更新消息状态
func (offline *OfflineMessageDao) UpdatePrivateOfflineMessageStatus(id int64) {
	model.DB.Table("im_offline_messages").
		Where("status=0 and receive_id=?", id).
		Updates(map[string]interface{}{"status": 1})
}
