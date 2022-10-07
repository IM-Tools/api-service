package services

import (
	"encoding/json"
	"im-services/internal/api/requests"
	"im-services/internal/helpers"
	"im-services/internal/models/group_message"
	"im-services/internal/models/im_group_users"
	AppClient "im-services/internal/service/client"
	"im-services/internal/service/queue/nsq_queue"
	"im-services/pkg/date"
	"im-services/pkg/model"
	"unsafe"
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

func (s ImMessageService) SendFriendActionMessage(msg AppClient.CreateFriendMessage) {
	AppClient.ImManager.SendFriendActionMessage(msg)
}

func (s ImMessageService) SendPrivateMessage(message requests.PrivateMessageRequest) (bool, string) {
	isOk, respMessage := AppClient.ImManager.SendPrivateMessage(message)
	return isOk, respMessage
}

func (s ImMessageService) SendGroupMessage(message requests.PrivateMessageRequest) bool {
	var users []Users

	model.DB.Model(&im_group_users.ImGroupUsers{}).Where("group_id=?", message.ToID).Select([]string{"user_id"}).Find(&users)
	var groupMesage group_message.ImGroupMessages

	Len := unsafe.Sizeof(&message)
	MessageBytes := &SliceMock{
		addr: uintptr(unsafe.Pointer(&message)),
		cap:  int(Len),
		len:  int(Len),
	}
	data := *(*[]byte)(unsafe.Pointer(MessageBytes))

	for _, val := range users {
		isOk := AppClient.ImManager.SendMessageToSpecifiedClient(data, val.UserId)
		if !isOk {
			//没有就丢入消息队列
			nsq_queue.ProducerQueue.SendGroupMessage(data)
		}

	}
	groupMesage.Message = message.Message
	groupMesage.SendTime = date.TimeUnix()
	groupMesage.MessageId = int(message.MsgClientId)
	groupMesage.ClientMessageId = int(message.MsgClientId)
	groupMesage.FormId = message.FormID
	groupMesage.GroupId = message.ToID

	model.DB.Model(&im_group_users.ImGroupUsers{}).Create(&groupMesage)

	return true

}

func (s ImMessageService) SendVideoMessage(message requests.VideoMessageRequest) bool {
	msg, _ := json.Marshal(message)
	isOk := AppClient.ImManager.SendMessageToSpecifiedClient(msg, helpers.Int64ToString(message.ToID))
	return isOk
}

type Users struct {
	UserId string `json:"user_id"`
}
