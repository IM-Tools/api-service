package services

import (
	"encoding/json"
	"fmt"
	"im-services/app/api/requests"
	"im-services/app/service/client"
	messageHandler "im-services/app/service/message"
)

type ImMessageService struct {
}

type ImMessageServiceInterface interface {

	// 发送-好友申请或者拒绝❌好友请求

	SendFriendActionMessage(msg messageHandler.CreateFriendMessage)

	// 发送私聊消息
	SendPrivateMessage(msg requests.PrivateMessageRequest) (bool, string)
}

func (s ImMessageService) SendFriendActionMessage(msg messageHandler.CreateFriendMessage) {
	jsonByte, _ := json.Marshal(msg)
	client.ImManager.BroadcastChannel <- jsonByte
}

func (s ImMessageService) SendPrivateMessage(message requests.PrivateMessageRequest) (bool, string) {

	var handler messageHandler.MessageHandler

	msgString := handler.GetPrivateChatMessages(message)

	// 将消费分发到不同的队列
	switch message.ChannelType {
	case 1:
		client.ImManager.PrivateChannel <- []byte(msgString)
	case 2:
		client.ImManager.GroupChannel <- []byte(msgString)
	default:
	}

	return true, "Success"

}

func getKey() {
	for key, _ := range client.ImManager.ImClientMap {
		fmt.Println(key)
	}
}
