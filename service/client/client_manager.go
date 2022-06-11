/**
  @author:panliang
  @data:2022/5/27
  @note
**/
package client

import (
	"fmt"
	"im-services/pkg/coroutine_poll"
	"im-services/pkg/logger"
	"sync"
)

type ImClientManager struct {
	ImClientMap map[int64]*ImClient
	Broadcast   chan []byte
	Register    chan *ImClient
	Unregister  chan *ImClient
	MutexKey    sync.RWMutex //读写锁
}

var (
	ImManager = ImClientManager{
		ImClientMap: make(map[int64]*ImClient),
		Broadcast:   make(chan []byte),
		Register:    make(chan *ImClient),
		Unregister:  make(chan *ImClient),
	}
)

type ClientManagerInterface interface {
	SetClient(client *ImClient)              // 设置客户端信息
	DelClient(client *ImClient)              // 删除客户端信息
	Start()                                  // 启动服务
	ImSend(message []byte, client *ImClient) // 给指定客户端投递消息 该方法可能用不着了..
	LaunchMessage(msg_byte []byte)
	ConsumingOfflineMessages(client *ImClient) // 消费离线消息
	RadioUserOnlineStatus(client *ImClient)    // 向好友广播在线状态
	GetOnlineNumber() int                      //在线人数
}

func (manager *ImClientManager) SetClient(client *ImClient) {
	manager.MutexKey.Lock()
	defer manager.MutexKey.Unlock()
	logger.Logger.Info(fmt.Sprintf("客户端链接:%d", client.ID))
	manager.ImClientMap[client.ID] = client

}

func (manager *ImClientManager) DelClient(client *ImClient) {
	manager.MutexKey.Lock()
	client.Close()
	defer manager.MutexKey.Unlock()
	delete(manager.ImClientMap, client.ID)
}

func (manager *ImClientManager) Start() {
	for {
		select {
		case client := <-ImManager.Register:
			// 设置客户端 拉去离线消息 推送在线状态
			manager.SetClient(client)
			manager.ConsumingOfflineMessages(client)
			manager.RadioUserOnlineStatus(client)

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

func (manager *ImClientManager) GetOnlineNumber() int {
	manager.MutexKey.RLock()
	defer manager.MutexKey.RUnlock()
	return len(manager.ImClientMap)
}
