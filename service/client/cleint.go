/**
  @author:panliang
  @data:2022/5/27
  @note
**/
package client

import (
	"github.com/gorilla/websocket"
	"im-services/service/message"
	"sync"
)

type ImClient struct {
	ID     int64           // 客户端用户id
	Socket *websocket.Conn // 当前socket握手对象
	Send   chan []byte     // 当前用户发送的消息
	Mux    sync.RWMutex    // 互斥锁
}

var (
	messageHandler message.MessageHandler
)

type ClientInterface interface {
	Read()
	Write()
	Close()
}

// 返回一个客户端实例
// 并且设置当前客户端id以及socket握手实例
func NewClient(ID int64, conn *websocket.Conn) *ImClient {
	client := new(ImClient)
	client.ID = ID
	client.Send = make(chan []byte)
	client.Socket = conn
	return client
}

func (client *ImClient) Read() {

	defer func() {
		ImManager.Unregister <- client
		client.Socket.Close()
	}()

	for {
		_, msg, err := client.Socket.ReadMessage()
		if err != nil {

			ImManager.Unregister <- client
			client.Close()
			break
		}

		errs, msgString, ackMsg := messageHandler.ValidationMsg(msg)

		if errs != nil {
			client.Socket.WriteMessage(websocket.TextMessage,
				[]byte(msgString))
		} else {

			client.Socket.WriteMessage(websocket.TextMessage,
				[]byte(ackMsg))
			ImManager.Broadcast <- []byte(msgString)

		}
	}

}

func (client *ImClient) Write() {
	defer client.Close()
	for {
		select {
		case message, ok := <-client.Send:
			if !ok {
				client.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			client.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func (client *ImClient) Close() {
	client.Socket.Close()
}
