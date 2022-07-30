/**
  @author:panliang
  @data:2022/7/30
  @note
**/
package server

import (
	"context"
	"im-services/app/api/requests"
	"im-services/app/enum"
	"im-services/app/service/client"
	messageHandler "im-services/app/service/message"
	"im-services/pkg/date"
	grpcMessage "im-services/server/grpc/message"
)

type GrpcMessageServiceInterface interface {
	// SendGpcMessage 消息发送到指定节点
	SendGpcMessage(message []byte, node string)
}

type MessageService struct {
}

// ReceivesGrpcPrivateMessage 接收消息
func (ps *MessageService) ReceivesGrpcPrivateMessage(ctx context.Context, request *grpcMessage.SendMessageRequest) (*grpcMessage.SendMessageResponse, error) {

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

	var handler messageHandler.MessageHandler

	msgString := handler.GetPrivateChatMessages(params)

	switch request.ChannelType {
	case 1:
		client.ImManager.PrivateChannel <- []byte(msgString)
	case 2:

	}
	return &grpcMessage.SendMessageResponse{Code: 200}, nil
}
