package message

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/valyala/fastjson"
	"im-services/app/api/requests"
	"im-services/app/enum"
	"im-services/app/helpers"
	"im-services/app/service/dispatch"
	"im-services/pkg/date"
)

type MessageClient struct {
	ReceiveId   string  `json:"receive_id"`
	ChannelType int     `json:"channel_type"`
	Msg         Message `json:"msg"`
}

// AckMsg ack消息
type AckMsg struct {
	Ack         int   `json:"ack"`           // 1.消息已经投递到服务器了
	MsgCode     int   `json:"msg_code"`      // 1.消息已经投递到服务器了
	MsgId       int64 `json:"msg_id"`        //服务器生成的消息id
	MsgClientId int64 `json:"msg_client_id"` //客户端生成的消息id
}

// CreateFriendMessage 私聊内容
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

// Message 用户发送的消息
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

// PingMessage 心跳消息
type PingMessage struct {
	MsgCode int    `json:"msg_code"`
	Message string `json:"message"`
}

type BroadcastMessages struct {
}

type MessageInterface interface {
	ValidationMsg(msg []byte) (error, string)
}

type MessageHandler struct {
}

// ValidationMsg 验证消息是否正确 此处可以做消息拦截
func (m *MessageHandler) ValidationMsg(msg []byte) (error, []byte, []byte, int) {

	var errs error

	var p fastjson.Parser
	v, _ := p.Parse(string(msg))

	msgCode, _ := v.Get("msg_code").Int()

	if msgCode == enum.WsPing {
		return nil, []byte(`{"msg_code":1004,"message":"ping"}`), []byte(``), 3
	}

	if len(msg) == 0 {
		return errs, []byte(`{"msg_code":500,"message":"请勿发送空消息"}`), []byte(``), 0
	}

	params := requests.PrivateMessageRequest{
		MsgId:       date.TimeUnixNano(),
		MsgCode:     enum.WsChantMessage,
		MsgClientId: v.GetInt64("msg_client_id"),
		FormID:      v.GetInt64("form_id"),
		ToID:        v.GetInt64("to_id"),
		ChannelType: v.GetInt("channel_type"),
		MsgType:     v.GetInt("msg_type"),
		Message:     v.Get("msg_client_id").String(),
		SendTime:    date.NewDate(),
		Data:        v.Get("data").String(),
	}

	err := validator.New().Struct(params)

	if err != nil {
		return err, []byte(`{"msg_code":500,"message":"用户消息解析异常"}`), []byte(``), 0

	}

	chatMessage := m.GetPrivateChatMessages(params)

	var dService dispatch.DispatchService

	ok, node := dService.IsDispatchNode(helpers.Int64ToString(params.ToID))
	if !ok && node != "" {
		// todo 将消息分发到指定的客户端
	}

	ack := AckMsg{
		MsgId:       params.MsgId,
		MsgClientId: params.MsgClientId,
		Ack:         1,
		MsgCode:     enum.WsAck,
	}

	ackMsg := m.GetAckMessages(ack)

	return nil, []byte(chatMessage), []byte(ackMsg), params.ChannelType

}

// GetPrivateChatMessages 组装私聊消息
func (m *MessageHandler) GetPrivateChatMessages(message requests.PrivateMessageRequest) string {
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

	msgString := fmt.Sprintf(`{
"receive_id":"%d",
"channel_type":%d,
"msg":%s
}`, message.ToID, message.ChannelType, msg)

	return msgString
}

// 获取ack消息
func (m *MessageHandler) GetAckMessages(ack AckMsg) string {
	msg := fmt.Sprintf(`{"ack": %d,"msg_code": %d,"msg_id": %d,"msg_client_id": %d,}`, 1, ack.MsgCode, ack.MsgId, ack.MsgClientId)
	return msg
}
