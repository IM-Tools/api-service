/**
  @author:panliang
  @data:2022/7/30
  @note
**/
package grpcMessage

import (
	"context"
	"im-services/app/api/requests"
	"im-services/app/enum"
	"im-services/app/service/client"
	messageHandler "im-services/app/service/message"
	"im-services/pkg/date"
	"im-services/pkg/logger"
)

type ImGrpcMessage struct {
}

// ReceivesGrpcPrivateMessage 接收消息
func (ps *ImGrpcMessage) SendMessageHandler(ctx context.Context, request *SendMessageRequest) (*SendMessageResponse, error) {

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

	logger.Logger.Error(params.Message)
	var handler messageHandler.MessageHandler

	msgString := handler.GetPrivateChatMessages(params)

	switch request.ChannelType {
	case 1:
		client.ImManager.PrivateChannel <- []byte(msgString)
	case 2:
		client.ImManager.GroupChannel <- []byte(msgString)
	}
	return &SendMessageResponse{Code: 200, Message: "Success"}, nil
}
