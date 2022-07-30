package friend_dao

import (
	"im-services/app/models/im_friends"
	"im-services/pkg/date"
	"im-services/pkg/model"
)

type FriendDao struct {
}

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

type APIUsers struct {
	ID     int64  `gorm:"column:id;primaryKey" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
	Email  string `gorm:"column:email" json:"email"`
	Avatar string `gorm:"column:avatar" json:"avatar"`
	Status int8   `gorm:"column:status" json:"status"`
	Bio    string `gorm:"column:bio" json:"bio"`
	Sex    int8   `gorm:"column:sex" json:"sex"`
}

// GetNotFriendList 获取非好友数据
func (f *FriendDao) GetNotFriendList(id interface{}, email string) []APIUsers {
	var users []APIUsers

	sqlQuery := model.DB.Table("im_friends").Where("form_id=?", id).Select("to_id")

	model.DB.Table("im_users").
		Where("email like ?", email).
		Where("id not in(?)", sqlQuery).
		Select("id,name,email,avatar,bio,sex,status").
		Limit(5).
		Find(&users)

	return users
}

// 查询好友详情
func (f *FriendDao) GetFriends(id interface{}) (error, interface{}) {
	var err error
	var list im_friends.ImFriends
	result := model.DB.Model(&im_friends.ImFriends{}).Preload("Users").
		Where("form_id=?", id).
		Order("status desc").
		Order("top_time desc").
		Find(&list)
	if result.RowsAffected == 0 {
		return err, list
	}
	return nil, list
}

// 查询好友详情
func (f *FriendDao) GetFriendLists(id interface{}) (error, interface{}) {
	var err error
	var list []im_friends.ImFriends
	result := model.DB.Model(&im_friends.ImFriends{}).Preload("Users").
		Where("form_id=?", id).
		Order("status desc").
		Order("top_time desc").
		Find(&list)
	if result.RowsAffected == 0 {
		return err, list
	}
	return nil, list
}
