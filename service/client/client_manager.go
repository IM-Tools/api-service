/**
  @author:panliang
  @data:2022/5/27
  @note
**/
package client

import (
	"fmt"
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
}

func (manager *ImClientManager) SetClient(client *ImClient) {
	mutexKey.Lock()
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
			fmt.Println(client)
		case client := <-ImManager.Unregister:
			fmt.Println(client)
		case message := <-ImManager.Broadcast:
			fmt.Println(message)
		}
	}
}

func (manager *ImClientManager) ImSend(message []byte, client *ImClient) {
	data, ok := manager.ImClientMap[client.ID]
	if ok {
		data.Send <- message
	}
}
