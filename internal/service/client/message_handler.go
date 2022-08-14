package client

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/valyala/fastjson"
	"im-services/internal/api/requests"
	"im-services/internal/config"
	"im-services/internal/enum"
	"im-services/internal/helpers"
	"im-services/internal/service/cache/firend_cache"
	"im-services/internal/service/dao"
	"im-services/internal/service/dispatch"
	"im-services/internal/service/queue/nsq_queue"
	GrpcClient "im-services/server/client"
)

func (manager *AppImClientManager) LaunchPrivateMessage(message []byte) {

	receiveId, userMsg := GetReceiveIdAndUserMsg(message)

	if client, ok := manager.ImClientMap[receiveId]; ok {
		client.Send <- []byte(userMsg)
	} else {
		nsq_queue.ProducerQueue.SendMessage([]byte(userMsg))
	}

}

func (manager *AppImClientManager) LaunchBroadcastMessage(message []byte) {

	var p fastjson.Parser
	v, _ := p.Parse(string(message))

	msgCode, _ := v.Get("msg_code").Int()

	var ReceiveId string
	if msgCode == enum.WsCreate {
		ReceiveId = v.Get("to_id").String()
	} else {
		ReceiveId = v.Get("form_id").String()
	}

	if client, ok := manager.ImClientMap[ReceiveId]; ok {
		client.Send <- message
	}
}

func (manager *AppImClientManager) LaunchGroupMessage(message []byte) {

	receiveId, userMsg := GetReceiveIdAndUserMsg(message)

	if client, ok := manager.ImClientMap[receiveId]; ok {
		client.Send <- []byte(userMsg)
	} else {
		nsq_queue.ProducerQueue.SendMessage([]byte(userMsg))

	}
}

func (manager *AppImClientManager) ConsumingOfflineMessages(client *ImClient) {
	// 读取离线消息
	list := dao.OfflineMessage.PullPrivateOfflineMessage(client.ID)
	for _, value := range list {
		_ = client.Socket.WriteMessage(websocket.TextMessage, []byte(value.Message))
	}
	// 更新离线消息状态
	if len(list) > 0 {
		dao.OfflineMessage.UpdatePrivateOfflineMessageStatus(client.ID)
	}
}

func (manager *AppImClientManager) RadioUserOnlineStatus(client *ImClient) {

	data, err := firend_cache.FriendCache.Get(client.ID)
	if err != nil {

	}
	for _, val := range data {
		if friendClient, ok := manager.ImClientMap[val.Uid]; ok {
			_ = friendClient.Socket.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(`{"code":200,"message":"用户上线了"',"fo_id":%d}`, int(val.ToId))))
		}
	}
}

// GetReceiveIdAndUserMsg 拿消息投递id
func GetReceiveIdAndUserMsg(msg []byte) (string, string) {
	var p fastjson.Parser
	v, _ := p.Parse(string(msg))
	return fastjson.GetString(msg, "receive_id"), v.GetObject("msg").String()
}

func (manager *AppImClientManager) SendMessageToSpecifiedClient(message []byte, toId string) bool {
	data, ok := manager.ImClientMap[toId]
	if ok {
		data.Send <- message
		return true
	}
	return false
}

var messageData MessageData

func (manager *AppImClientManager) SendFriendActionMessage(msg CreateFriendMessage) {
	message := messageData.GetCreateFriendMessage(msg)
	manager.SendMessageToSpecifiedClient([]byte(message), helpers.Int64ToString(msg.ToID))
}

func (manager *AppImClientManager) SendPrivateMessage(message requests.PrivateMessageRequest) (bool, string) {
	// 判断是否开启集群
	if config.Conf.Server.ServiceOpen {
		msgString := messageHandler.GetPrivateChatMessages(message, true)
		var dService dispatch.DispatchService
		ok, node := dService.IsDispatchNode(helpers.Int64ToString(message.ToID))
		// 判断用户在线并且返回节点信息
		if ok && node != "" {
			var messageClient GrpcClient.GrpcMessageService
			messageClient.SendGpcMessage(msgString, node)
			return true, "消息投递成功"
		}
	}
	msgString := messageHandler.GetPrivateChatMessages(message, true)

	// 将消费分发到不同的队列
	switch message.ChannelType {
	case 1:
		if !ImManager.SendMessageToSpecifiedClient([]byte(msgString), helpers.Int64ToString(message.ToID)) {
			nsq_queue.ProducerQueue.SendMessage([]byte(msgString))
		}
	case 2:

	default:
	}
	return true, "Success"

}
