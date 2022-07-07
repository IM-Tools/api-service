/**
  @author:panliang
  @data:2022/7/6
  @note
**/
package friend_dao

import (
	"im-services/app/models/im_friends"
	"im-services/pkg/date"
	"im-services/pkg/model"
)

type FriendDao struct {
}

// 添加好友
func (f *FriendDao) AgreeFriendRequest(toId int64, formId int64) {

	friend := im_friends.ImFriends{
		FormId:    formId,
		ToId:      toId,
		Note:      "",
		CreatedAt: date.NewDate(),
		UpdatedAt: date.NewDate(),
		TopTime:   date.NewDate(),
		Status:    0,
	}
	model.DB.Save(&friend)
}
