/**
  @author:panliang
  @data:2022/6/22
  @note
**/
package dispatch

import (
	"github.com/gorilla/websocket"
	"im-services/config"
	"im-services/pkg/redis"
	"sync"
)

var (
	node = config.Conf.Server.Node
	mux  sync.Mutex
)

type DispatchService struct {
}

type DispatchServiceInterface interface {
	SetDispatchNode(uid string, node string) // 设置当前节点信息
	GetDispatchNode(uid string, node string) // 获取当前节点信息
	MessageDispatch(uid string, node string) // 获取当前节点信息
	IsDispatchNode(uid string, node string)  // 获取当前节点信息
}

func (Service *DispatchService) IsDispatchNode(uid string) (bool, string) {

	n, _ := redis.RedisDB.Exists(uid).Result()
	if n > 0 {
		return true, ""
	} else {
		uNode := Service.GetDispatchNode(uid)
		if uNode != node {
			return false, uNode
		}
		return true, ""
	}
}

func (Service *DispatchService) GetDispatchNode(uid string) string {
	return redis.RedisDB.Get(uid).Val()
}

func (Service *DispatchService) SetDispatchNode(uid string) {
	mux.Lock()
	redis.RedisDB.Set(uid, node, 3600*10)
	mux.Unlock()
}

func (Service *DispatchService) MessageDispatch(conn *websocket.Conn) {

}
