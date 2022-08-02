/**
  @author:panliang
  @data:2022/7/30
  @note
**/
package client

import (
	"context"
	"fmt"
	"github.com/valyala/fastjson"
	"google.golang.org/grpc"
	"im-services/internal/enum"
	"im-services/pkg/date"
	"im-services/pkg/logger"
	grpcMessage "im-services/server/grpc/message"
)

type GrpcMessageServiceInterface interface {
	// SendGpcMessage 消息发送到指定节点
	SendGpcMessage(message []byte, node string)
}

type GrpcMessageService struct {
}

// 发送grpc消息
func (messageService *GrpcMessageService) SendGpcMessage(message []byte, node string) {
	conn, err := grpc.Dial(node, grpc.WithInsecure())
	if err != nil {

	}
	defer conn.Close()

	ImRpcServiceClient := grpcMessage.NewImMessageClient(conn)

	var p fastjson.Parser
	v, _ := p.Parse(string(message))

	fmt.Println(v.Get("message").String())
	params := &grpcMessage.SendMessageRequest{
		MsgId:       date.TimeUnixNano(),
		MsgClientId: date.TimeUnix(),
		MsgCode:     enum.WsChantMessage,
		FormId:      v.GetInt64("form_id"),
		ToId:        v.GetInt64("to_id"),
		MsgType:     v.GetInt64("msg_type"),
		ChannelType: v.GetInt64("channel_type"),
		Message:     v.Get("message").String(),
		SendTime:    date.TimeUnixNano(),
		Data:        v.Get("data").String(),
	}

	fmt.Println(params)
	resp, err := ImRpcServiceClient.
		SendMessageHandler(context.Background(), params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	logger.Logger.Error(resp.Message)
	return
}
