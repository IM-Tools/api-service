package dao

import (
	"im-services/internal/models/offline_message"
	"im-services/pkg/date"
	"im-services/pkg/model"

	"github.com/valyala/fastjson"
)

type DataInterface interface {
	PrivateOfflineMessageSave(string)                                    //私聊离线消息
	GroupOfflineMessageSave()                                            //群组离线消息
	PullPrivateOfflineMessage(int64) []offline_message.ImOfflineMessages //拉取离线消息
	UpdatePrivateOfflineMessageStatus(int64)                             //更新离线消息
}

var (
	OfflineMessage OfflineMessageDao
)

type OfflineMessageDao struct {
}

// todo 优化合并成多条进行插入数据库 从主
func (offline *OfflineMessageDao) PrivateOfflineMessageSave(msg string) {
	var p fastjson.Parser
	v, _ := p.Parse(msg)
	ReceiveId := v.GetInt64("to_id")
	model.DB.Table("im_offline_messages").Create(&offline_message.ImOfflineMessages{
		Status:    0,
		SendTime:  int(date.TimeUnix()),
		ReceiveId: ReceiveId,
		Message:   msg,
	})
}

// todo 优化合并成多条进行插入数据库 主从
func (offline *OfflineMessageDao) GroupOfflineMessageSave(msg string) {
	var p fastjson.Parser
	v, _ := p.Parse(msg)
	userId := v.GetInt("user_id")
	sendTime := date.TimeUnix()

	model.DB.Table("im_group_offline_messages").Create(&offline_message.ImGroupOfflineMessages{
		SendTime:  int(sendTime),
		Message:   msg,
		Status:    0,
		ReceiveId: userId,
	})

}
