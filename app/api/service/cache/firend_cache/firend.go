/**
  @author:panliang
  @data:2022/6/8
  @note
**/
package firend_cache

import (
	"im-services/app/models/im_friends"
	"im-services/pkg/model"
	"sync"
)

type CacheInterface interface {
	Set(uid string, friends *[]im_friends.ImFriends) // 设置缓存
	Get(uid string) (*[]im_friends.ImFriends, error) // 读取缓存
}

var (
	FriendCache = FriendCacheClient{
		CachetMap: make(map[string]*[]im_friends.ImFriends),
	}
	mux sync.Mutex
)

type FriendCacheClient struct {
	CachetMap map[string]*[]im_friends.ImFriends
}

// 设置好友缓存
func (FriendCache *FriendCacheClient) Set(id string, friends *[]im_friends.ImFriends) {
	mux.Lock()
	FriendCache.CachetMap[id] = friends
	mux.Unlock()
}

// 获取好友缓存
func (FriendCache *FriendCacheClient) Get(id string) ([]im_friends.ImFriends, error) {
	var err error
	data, ok := FriendCache.CachetMap[id]
	if ok {
		return *data, nil
	}

	var list []im_friends.ImFriends
	model.DB.Table("im_friends").Where("m_id=?", id).Find(&list)

	FriendCache.Set(id, &list)

	return list, err
}
