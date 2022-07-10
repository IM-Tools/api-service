/**
  @author:panliang
  @data:2022/5/27
  @note
**/
package message

import (
	"encoding/json"
	"fmt"
	"im-services/app/helpers"
	"im-services/app/service/dispatch"
	"im-services/pkg/date"
)

type MessageClient struct {
	ReceiveId   string  `json:"receive_id"`
	ChannelType int     `json:"channel_type"`
	Msg         Message `json:"msg"`
}

// ack机制
type AckMsg struct {
	Ack         int   `json:"ack"`           // 1.消息已经投递到服务器了
	MsgId       int64 `json:"msg_id"`        //服务器生成的消息id
	MsgClientId int64 `json:"msg_client_id"` //客户端生成的消息id
}

type CreateFriendMessage struct {
	MsgCode     int    `json:"msg_code"`    // 定义的消息code
	ID          int64  `json:"id"`          // 定义的消息code
	FormId      int64  `json:"form_id"`     // 发消息的人
	Status      int    `json:"status"`      // 发消息的人
	CreatedAt   string `json:"created_at"`  // 发消息的人
	ToID        int64  `json:"to_id"`       // 接收消息人的id
	Information string `json:"information"` //内容
	Users       Users  `json:"users"`       //请求人
}

type Users struct {
	Name   string `json:"name"`
	ID     int64  `json:"id"`
	Avatar string `json:"avatar"`
}

// 用户发送的消息数据
type Message struct {
	MsgId       int64       `json:"msg_id"`        // 服务端消息唯一id
	MsgClientId int64       `json:"msg_client_id"` // 客户端消息唯一id
	MsgCode     int         `json:"msg_code"`      // 定义的消息code
	FormID      int64       `json:"form_id"`       // 发消息的人
	ToID        int64       `json:"to_id"`         // 接收消息人的id
	Uid         string      `json:"uid"`           // uid
	ToUid       string      `json:"to_uid"`        // to uid
	MsgType     int         `json:"msg_type"`      // 消息类型 例如 1.文本 2.语音 3.文件
	ChannelType int         `json:"channel_type"`  // 频道类型 1.私聊 2.频道 3.广播
	Message     string      `json:"message"`       // 消息
	SendTime    string      `json:"send_time"`     // 消息发送时间
	Data        interface{} `json:"data"`          // 自定义携带的数据
}

type BroadcastMessages struct {
}

type MessageInterface interface {
	ValidationMsg(msg []byte) (error, string)
}

type MessageHandler struct {
}

// 验证消息是否正确 此处可以做消息拦截
func (m *MessageHandler) ValidationMsg(msg []byte) (error, []byte, []byte, int) {

	var errs error

	if len(msg) == 0 {
		return errs, []byte(`{"code":500,"message":"请勿发送空消息"}`), []byte(``), 0
	}

	var userMsg Message

	err := json.Unmarshal(msg, &userMsg)

	if err != nil {
		return err, []byte(`{"code":500,"message":"用户消息解析异常"}`), []byte(``), 0
	}
	userMsg.MsgId = date.TimeUnixNano()
	userMsg.SendTime = date.NewDate()

	var ackMsg AckMsg

	ackMsg.MsgId = userMsg.MsgId
	ackMsg.MsgClientId = userMsg.MsgClientId
	ackMsg.Ack = 1

	fmt.Println(userMsg)
	msgByte, _ := json.Marshal(&MessageClient{
		ReceiveId:   helpers.Int64ToString(userMsg.ToID),
		ChannelType: userMsg.ChannelType,
		Msg:         userMsg,
	})

	fmt.Println(string(msgByte))

	var dService dispatch.DispatchService

	ok, node := dService.IsDispatchNode(helpers.Int64ToString(userMsg.ToID))
	if !ok && node != "" {
		// 将消息分发到指定的客户端
	}
	ackMsgByte, _ := json.Marshal(ackMsg)

	return nil, msgByte, ackMsgByte, userMsg.ChannelType

}
