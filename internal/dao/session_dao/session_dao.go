package session_dao

import (
	"im-services/internal/models/im_sessions"
	"im-services/internal/models/user"
	"im-services/pkg/date"
	"im-services/pkg/model"
)

type SessionDao struct {
}

func (s *SessionDao) CreateSession(formId int64, toId int64, channelType int) {

	var users user.ImUsers
	model.DB.Table("im_users").Where("id=?", toId).First(&users)
	session := im_sessions.ImSessions{
		ToId:        toId,
		FormId:      formId,
		CreatedAt:   date.NewDate(),
		TopStatus:   im_sessions.TopStatus,
		TopTime:     date.NewDate(),
		Note:        users.Name,
		ChannelType: channelType,
		Name:        users.Name,
		Avatar:      users.Avatar,
		Status:      im_sessions.SessionStatusOk,
	}

	model.DB.Save(&session)

}
