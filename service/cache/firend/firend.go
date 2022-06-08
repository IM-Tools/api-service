/**
  @author:panliang
  @data:2022/6/8
  @note
**/
package firend

import (
	"im-services/app/models/friend"
	"im-services/pkg/model"
	"sync"
)

type CacheInterface interface {
	Set(id int64, friends *[]friend.ImFriends) // 设置缓存
	Get(id int64) (*[]friend.ImFriends, error) // 读取缓存
}

var (
	FriendCache = FriendCacheClient{
		CachetMap: make(map[int64]*[]friend.ImFriends),
	}
	mux sync.Mutex
)

type FriendCacheClient struct {
	CachetMap map[int64]*[]friend.ImFriends
}

// 设置好友缓存
func (FriendCache *FriendCacheClient) Set(id int64, friends *[]friend.ImFriends) {
	mux.Lock()
	FriendCache.CachetMap[id] = friends
	mux.Unlock()
}

// 获取好友缓存
func (FriendCache *FriendCacheClient) Get(id int64) ([]friend.ImFriends, error) {
	var err error
	data, ok := FriendCache.CachetMap[id]
	if ok {
		return *data, nil
	}

	var list []friend.ImFriends
	model.DB.Table("im_friends").Where("m_id=?", id).Find(&list)

	FriendCache.Set(id, &list)

	return list, err
}
