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
		sessionsData[key].ToId = int64(groupId)
		sessionsData[key].CreatedAt = createdAt
		sessionsData[key].ChannelType = im_sessions.GroupType
		sessionsData[key].Name = name
		sessionsData[key].Avatar = avatar
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
	session.ChannelType = im_sessions.GroupType
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
