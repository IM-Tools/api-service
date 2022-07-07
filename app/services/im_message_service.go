/**
  @author:panliang
  @data:2022/7/3
  @note
**/
package services

import (
	"encoding/json"
	"im-services/service/client"
	"im-services/service/message"
)

type ImMessageService struct {
}

type ImMessageServiceInterface interface {

	// 发送-好友申请或者拒绝❌好友请求

	SendFriendActionMessage(msg message.CreateFriendMessage)
}

func (s ImMessageService) SendFriendActionMessage(msg message.CreateFriendMessage) {
	jsonByte, _ := json.Marshal(msg)
	client.ImManager.BroadcastChannel <- jsonByte
}
