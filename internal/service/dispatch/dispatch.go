package dispatch

import (
	"im-services/internal/config"
	"im-services/pkg/redis"
	"time"
)

type DispatchService struct {
}

type DispatchServiceInterface interface {
	SetDispatchNode(uid string, node string) // 设置当前节点信息
	GetDispatchNode(uid string, node string) // 获取当前节点信息
	MessageDispatch(uid string, node string) // 获取当前节点信息
	IsDispatchNode(uid string, node string)  // 获取当前节点信息
	DetDispatchNode(uid string)              //删除当前节点
}

func (Service *DispatchService) IsDispatchNode(uid string) (bool, string) {

	n, _ := redis.RedisDB.Exists(uid).Result()

	if n > 0 {
		uNode := Service.GetDispatchNode(uid)
		return true, uNode
	} else {
		return false, ""
	}
}

func (Service *DispatchService) GetDispatchNode(uid string) string {
	return redis.RedisDB.Get(uid).Val()
}

func (Service *DispatchService) DetDispatchNode(uid string) {
	redis.RedisDB.Del(uid)
}

func (Service *DispatchService) SetDispatchNode(uid string) {

	redis.RedisDB.Set(uid, config.Conf.Server.Node, time.Hour*24)
}

func (Service *DispatchService) MessageDispatch() {

}
