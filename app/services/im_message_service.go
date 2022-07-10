/**
  @author:panliang
  @data:2022/7/3
  @note
**/
package services

import (
	"encoding/json"
	"fmt"
	"im-services/app/api/requests"
	"im-services/app/api/service/client"
	"im-services/app/api/service/message"
)

type ImMessageService struct {
}

type ImMessageServiceInterface interface {

	// 发送-好友申请或者拒绝❌好友请求

	SendFriendActionMessage(msg message.CreateFriendMessage)

	// 发送私聊消息
	SendPrivateMessage(msg requests.PrivateMessageRequest)
}

func (s ImMessageService) SendFriendActionMessage(msg message.CreateFriendMessage) {
	jsonByte, _ := json.Marshal(msg)
	client.ImManager.BroadcastChannel <- jsonByte
}

func (s ImMessageService) SendPrivateMessage(msg requests.PrivateMessageRequest) (bool, string) {

	var handler message.MessageHandler
	jsonByte, _ := json.Marshal(msg)

	errs, msgString, _, channel := handler.ValidationMsg(jsonByte)

	if errs != nil {
		return false, "消息解析失败"
	} else {
		// 将消费分发到不同的队列
		switch channel {
		case 1:
			client.ImManager.PrivateChannel <- msgString
		case 2:

		default:
		}
	}

	return true, "Success"

}

func getKey() {
	for key, _ := range client.ImManager.ImClientMap {
		fmt.Println(key)
	}
}
