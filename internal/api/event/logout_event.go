package event

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"im-services/internal/enum"
	"im-services/internal/helpers"
	AppClient "im-services/internal/service/client"
)

type EventHandle struct {
}

func (event *EventHandle) LogoutEvent(id string, ip string) {
	client, ok := AppClient.ImManager.ImClientMap[id]
	if ok {
		var message LogoutMessage
		message.MsgCode = enum.WsLoginOut
		message.Message = "您的账号在异地登录了!"
		err, info := helpers.GetIpInfo(ip)
		if err == nil {
			message.City = info.City
			message.Pro = info.Pro
			message.Ip = info.IP
		}
		msg, _ := json.Marshal(message)
		client.Socket.WriteMessage(websocket.TextMessage, msg)
		// 移除当前客户端
		AppClient.ImManager.DelClient(client)
	}
}

type LogoutMessage struct {
	MsgCode int    `json:"msg_code"`
	Message string `json:"message"`
	Ip      string `json:"ip"`
	Pro     string `json:"pro"`
	City    string `json:"city"`
}
