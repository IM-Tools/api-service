package message

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"im-services/internal/api/requests"
	"im-services/internal/api/services"
	"im-services/internal/enum"
	"im-services/internal/helpers"
	"im-services/internal/models/im_messages"
	"im-services/internal/models/user"
	"im-services/pkg/date"
	"im-services/pkg/model"
	"im-services/pkg/response"
	"net/http"
)

type MessageHandler struct {
}

func (m *MessageHandler) Index(cxt *gin.Context) {

	id := cxt.MustGet("id")
	page := helpers.StringToInt(cxt.DefaultQuery("page", "1"))
	toId := cxt.Query("to_id")
	pageSize := helpers.StringToInt(cxt.DefaultQuery("pageSize", "20"))

	var list []im_messages.ImMessages

	query := model.DB.Table("im_messages").
		Where("(form_id=? and to_id=?) or (form_id=? and to_id=?)", id, toId, toId, id).
		Order("created_at desc")

	var users user.ImUsers

	model.DB.Table("im_users").Where("id=?", toId).First(&users)

	var total int64
	query.Count(&total)

	if result := query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&list); result.RowsAffected == 0 {
		response.SuccessResponse(gin.H{
			"list": struct {
			}{},
			"mate": gin.H{
				"pageSize": pageSize,
				"page":     page,
				"total":    total,
			}}, http.StatusOK).ToJson(cxt)
		return
	}

	for key := range list {
		list[key].Users.ID = users.ID
		list[key].Users.Name = users.Name
		list[key].Users.Email = users.Email
		list[key].Users.Avatar = users.Avatar
	}

	response.SuccessResponse(gin.H{
		"list": list,
		"mate": gin.H{
			"pageSize": pageSize,
			"page":     page,
			"total":    total,
		}}, http.StatusOK).ToJson(cxt)
	return

}

func (m *MessageHandler) RecallMessage(cxt *gin.Context) {
	response.SuccessResponse().ToJson(cxt)
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

	var count int64
	model.DB.Table("im_friends").
		Where("to_id=? and form_id=?", id, params.ToID).
		Count(&count)

	if count == 0 {
		response.FailResponse(enum.WsNotFriend, "非好友关系,不能聊天...").ToJson(cxt)
		return
	}

	// 消息投递
	var messages services.ImMessageService
	ok, msg := messages.SendPrivateMessage(params)
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
