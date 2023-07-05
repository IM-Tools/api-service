package event

import (
	"encoding/json"
	"fmt"
	"im-services/internal/api/requests"
	"im-services/internal/api/services"
	"im-services/internal/dao/group_dao"
	"im-services/internal/enum"
	"im-services/internal/helpers"
	"im-services/internal/models/im_group_users"
	"im-services/internal/models/im_groups"
	"im-services/internal/models/user"
	"im-services/pkg/date"
	"im-services/pkg/model"
)

var (
	groupDao       group_dao.GroupDao
	messageService services.ImMessageService
)

// 注册后邀请该人进群
func (event *EventHandle) RegisterEvent(id int64, name string) {
	var group im_groups.ImGroups
	if result := model.DB.Model(&im_groups.ImGroups{}).Where("id=?", 39).Find(&group); result.RowsAffected == 0 {

		return
	}

	var selectUser = []string{fmt.Sprintf("%s", helpers.Int64ToString(id))}
	groupDao.CreateSelectGroupUser(selectUser, 39, group.Avatar, group.Name)
	// 发送群聊会话消息
	messageService.SendGroupSessionMessage(selectUser, group.Id)

	var users []user.ImUsers

	model.DB.Model(&user.ImUsers{}).
		Where("id in(?)", model.DB.Model(&im_group_users.ImGroupUsers{}).
			Where("group_id=?", group.Id).Select("user_id")).
		Find(&users)

	groupStr, _ := json.Marshal(group)
	message := requests.PrivateMessageRequest{
		MsgId:       date.TimeUnixNano(),
		MsgCode:     enum.WsChantMessage,
		MsgClientId: date.TimeUnixNano(),
		FormID:      group.Id,
		ChannelType: enum.GroupMessage,
		MsgType:     enum.JOIN_GROUP,
		Message:     "",
		SendTime:    date.NewDate(),
		Data:        string(groupStr),
		CreatedAt:   date.NewDate(),
	}

	messageService.SendCreateUserGroupMessage(users, message, name, 1, selectUser, id)
}
