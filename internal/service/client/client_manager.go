package client

import (
	"im-services/internal/api/requests"
	"im-services/internal/helpers"
	"im-services/pkg/coroutine_poll"
	"sync"
)

type AppImClientManager struct {
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
	ImManager = AppImClientManager{
		ImClientMap:      make(map[string]*ImClient),
		BroadcastChannel: make(chan []byte),
		PrivateChannel:   make(chan []byte),
		GroupChannel:     make(chan []byte),
		Register:         make(chan *ImClient),
		Unregister:       make(chan *ImClient),
	}
)

type AppClientManagerInterface interface {
	// SetClient 设置客户端信息
	SetClient(client *ImClient)
	// DelClient 删除客户端信息
	DelClient(client *ImClient)
	// Start 启动服务
	Start()
	// ImSend 消息投递到指定客户端 消息投递到指定客户端
	SendMessageToSpecifiedClient(message []byte, toId string) bool
	// LaunchPrivateMessage 私聊信息消费
	LaunchPrivateMessage(msgByte []byte)
	// LaunchGroupMessage 群聊信息消费
	LaunchGroupMessage(msgByte []byte)
	// LaunchBroadcastMessage 广播消息
	LaunchBroadcastMessage(msgByte []byte)
	// ConsumingOfflineMessages 消费离线消息
	ConsumingOfflineMessages(client *ImClient)
	// RadioUserOnlineStatus 向好友广播在线状态
	RadioUserOnlineStatus(client *ImClient)

	// GetOnlineNumber 获取在线人数
	GetOnlineNumber() int

	SendPrivateMessage(message requests.PrivateMessageRequest) (bool, string)

	SendFriendActionMessage(msg CreateFriendMessage)
}

func (manager *AppImClientManager) SetClient(client *ImClient) {
	manager.MutexKey.Lock()
	defer manager.MutexKey.Unlock()
	manager.ImClientMap[client.ID] = client

}

func (manager *AppImClientManager) DelClient(client *ImClient) {

	manager.MutexKey.Lock()
	client.Close()
	defer manager.MutexKey.Unlock()
	delete(manager.ImClientMap, client.ID)
}

func (manager *AppImClientManager) Start() {
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

func (manager *AppImClientManager) GetOnlineNumber() int {
	return len(manager.ImClientMap)
}

func (manager *AppImClientManager) IsUserOline(toId string) bool {
	_, ok := manager.ImClientMap[toId]
	if ok {
		return true
	}
	return false
}
