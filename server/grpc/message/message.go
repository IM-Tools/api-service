/**
  @author:panliang
  @data:2022/7/30
  @note
**/
package grpcMessage

import (
	"context"
	"fmt"
	"im-services/internal/api/requests"
	"im-services/internal/enum"
	"im-services/pkg/date"
	"im-services/pkg/logger"
)

// ImGrpcMessage 实现 ImMessageServer 接口
type ImGrpcMessage struct {
}

func (ImGrpcMessage) mustEmbedUnimplementedImMessageServer() {}

// ReceivesGrpcPrivateMessage 接收消息
func (ImGrpcMessage) SendMessageHandler(c context.Context, request *SendMessageRequest) (*SendMessageResponse, error) {
	logger.Logger.Error(request.Message)
	params := requests.PrivateMessageRequest{
		MsgId:       date.TimeUnixNano(),
		MsgCode:     enum.WsChantMessage,
		MsgClientId: request.MsgClientId,
		FormID:      request.FormId,
		ToID:        request.ToId,
		ChannelType: int(request.ChannelType),
		MsgType:     int(request.MsgType),
		Message:     request.Message,
		SendTime:    date.NewDate(),
		Data:        request.Data,
	}

	msgString := GetGrpcPrivateChatMessages(params)

	switch request.ChannelType {
	case 1:
		fmt.Println(msgString)
		//client.ImManager.PrivateChannel <- []byte(msgString)
	case 2:

	}
	return &SendMessageResponse{Code: 200, Message: "Success"}, nil
}
func GetGrpcPrivateChatMessages(message requests.PrivateMessageRequest) string {
	msg := fmt.Sprintf(`{
                "msg_id": %d,
                "msg_client_id": %d,
                "msg_code": %d,
                "form_id": %d,
                "to_id": %d,
                "msg_type": %d,
                "channel_type": %d,
                "message": %s,
                "data": %s
        }`, message.MsgId, message.MsgClientId, message.MsgCode, message.FormID, message.ToID, message.MsgType, message.ChannelType, message.Message, message.Data)

	return msg
}
