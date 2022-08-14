package message

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"im-services/internal/api/requests"
	"im-services/internal/api/services"
	"im-services/internal/dao/friend_dao"
	"im-services/internal/enum"
	"im-services/internal/helpers"
	"im-services/internal/models/im_messages"
	"im-services/internal/models/user"
	"im-services/pkg/date"
	"im-services/pkg/model"
	"im-services/pkg/response"
	"net/http"
	"sort"
)

type MessageHandler struct {
}

var (
	messagesServices services.ImMessageService
	friend           friend_dao.FriendDao
)

func (m *MessageHandler) Index(cxt *gin.Context) {

	id := cxt.MustGet("id")
	page := cxt.Query("page")
	toId := cxt.Query("to_id")
	pageSize := helpers.StringToInt(cxt.DefaultQuery("pageSize", "50"))

	var list []im_messages.ImMessages

	query := model.DB.Table("im_messages").
		Where("(form_id=? and to_id=?) or (form_id=? and to_id=?)", id, toId, toId, id).
		Order("created_at desc")

	var users user.ImUsers

	model.DB.Table("im_users").Where("id=?", toId).First(&users)

	if len(page) > 0 {
		query = query.Where("id>?", page)
	}

	if result := query.Limit(pageSize).Find(&list); result.RowsAffected == 0 {
		response.SuccessResponse(gin.H{
			"list": struct {
			}{},
			"mate": gin.H{
				"pageSize": pageSize,
				"page":     page,
			}}, http.StatusOK).ToJson(cxt)
		return
	}

	SortByMessage(list, users)
	response.SuccessResponse(gin.H{
		"list": list,
		"mate": gin.H{
			"pageSize": pageSize,
			"page":     page,
		}}, http.StatusOK).ToJson(cxt)
	return

}
func SortByMessage(list []im_messages.ImMessages, users user.ImUsers) {
	sort.Slice(list, func(i, j int) bool {
		list[i].Users.ID = users.ID
		list[i].Users.Name = users.Name
		list[i].Users.Email = users.Email
		list[i].Users.Avatar = users.Avatar
		return list[i].Id < list[j].Id
	})
}
func (m *MessageHandler) RecallMessage(cxt *gin.Context) {
	response.SuccessResponse().ToJson(cxt)
	return
}

func (m *MessageHandler) SendVideoMessage(cxt *gin.Context) {

	id := cxt.MustGet("id")
	toId := cxt.PostForm("to_id")

	if !friend.IsFriends(id, toId) {
		response.FailResponse(enum.WsNotFriend, "非好友关系,不能聊天...").ToJson(cxt)
		return
	}
	var users user.ImUsers
	model.DB.Table("im_users").Where("id=?", id).First(&users)

	params := requests.VideoMessageRequest{
		MsgCode:  enum.VideoChantMessage,
		FormID:   helpers.InterfaceToInt64(id),
		ToID:     helpers.StringToInt64(toId),
		Message:  "视频请求...",
		SendTime: date.NewDate(),
		Users: requests.Users{
			Email:  users.Email,
			Name:   users.Name,
			Avatar: users.Avatar,
		},
	}
	ok := messagesServices.SendVideoMessage(params)
	if !ok {
		response.FailResponse(http.StatusInternalServerError, "用户不在线").ToJson(cxt)
		return
	}
	response.SuccessResponse(params).ToJson(cxt)
	return

}

func (m *MessageHandler) SendPrivateMessage(cxt *gin.Context) {

	id := cxt.MustGet("id")
	params := requests.PrivateMessageRequest{
		MsgId:       date.TimeUnixNano(),
		MsgCode:     enum.WsChantMessage,
		MsgClientId: helpers.StringToInt64(cxt.PostForm("msg_client_id")),
		FormID:      helpers.InterfaceToInt64(id),
		ToID:        helpers.StringToInt64(cxt.PostForm("to_id")),
		ChannelType: 1,
		MsgType:     helpers.StringToInt(cxt.PostForm("msg_type")),
		Message:     cxt.PostForm("message"),
		SendTime:    date.NewDate(),
		Data:        cxt.PostForm("data"),
	}

	errs := validator.New().Struct(params)

	if errs != nil {
		response.FailResponse(enum.ParamError, errs.Error()).ToJson(cxt)
		return
	}

	if !friend.IsFriends(id, params.ToID) {
		response.FailResponse(enum.WsNotFriend, "非好友关系,不能聊天...").ToJson(cxt)
		return
	}

	// 消息投递
	ok, msg := messagesServices.SendPrivateMessage(params)
	if !ok {

		response.FailResponse(http.StatusInternalServerError, msg).ToJson(cxt)
		return
	}

	message := im_messages.ImMessages{
		Msg:       params.Message,
		FormId:    params.FormID,
		ToId:      params.ToID,
		CreatedAt: params.SendTime,
		IsRead:    0,
		MsgType:   params.MsgType,
		Status:    1,
		Data:      helpers.InterfaceToString(params.Data),
	}

	model.DB.Save(&message)

	response.SuccessResponse(params).ToJson(cxt)
	return
}
