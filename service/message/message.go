/**
  @author:panliang
  @data:2022/5/27
  @note
**/
package message

import (
	"encoding/json"
)

type MessageClient struct {
	SendId int64
	Msg    *Message
}

type Message struct {
	MsgCode     string      `json:"msg_code"`     // 定义的消息code
	FormID      int64       `json:"form_id"`      // 发消息的人
	ToID        int64       `json:"to_id"`        // 接收消息人的id
	MsgType     int         `json:"msg_type"`     // 消息类型 例如 1.文本 2.语音 3.文件
	ChannelType int         `json:"channel_type"` // 频道类型 1.私聊 2.频道 3.广播
	Data        interface{} // 自定义携带的消息
}

var (
	MsgHandler *MessageHandler
	userMsg    *Message
)

type MessageInterface interface {
	ValidationMsg(msg []byte) (error, string)
}

type MessageHandler struct {
}

func NewMsg() *Message {
	userMsg = new(Message)
	return userMsg
}

func New() *MessageHandler {
	MsgHandler = new(MessageHandler)
	return MsgHandler
}

// 验证消息是否正确 此处可以做消息拦截。
func (m *MessageHandler) ValidationMsg(msg []byte) (error, string) {

	var errs error

	if len(msg) == 0 {
		return errs, "请勿发送空消息"
	}

	userMsg = NewMsg()

	err := json.Unmarshal(msg, &userMsg)

	if err != nil {
		return err, "用户消息解析异常"
	}
	msgByte, _ := json.Marshal(&MessageClient{
		SendId: userMsg.ToID,
		Msg:    userMsg,
	})

	return nil, string(msgByte)

}
