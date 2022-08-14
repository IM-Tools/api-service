package services

import (
	"encoding/json"
	"im-services/internal/api/requests"
	"im-services/internal/helpers"
	AppClient "im-services/internal/service/client"
)

type ImMessageService struct {
}

type ImMessageServiceInterface interface {

	// 发送-好友申请或者拒绝❌好友请求

	SendFriendActionMessage(msg AppClient.CreateFriendMessage)

	// 发送私聊消息
	SendPrivateMessage(msg requests.PrivateMessageRequest) (bool, string)

	// 发送视频请求
	SendVideoMessage(msg requests.VideoMessageRequest) bool
}

func (s ImMessageService) SendFriendActionMessage(msg AppClient.CreateFriendMessage) {
	AppClient.ImManager.SendFriendActionMessage(msg)
}

func (s ImMessageService) SendPrivateMessage(message requests.PrivateMessageRequest) (bool, string) {
	isOk, respMessage := AppClient.ImManager.SendPrivateMessage(message)
	return isOk, respMessage
}

func (s ImMessageService) SendVideoMessage(message requests.VideoMessageRequest) bool {
	msg, _ := json.Marshal(message)
	isOk := AppClient.ImManager.SendMessageToSpecifiedClient(msg, helpers.Int64ToString(message.ToID))
	return isOk
}
