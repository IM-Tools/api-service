/**
  @author:panliang
  @data:2022/6/6
  @note
**/
package dao

import (
	"github.com/valyala/fastjson"
	"im-services/app/models/offline_message"
	"im-services/pkg/model"
)

type DataInterface interface {
	PrivateOfflineMessageSave(string) //私聊离线消息
	GroupOfflineMessageSave()         //群组离线消息
}

var (
	OfflineMessageSave *OfflineMessageDao
)

func New() *OfflineMessageDao {
	OfflineMessageSave = new(OfflineMessageDao)
	return OfflineMessageSave
}

type OfflineMessageDao struct {
}

// 消息入库
func (offline *OfflineMessageDao) PrivateOfflineMessageSave(msg string) {
	var p fastjson.Parser
	v, _ := p.Parse(msg)
	ReceiveId := v.GetInt64("to_id")
	msgId := v.GetInt64("msg_id")
	sendTime := v.GetInt("send_time")

	model.DB.Table("im_offline_messages").Create(&offline_message.ImOfflineMessages{
		MsgId:     msgId,
		Status:    0,
		SendTime:  sendTime,
		ReceiveId: ReceiveId,
		Message:   msg,
	})
}

func (offline *OfflineMessageDao) GroupOfflineMessageSave() {

}
