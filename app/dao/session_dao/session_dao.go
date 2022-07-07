/**
  @author:panliang
  @data:2022/7/6
  @note
**/
package session_dao

import (
	"im-services/app/models/im_sessions"
	"im-services/app/models/user"
	"im-services/pkg/date"
	"im-services/pkg/model"
)

type SessionDao struct {
}

// 创建会话
func (s *SessionDao) CreateSession(formId int64, toId int64, channelType int) {

	var users user.ImUsers

	model.DB.Table("im_users").Where("id=?", toId).First(&users)
	session := im_sessions.ImSessions{
		ToId:        toId,
		FormId:      formId,
		CreatedAt:   date.NewDate(),
		TopStatus:   im_sessions.TOP_STATUS,
		TopTime:     date.NewDate(),
		Note:        users.Name,
		ChannelType: channelType,
		Name:        users.Name,
		Avatar:      users.Avatar,
		Status:      im_sessions.SESSION_STATUS_OK,
	}

	model.DB.Save(&session)

}
