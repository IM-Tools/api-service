package group_dao

import (
	"im-services/internal/api/requests"
	"im-services/internal/helpers"
	"im-services/internal/models/im_group_users"
	"im-services/internal/models/im_groups"
	"im-services/internal/models/im_sessions"
	"im-services/pkg/date"
	"im-services/pkg/model"
)

type GroupDao struct {
}

type ImGroups struct {
	Id        int64  `gorm:"column:id" json:"id"`                 //群聊id
	UserId    int64  `gorm:"column:user_id" json:"user_id"`       //创建者
	Name      string `gorm:"column:name" json:"name"`             //群聊名称
	CreatedAt string `gorm:"column:created_at" json:"created_at"` //添加时间
	Info      string `gorm:"column:info" json:"info"`             //群聊描述
	Avatar    string `gorm:"column:avatar" json:"avatar"`         //群聊头像
	Password  string `gorm:"column:password" json:"password"`     //密码
	IsPwd     int8   `gorm:"column:is_pwd" json:"is_pwd"`         //是否加密 0 否 1 是
	Hot       int    `gorm:"column:hot" json:"hot"`               //热度
}

// 将人员添加到群组表中 并创建会话
func (*GroupDao) CreateSelectGroupUser(userIds []string, groupId int, avatar string, name string) {

	count := len(userIds)
	createdAt := date.NewDate()
	var groupUser = make([]im_group_users.ImGroupUsers, count)
	var sessionsData = make([]im_sessions.ImSessions, count)
	for key, id := range userIds {
		groupUser[key].UserId = helpers.StringToInt(id)
		groupUser[key].CreatedAt = createdAt
		groupUser[key].Avatar = avatar
		groupUser[key].GroupId = groupId
		groupUser[key].Name = name

		sessionsData[key].FormId = helpers.StringToInt64(id)
		sessionsData[key].GroupId = int64(groupId)
		sessionsData[key].CreatedAt = createdAt
		sessionsData[key].ChannelType = im_sessions.GROUP_TYPE
		sessionsData[key].Name = name
		sessionsData[key].Avatar = avatar
		sessionsData[key].TopTime = date.NewDate()
	}
	model.DB.Model(&im_group_users.ImGroupUsers{}).Create(&groupUser)
	model.DB.Model(&im_sessions.ImSessions{}).Create(&sessionsData)
	return
}

// 将人员添加到群组表中 并创建会话
func (*GroupDao) CreateOneGroupUser(group im_groups.ImGroups, id int) {

	var groupUser im_group_users.ImGroupUsers

	groupUser.UserId = id
	groupUser.CreatedAt = date.NewDate()
	groupUser.Avatar = group.Avatar
	groupUser.GroupId = int(group.Id)
	groupUser.Name = group.Name
	model.DB.Model(&im_group_users.ImGroupUsers{}).Create(&groupUser)
	var session im_sessions.ImSessions

	session.FormId = int64(id)
	session.ToId = group.Id
	session.CreatedAt = date.NewDate()
	session.ChannelType = im_sessions.GROUP_TYPE
	session.Name = group.Name
	session.Avatar = group.Avatar
	model.DB.Model(&im_sessions.ImSessions{}).Create(&session)

	return
}

// 创建群聊
func (*GroupDao) CreateGroup(params requests.CreateGroupRequest) (error, im_groups.ImGroups) {

	var imGroups im_groups.ImGroups
	imGroups.UserId = params.UserId
	imGroups.Name = params.Name
	imGroups.Info = params.Info
	imGroups.Avatar = params.Avatar
	imGroups.CreatedAt = date.NewDate()
	if model.DB.Model(&im_groups.ImGroups{}).Create(&imGroups).Error != nil {
		var err error
		return err, imGroups
	}
	return nil, imGroups
}

// 判断用户是否入群
func (*GroupDao) IsGroupsUser(userId interface{}, groupId interface{}) bool {
	var count int64
	model.DB.Model(&im_group_users.ImGroupUsers{}).Where("user_id = ? and group_id = ?", userId, groupId).Count(&count)
	if count > 0 {
		return true
	}
	return false
}

// 查询群聊用户数据
func (*GroupDao) GetGroupUsers(groupId string) []im_group_users.ImGroupUsers {
	var groupUser []im_group_users.ImGroupUsers
	model.DB.Model(&im_group_users.ImGroupUsers{}).Where("group_id=?", groupId).Preload("Users").Find(&groupUser)
	return groupUser
}

func (*GroupDao) DeleteGroupUser(id interface{}, groupId string) {
	var groupUsers im_group_users.ImGroupUsers
	model.DB.Model(&im_group_users.ImGroupUsers{}).
		Where("user_id = ? and group_id =?", id, groupId).
		Delete(&groupUsers)
}
