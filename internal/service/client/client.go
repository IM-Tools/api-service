package client

import (
	"github.com/gorilla/websocket"
	"im-services/internal/service/dispatch"
	"im-services/pkg/logger"
	GrpcClient "im-services/server/client"
	"sync"
)

type ImClient struct {
	ID       string          // 客户端用户id
	Uuid     string          // 用户唯一id
	Socket   *websocket.Conn // 当前socket握手对象
	Send     chan []byte     // 当前用户发送的消息
	Mux      sync.RWMutex    // 互斥锁
	Identity int             // 身份 1.游客 2.用户
}

var (
	messageHandler MessageHandler
	grpcClient     GrpcClient.GrpcMessageService
	dispatchNode   dispatch.DispatchService
)

type WsClientInterface interface {
	Read()
	Write()
	Close()
}

// NewClient 返回一个客户端实例
// 并且设置当前客户端id以及socket握手实例
func NewClient(ID string, uid string, identity int, conn *websocket.Conn) *ImClient {
	client := new(ImClient)
	client.ID = ID
	client.Uuid = uid
	client.Identity = identity
	client.Send = make(chan []byte)
	client.Socket = conn
	return client
}

func (client *ImClient) Read() {

	defer func() {
		ImManager.Unregister <- client
		_ = client.Socket.Close()
	}()

	for {
		_, msg, err := client.Socket.ReadMessage()
		if err != nil {
			ImManager.Unregister <- client
			client.Close()
			break
		}

		errs, msgByte, ackMsg, channel, node := messageHandler.ValidationMsg(msg)

		if errs != nil {
			logger.Logger.Info(string(msgByte))
			_ = client.Socket.WriteMessage(websocket.TextMessage, msgByte)
		} else {
			// 将消费分发到不同的队列
			switch channel {
			case PRIVATE:
				_ = client.Socket.WriteMessage(websocket.TextMessage,
					ackMsg)
				ImManager.PrivateChannel <- msgByte
			case GROUP:
				_ = client.Socket.WriteMessage(websocket.TextMessage,
					msgByte)
			case PING:
				_ = client.Socket.WriteMessage(websocket.TextMessage,
					ackMsg)
			case FORWARDING:
				_ = client.Socket.WriteMessage(websocket.TextMessage,
					msgByte)
				grpcClient.SendGpcMessage(string(msgByte), node)
			default:
				_ = client.Socket.WriteMessage(websocket.TextMessage,
					ackMsg)
				ImManager.BroadcastChannel <- msgByte
			}
		}
	}

}

func (client *ImClient) Write() {
	defer client.Close()
	for {
		select {
		case msg, ok := <-client.Send:
			if !ok {
				_ = client.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			_ = client.Socket.WriteMessage(websocket.TextMessage, msg)
		}
	}
}

func (client *ImClient) Close() {
	dispatchNode.DetDispatchNode(client.ID)
	_ = client.Socket.Close()
}
