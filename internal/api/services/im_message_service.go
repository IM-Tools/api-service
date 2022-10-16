package services

import (
	"encoding/json"
	"im-services/internal/api/requests"
	"im-services/internal/enum"
	"im-services/internal/helpers"
	"im-services/internal/models/group_message"
	"im-services/internal/models/im_group_users"
	AppClient "im-services/internal/service/client"
	"im-services/internal/service/queue/nsq_queue"
	"im-services/pkg/date"
	"im-services/pkg/model"
)

type ImMessageService struct {
}

type ImMessageServiceInterface interface {

	// 发送-好友申请或者拒绝❌好友请求

	SendFriendActionMessage(msg AppClient.CreateFriendMessage)

	// 发送私聊消息
	SendPrivateMessage(msg requests.PrivateMessageRequest) (bool, string)

	// 发送群聊消息
	SendGroupMessage(msg requests.PrivateMessageRequest) (bool, string)

	// 发送视频请求
	SendVideoMessage(msg requests.VideoMessageRequest) bool
}

type SliceMock struct {
	addr uintptr
	len  int
	cap  int
}

func (*ImMessageService) SendFriendActionMessage(msg AppClient.CreateFriendMessage) {
	AppClient.ImManager.SendFriendActionMessage(msg)
}

func (*ImMessageService) SendPrivateMessage(message requests.PrivateMessageRequest) (bool, string) {
	isOk, respMessage := AppClient.ImManager.SendPrivateMessage(message)
	return isOk, respMessage
}

func (*ImMessageService) SendChatMessage(message requests.PrivateMessageRequest) (bool, string) {
	message.ToID = message.FormID
	message.FormID = 1
	message.Message = GetMessage(message.Message)
	messageDao.CreateMessage(message)
	isOk, respMessage := AppClient.ImManager.SendPrivateMessage(message)
	return isOk, respMessage
}
func (*ImMessageService) SendGroupMessage(message requests.PrivateMessageRequest) bool {
	var users []Users
	model.DB.Model(&im_group_users.ImGroupUsers{}).Where("group_id=?", message.ToID).Select([]string{"user_id"}).Find(&users)
	var groupMessage group_message.ImGroupMessages

	msg, _ := json.Marshal(message)

	for _, val := range users {
		isOk := AppClient.ImManager.SendMessageToSpecifiedClient(msg, val.UserId)
		if !isOk {
			//没有就丢入消息队列
			nsq_queue.ProducerQueue.SendGroupMessage(msg)
		}

	}
	groupMessage.Message = message.Message
	groupMessage.SendTime = date.TimeUnix()
	groupMessage.MessageId = message.MsgId
	groupMessage.ClientMessageId = message.MsgClientId
	groupMessage.FormId = message.FormID
	groupMessage.GroupId = message.ToID

	model.DB.Model(&group_message.ImGroupMessages{}).Create(&groupMessage)

	return true

}

func (*ImMessageService) SendVideoMessage(message requests.VideoMessageRequest) bool {
	msg, _ := json.Marshal(message)
	isOk := AppClient.ImManager.SendMessageToSpecifiedClient(msg, helpers.Int64ToString(message.ToID))
	return isOk
}

type Users struct {
	UserId string `json:"user_id"`
}

type Sessions struct {
	Id          int64    `gorm:"column:id;primaryKey" json:"id"` //会话表
	FormId      int64    `gorm:"column:form_id" json:"form_id"`
	ToId        int64    `gorm:"column:to_id" json:"to_id"`
	GroupId     int64    `gorm:"column:group_id" json:"group_id"` // 群组id
	CreatedAt   string   `gorm:"column:created_at" json:"created_at"`
	TopStatus   int      `gorm:"column:top_status" json:"top_status"` //0.否 1.是
	TopTime     string   `gorm:"column:top_time" json:"top_time"`
	Note        string   `gorm:"column:note" json:"note"`                 //备注
	ChannelType int      `gorm:"column:channel_type" json:"channel_type"` //0.单聊 1.群聊
	Name        string   `gorm:"column:name" json:"name"`                 //会话名称
	Avatar      string   `gorm:"column:avatar" json:"avatar"`             //会话头像
	Status      int      `gorm:"column:status" json:"status"`             //会话状态 0.正常 1.禁用
	Groups      ImGroups `gorm:"foreignKey:ID;references:GroupId"`
}

type ImSessionsMessage struct {
	MsgCode  int      `json:"msg_code"` // 定义的消息code
	Sessions Sessions `json:"sessions"` // 会话内容
}

type ImGroups struct {
	ID        int64  `gorm:"column:id" json:"id"`                 //群聊id
	UserId    int64  `gorm:"column:user_id" json:"user_id"`       //创建者
	Name      string `gorm:"column:name" json:"name"`             //群聊名称
	CreatedAt string `gorm:"column:created_at" json:"created_at"` //添加时间
	Info      string `gorm:"column:info" json:"info"`             //群聊描述
	Avatar    string `gorm:"column:avatar" json:"avatar"`         //群聊头像
	IsPwd     int8   `gorm:"column:is_pwd" json:"is_pwd"`         //是否加密 0 否 1 是
	Hot       int    `gorm:"column:hot" json:"hot"`               //热度
}

// 会话消息投递
func (*ImMessageService) SendGroupSessionMessage(userIds []string, groupId int64) {
	var message ImSessionsMessage
	message.MsgCode = enum.WsSession
	model.DB.Table("im_sessions").Where("group_id=?", groupId).Preload("Groups").Find(&message.Sessions)
	for _, id := range userIds {
		message.Sessions.FormId = helpers.StringToInt64(id)
		msg, _ := json.Marshal(message)

		data, ok := AppClient.ImManager.ImClientMap[id]
		if ok {
			data.Send <- msg
		}
	}
}
