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
	"im-services/pkg/logger"
	"im-services/service/cache/firend_cache"
	"im-services/service/dao"
	"im-services/service/queue/nsq_queue"
)

//
func (manager *ImClientManager) LaunchPrivateMessage(message []byte) {
	var p fastjson.Parser
	v, _ := p.Parse(string(message))
	ReceiveId := v.Get("receive_id").String()
	msg := v.Get("msg").String()
	if client, ok := manager.ImClientMap[ReceiveId]; ok {
		client.Send <- []byte(msg)
	} else {
		nsq_queue.ProducerQueue.SendMessage([]byte(msg))
	}
}

func (manager *ImClientManager) LaunchBroadcastMessage(message []byte) {
	logger.Logger.Info("广播消息")
	var p fastjson.Parser
	v, _ := p.Parse(string(message))

	msgCode, _ := v.Get("msg_code").Int()

	var ReceiveId string
	if msgCode == enum.WS_CREATE {
		ReceiveId = v.Get("to_id").String()
	} else {
		ReceiveId = v.Get("form_id").String()
	}
	
	logger.Logger.Info("广播消息id:" + ReceiveId)
	if client, ok := manager.ImClientMap[ReceiveId]; ok {
		client.Send <- message
	}
}

func (manager *ImClientManager) LaunchGroupMessage(message []byte) {

	var p fastjson.Parser
	v, _ := p.Parse(string(message))
	ReceiveId := v.Get("receive_id").String()
	msg := v.Get("msg").String()
	if client, ok := manager.ImClientMap[ReceiveId]; ok {
		client.Send <- []byte(msg)
	} else {
		logger.Logger.Info("用户离线了" + string(message))
		//离线消息进入nsq
		nsq_queue.ProducerQueue.SendMessage([]byte(msg))

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
			// todo 消息持久化
			friendClient.Socket.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(`{"code":200,"message":"用户上线了"',"fo_id":%d}`, int(val.ToId))))
		}
	}
}
