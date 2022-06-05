/**
  @author:panliang
  @data:2022/5/27
  @note
**/
package client

import (
	"fmt"
	"github.com/valyala/fastjson"
	"im-services/pkg/coroutine_poll"
	"im-services/pkg/logger"
	"sync"
)

type ImClientManager struct {
	ImClientMap map[int64]*ImClient
	Broadcast   chan []byte
	Register    chan *ImClient
	Unregister  chan *ImClient
}

var (
	ImManager = ImClientManager{
		ImClientMap: make(map[int64]*ImClient),
		Broadcast:   make(chan []byte),
		Register:    make(chan *ImClient),
		Unregister:  make(chan *ImClient),
	}
	mutexKey sync.Mutex
)

type ClientManagerInterface interface {
	SetClient(client *ImClient)              // 设置客户端信息
	DelClient(client *ImClient)              // 删除客户端信息
	Start()                                  // 启动服务
	ImSend(message []byte, client *ImClient) // 给指定客户端投递消息 该方法可能用不着了..
	LaunchMessage(msg_byte []byte)
}

func (manager *ImClientManager) SetClient(client *ImClient) {
	mutexKey.Lock()
	logger.Logger.Info(fmt.Sprintf("客户端链接:%d", client.ID))
	manager.ImClientMap[client.ID] = client
	mutexKey.Unlock()
}

func (manager *ImClientManager) DelClient(client *ImClient) {
	client.Close()
	mutexKey.Lock()
	delete(manager.ImClientMap, client.ID)
	mutexKey.Unlock()
}

func (manager *ImClientManager) Start() {
	for {
		select {
		case client := <-ImManager.Register:
			manager.SetClient(client)
			logger.Logger.Debug(fmt.Sprintf("注册的客户端:%s", client.ID))

		case client := <-ImManager.Unregister:
			manager.DelClient(client)
			logger.Logger.Debug(fmt.Sprintf("离线的客户端%s:", client.ID))

		case message := <-ImManager.Broadcast:
			coroutine_poll.AntsPool.Submit(func() {
				manager.LaunchMessage(message)
			})
			logger.Logger.Debug(fmt.Sprintf("收到的消息:%s", string(message)))
		}
	}
}

func (manager *ImClientManager) ImSend(message []byte, client *ImClient) {
	data, ok := manager.ImClientMap[client.ID]
	if ok {
		data.Send <- message
	}
}

func (manager *ImClientManager) LaunchMessage(message []byte) {
	var p fastjson.Parser
	v, _ := p.Parse(string(message))
	channelType := v.GetInt("channel_type")
	ReceiveId := v.GetInt64("receive_id")
	logger.Logger.Info("消息方法" + string(message))
	if channelType == 1 || channelType == 3 {
		if client, ok := manager.ImClientMap[ReceiveId]; ok {
			// todo 消息持久化
			logger.Logger.Info("消息已经投递" + string(message))
			client.Send <- message
		} else {
			logger.Logger.Info("用户离线了" + string(message))
		}
	} else {
		// todo 群聊消息
	}
}
