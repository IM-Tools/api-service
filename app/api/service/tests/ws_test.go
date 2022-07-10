/**
  @author:panliang
  @data:2022/6/12
  @note
**/
package tests

import (
	"fmt"
	"github.com/gorilla/websocket"
	"im-services/config"
	"im-services/pkg/jwt"
	"log"
	"math/rand"
	"sync"
	"testing"
)

func init() {
	config.InitConfig("config.yaml")
}

var (
	ClientMapALL = ClientMap{
		ClientMap: make(map[int64]*Client),
	}
	userCount int64 = 10000
)

type ClientMap struct {
	ClientMap map[int64]*Client
}

type Client struct {
	Conn *websocket.Conn `json:"conn"`
	ID   int64           `json:"ID"`
}

var wg sync.WaitGroup
var mux sync.RWMutex

func TestWs(t *testing.T) {
	var i int64
	i = 1
	for i = 0; i < userCount; i++ {
		token := jwt.NewJWT().IssueToken(
			i,
			"",
			fmt.Sprintf("用户%d", i),
			fmt.Sprintf("%dxxx@qq.com", i),
			1749738653,
		)
		dl := websocket.Dialer{}
		conn, _, err := dl.Dial("ws://127.0.0.1:8000/im/connect?token="+token, nil)
		if err != nil {
			log.Println(fmt.Sprintf("用户%d建立链接失败！！", i))
			return
		}
		defer conn.Close()

		client := new(Client)
		client.ID = i
		client.Conn = conn
		ClientMapALL.ClientMap[i] = client

	}

	wg.Add(60)

	var x int64
	for x = 0; x < 10; x++ {
		go testUserSendMsg()
		go testUserSendMsg()
		go testUserSendMsg()
		go testUserSendMsg()
		go testUserSendMsg()
	}
	wg.Wait()

}

func testUserSendMsg() {
	defer wg.Done()

	var x int64 = 1
	for x = 1; x < 1000; x++ {
		userId := getUserInt(userCount, 1)
		mux.RLock()
		if data, ok := ClientMapALL.ClientMap[userId]; ok {
			sendMsg(data, userId, getUserInt(10000000, userId))
		}
		mux.RUnlock()
	}
}

func getUserInt(count int64, repeat int64) int64 {

	userId := int64(rand.Intn(int(count)))
	if repeat == userId {
		getUserInt(count, repeat)
	}
	return int64(userId)
}

func sendMsg(client *Client, formId int64, toiD int64) {

	msg := fmt.Sprintf(`{"msg_id":1,"msg_client_id":1,"msg_code":200,"form_id":%d,"to_id":%d,"msg_type":1,"channel_type":1,"message":"你好！"}`, formId, toiD)

	fmt.Sprintf("%d向%d发送消息", formId, toiD)
	mux.Lock()

	defer mux.Unlock()

	client.Conn.WriteMessage(websocket.TextMessage, []byte(msg))

}
