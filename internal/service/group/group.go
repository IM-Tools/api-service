package group

import (
	"im-services/internal/helpers"
	"im-services/internal/models/im_groups"
	"sync"
)

type Groups struct {
	Id           string `json:"id"`
	OwnerId      int64  `json:"owner_id"`
	Name         string `json:"group_name"`
	Info         string `json:"info"`
	OnlineNumber int    `json:"online_number"`
	TotalNumber  int    `json:"total_number"`
	Avatar       string `json:"avatar"`
	UserGather   map[string]string
	MutexKey     sync.RWMutex
}

type GroupsInterface interface {
}

func NewGroup(imGroups im_groups.ImGroups) *Groups {
	groups := new(Groups)
	groups.Id = helpers.Int64ToString(imGroups.Id)
	groups.Name = imGroups.Name
	groups.Info = imGroups.Info
	groups.Avatar = imGroups.Avatar
	groups.OnlineNumber = 1
	groups.TotalNumber = 1
	return groups
}

// 将用户id添加到组内
func (group *Groups) AddGroupNumber(userId string) {
	group.MutexKey.Lock()
	group.UserGather[userId] = userId
	defer group.MutexKey.Unlock()
}

func (group *Groups) DetGroup() {

}

// 组内用户在线++
func (group *Groups) IncrementGroupOnlineNumber() {
	group.MutexKey.Lock()
	group.OnlineNumber++
	defer group.MutexKey.Unlock()
}

// 组内用户在线--
func (group *Groups) DecreaseGroupOnlineNumber() {
	group.MutexKey.Lock()
	group.OnlineNumber--
	defer group.MutexKey.Unlock()
}
