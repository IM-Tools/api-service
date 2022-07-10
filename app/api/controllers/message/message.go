/**
  @author:panliang
  @data:2022/7/7
  @note
**/
package message

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"im-services/app/api/requests"
	"im-services/app/enum"
	"im-services/app/helpers"
	"im-services/app/models/im_messages"
	"im-services/app/models/user"
	"im-services/app/services"
	"im-services/pkg/date"
	"im-services/pkg/model"
	"im-services/pkg/response"
	"net/http"
)

type MessageController struct {
}

/**
获取消息列表
*/
func (m MessageController) Index(cxt *gin.Context) {

	id := cxt.MustGet("id")
	page := helpers.StringToInt(cxt.DefaultQuery("page", "1"))
	toId := cxt.Query("to_id")
	pageSize := helpers.StringToInt(cxt.DefaultQuery("pageSize", "20"))

	var list []im_messages.ImMessages

	query := model.DB.Table("im_messages").Where("form_id=? and to_id=?", id, toId).
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

	for key, _ := range list {
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

// 私聊消息投递

func (m MessageController) SendPrivateMessage(cxt *gin.Context) {

	id := cxt.MustGet("id")
	params := requests.PrivateMessageRequest{
		MsgId:       date.TimeUnixNano(),
		MsgCode:     http.StatusOK,
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
		response.FailResponse(enum.PARAMS_ERROR, errs.Error()).ToJson(cxt)
		return
	}

	var count int64
	model.DB.Table("im_friends").
		Where("to_id=? and form_id=?", id, params.ToID).
		Count(&count)

	if count == 0 {
		response.FailResponse(enum.WS_NOT_FRIEND, "非好友关系,不能聊天...").ToJson(cxt)
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