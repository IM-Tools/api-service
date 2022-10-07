package dao

import (
	"im-services/internal/models/group_message"
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

// PrivateOfflineMessageSave 消息入库
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

func (offline *OfflineMessageDao) GroupOfflineMessageSave(msg string) {
	var p fastjson.Parser
	v, _ := p.Parse(msg)
	ReceiveId := v.GetInt64("to_id")
	FormId := v.GetInt64("form_id")
	sendTime := v.GetInt64("send_time")

	model.DB.Table("im_offline_messages").Create(&group_message.ImGroupMessages{
		SendTime: sendTime,
		GroupId:  ReceiveId,
		FormId:   FormId,
		Message:  msg,
	})

}
