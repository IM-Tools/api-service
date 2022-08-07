/**
  @author:panliang
  @data:2022/8/6
  @note
**/
package handler

import "im-services/internal/service/client"

type SendMessageHandler struct {
}

func (*SendMessageHandler) SendMessageToSpecifiedClient(message []byte, toId string) bool {
	data, ok := client.ImManager.ImClientMap[toId]
	if ok {
		data.Send <- message
		return true
	}
	return false
}
