package messsage_dao

import (
	"im-services/internal/api/requests"
	"im-services/internal/helpers"
	"im-services/internal/models/im_messages"
	"im-services/pkg/model"
)

type MessageDao struct {
}

// 私聊消息入库
func (*MessageDao) CreateMessage(params requests.PrivateMessageRequest) {
	var message im_messages.ImMessages
	message.Msg = params.Message
	message.FormId = params.FormID
	message.ToId = params.ToID
	message.CreatedAt = params.SendTime
	message.IsRead = 0
	message.MsgType = params.MsgType
	message.Status = 1
	message.Data = helpers.InterfaceToString(params.Data)
	model.DB.Save(&message)
}
