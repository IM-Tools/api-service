/**
  @author:panliang
  @data:2022/6/8
  @note
**/
package client

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/valyala/fastjson"
	"im-services/app/enum"
	"im-services/app/service/cache/firend_cache"
	"im-services/app/service/dao"
	"im-services/app/service/queue/nsq_queue"
	"im-services/pkg/logger"
)

//
func (manager *ImClientManager) LaunchPrivateMessage(message []byte) {

	receiveId, userMsg := GetReceiveIdAndUserMsg(message)

	logger.Logger.Info("私聊消息" + string(userMsg))
	if client, ok := manager.ImClientMap[receiveId]; ok {
		client.Send <- []byte(userMsg)
	} else {
		nsq_queue.ProducerQueue.SendMessage([]byte(userMsg))
	}

}

func (manager *ImClientManager) LaunchBroadcastMessage(message []byte) {

	var p fastjson.Parser
	v, _ := p.Parse(string(message))

	msgCode, _ := v.Get("msg_code").Int()

	var ReceiveId string
	if msgCode == enum.WS_CREATE {
		ReceiveId = v.Get("to_id").String()
	} else {
		ReceiveId = v.Get("form_id").String()
	}

	if client, ok := manager.ImClientMap[ReceiveId]; ok {
		client.Send <- message
	}
}

func (manager *ImClientManager) LaunchGroupMessage(message []byte) {

	receiveId, userMsg := GetReceiveIdAndUserMsg(message)

	if client, ok := manager.ImClientMap[receiveId]; ok {
		client.Send <- []byte(userMsg)
	} else {
		nsq_queue.ProducerQueue.SendMessage([]byte(userMsg))

	}
}

// 消费离线消息
func (manager *ImClientManager) ConsumingOfflineMessages(client *ImClient) {
	// 读取离线消息
	list := dao.OfflineMessage.PullPrivateOfflineMessage(client.ID)
	for _, value := range list {
		client.Socket.WriteMessage(websocket.TextMessage, []byte(value.Message))
	}
	// 更新离线消息状态
	if len(list) > 0 {
		dao.OfflineMessage.UpdatePrivateOfflineMessageStatus(client.ID)
	}
}

// 广播在线用户在线状态
func (manager *ImClientManager) RadioUserOnlineStatus(client *ImClient) {
	// 从数据库拿好友列表id 从客户端拿好友在线id 进行在线状态推送
	data, err := firend_cache.FriendCache.Get(client.ID)
	if err != nil {

	}
	for _, val := range data {
		if friendClient, ok := manager.ImClientMap[val.Uid]; ok {
			_ = friendClient.Socket.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(`{"code":200,"message":"用户上线了"',"fo_id":%d}`, int(val.ToId))))
		}
	}
}

// 拿消息投递id
func GetReceiveIdAndUserMsg(msg []byte) (string, string) {
	var p fastjson.Parser
	v, _ := p.Parse(string(msg))
	return fastjson.GetString(msg, "receive_id"), v.GetObject("msg").String()
}
