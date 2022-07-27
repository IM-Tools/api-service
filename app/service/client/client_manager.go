package client

import (
	"im-services/app/helpers"
	"im-services/pkg/coroutine_poll"
	"im-services/pkg/logger"
	"sync"
)

type ImClientManager struct {
	// 储存客户端实例
	ImClientMap map[string]*ImClient
	// 公共频道
	BroadcastChannel chan []byte
	// 私聊频道
	PrivateChannel chan []byte
	// 群聊频道
	GroupChannel chan []byte
	// 注册客户端
	Register chan *ImClient
	// 关闭客户端
	Unregister chan *ImClient
	// 锁
	MutexKey sync.RWMutex
}

var (
	ImManager = ImClientManager{
		ImClientMap:      make(map[string]*ImClient),
		BroadcastChannel: make(chan []byte),
		PrivateChannel:   make(chan []byte),
		GroupChannel:     make(chan []byte),
		Register:         make(chan *ImClient),
		Unregister:       make(chan *ImClient),
	}
)

type ClientManagerInterface interface {
	// SetClient 设置客户端信息
	SetClient(client *ImClient)
	// DelClient 删除客户端信息
	DelClient(client *ImClient)
	// Start 启动服务
	Start()
	// ImSend 消息投递到指定客户端 消息投递到指定客户端
	ImSend(message []byte, client *ImClient)
	// LaunchPrivateMessage 私聊信息消费
	LaunchPrivateMessage(msg_byte []byte)
	// LaunchGroupMessage 群聊信息消费
	LaunchGroupMessage(msg_byte []byte)
	// LaunchBroadcastMessage 广播消息
	LaunchBroadcastMessage(msg_byte []byte)
	// ConsumingOfflineMessages 消费离线消息
	ConsumingOfflineMessages(client *ImClient)
	// RadioUserOnlineStatus 向好友广播在线状态
	RadioUserOnlineStatus(client *ImClient)
	// GetOnlineNumber 获取在线人数
	GetOnlineNumber() int
}

func (manager *ImClientManager) SetClient(client *ImClient) {
	manager.MutexKey.Lock()
	defer manager.MutexKey.Unlock()
	manager.ImClientMap[client.ID] = client

}

func (manager *ImClientManager) DelClient(client *ImClient) {

	manager.MutexKey.Lock()
	client.Close()
	defer manager.MutexKey.Unlock()
	logger.Logger.Info("客户端断开:" + client.ID)
	delete(manager.ImClientMap, client.ID)
}

func (manager *ImClientManager) Start() {
	for {
		select {
		case client := <-ImManager.Register:
			// 设置客户端 拉去离线消息 推送在线状态
			manager.SetClient(client)
			manager.ConsumingOfflineMessages(client)
			//manager.RadioUserOnlineStatus(client)
		case client := <-ImManager.Unregister:
			manager.DelClient(client)

		case message := <-ImManager.PrivateChannel:
			err := coroutine_poll.AntsPool.Submit(func() {
				manager.LaunchPrivateMessage(message)
			})
			helpers.ErrorHandler(err)
		case groupMessage := <-ImManager.GroupChannel:
			err := coroutine_poll.AntsPool.Submit(func() {
				manager.LaunchPrivateMessage(groupMessage)
			})
			helpers.ErrorHandler(err)
		case publicMessage := <-ImManager.BroadcastChannel:
			err := coroutine_poll.AntsPool.Submit(func() {
				manager.LaunchBroadcastMessage(publicMessage)
			})
			helpers.ErrorHandler(err)
		}

	}
}

func (manager *ImClientManager) ImSend(message []byte, client *ImClient) {
	data, ok := manager.ImClientMap[client.ID]
	if ok {
		data.Send <- message
	}
}

func (manager *ImClientManager) GetOnlineNumber() int {
	manager.MutexKey.RLock()
	defer manager.MutexKey.RUnlock()
	return len(manager.ImClientMap)
}
