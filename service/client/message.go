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
	"im-services/pkg/logger"
	"im-services/service/cache/firend"
	"im-services/service/dao"
	"im-services/service/queue/nsq_queue"
)

// 消息分发
func (manager *ImClientManager) LaunchMessage(message []byte) {
	var p fastjson.Parser
	v, _ := p.Parse(string(message))
	channelType := v.GetInt("channel_type")
	ReceiveId := v.GetInt64("receive_id")
	msg := v.Get("msg").String()
	if channelType == 1 || channelType == 3 {
		if client, ok := manager.ImClientMap[ReceiveId]; ok {
			// todo 消息持久化
			logger.Logger.Info("消息已经投递" + msg)
			client.Send <- []byte(msg)
		} else {
			logger.Logger.Info("用户离线了" + string(message))
			//离线消息进入nsq
			nsq_queue.ProducerQueue.SendMessage([]byte(msg))

		}
	} else {
		// todo 群聊消息
	}
}

// 消费离线消息
func (manager *ImClientManager) ConsumingOfflineMessages(client *ImClient) {

	// 读取离线消息
	list := dao.OfflineMessage.PullPrivateOfflineMessage(client.ID)
	logger.Logger.Info(fmt.Sprintf("ConsumingOfflineMessages 客户端链接:%d", client.ID))
	for _, value := range list {
		logger.Logger.Info(fmt.Sprintf("消息消费 客户端链接:%d", value.Message))
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
	data, err := firend.FriendCache.Get(client.ID)
	if err != nil {

	}
	for _, val := range data {
		logger.Logger.Info("好友")
		if friendClient, ok := manager.ImClientMap[val.FId]; ok {
			// todo 消息持久化
			friendClient.Socket.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(`{"code":200,"message":"用户上线了"',"fo_id":%d}`, int(val.MId))))
		}
	}
}
