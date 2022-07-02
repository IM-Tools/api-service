/**
  @author:panliang
  @data:2022/6/22
  @note
**/
package dispatch

import (
	"github.com/gorilla/websocket"
	"im-services/app/helpers"
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
	SetDispatchNode(uid int64, node string) // 设置当前节点信息
	GetDispatchNode(uid int64, node string) // 获取当前节点信息
	MessageDispatch(uid int64, node string) // 获取当前节点信息
	IsDispatchNode(uid int64, node string)  // 获取当前节点信息
}

func (Service *DispatchService) IsDispatchNode(uid int64) (bool, string) {

	n, _ := redis.RedisDB.Exists(helpers.Int64ToString(uid)).Result()
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

func (Service *DispatchService) GetDispatchNode(uid int64) string {
	return redis.RedisDB.Get(helpers.Int64ToString(uid)).Val()
}

func (Service *DispatchService) SetDispatchNode(uid int64) {
	mux.Lock()
	redis.RedisDB.Set(helpers.Int64ToString(uid), node, 3600*10)
	mux.Unlock()
}

func (Service *DispatchService) MessageDispatch(conn *websocket.Conn) {

}
