/**
  @author:panliang
  @data:2022/7/6
  @note
**/
package friend

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"im-services/app/api/requests"
	"im-services/app/api/services"
	"im-services/app/dao/friend_dao"
	"im-services/app/dao/session_dao"
	"im-services/app/enum"
	"im-services/app/helpers"
	"im-services/app/models/im_friend_records"
	"im-services/app/models/im_friends"
	"im-services/app/models/user"
	"im-services/app/service/message"
	"im-services/pkg/date"
	"im-services/pkg/model"
	"im-services/pkg/response"
	"net/http"
)

type FriendRecordController struct {
}

// 获取好友请求列表
func (friend FriendRecordController) Index(cxt *gin.Context) {
	var list im_friend_records.ImFriendRecords
	id := cxt.MustGet("id")
	if result := model.DB.Model(&im_friend_records.ImFriendRecords{}).Preload("Users").
		Where("to_id=? or form_id=?", id, id).
		Order("created_at desc").Find(&list); result.RowsAffected == 0 {
		response.SuccessResponse().ToJson(cxt)
		return
	}

	response.SuccessResponse(list).ToJson(cxt)
	return

}

// 添加好友请求
func (friend FriendRecordController) Store(cxt *gin.Context) {
	id := cxt.MustGet("id")

	params := requests.CreateFriendRequest{
		ToId:        cxt.PostForm("to_id"),
		Information: cxt.PostForm("information"),
	}

	errs := validator.New().Struct(params)

	if errs != nil {
		response.ErrorResponse(enum.PARAMS_ERROR, errs.Error()).ToJson(cxt)
		return
	}
	var users user.ImUsers

	if result := model.DB.Table("im_users").Where("id=?", params.ToId).First(&users); result.RowsAffected == 0 {
		response.ErrorResponse(enum.PARAMS_ERROR, "用户不存在").ToJson(cxt)
		return
	}

	var friends im_friends.ImFriends

	if result := model.DB.Table("im_friends").Where("form_id=? and to_id=?", id, params.ToId).First(&friends); result.RowsAffected > 0 {
		response.ErrorResponse(enum.PARAMS_ERROR, "用户已经是好友关系了...").ToJson(cxt)
		return
	}

	records := im_friend_records.ImFriendRecords{
		FormId:      helpers.InterfaceToInt64(id),
		ToId:        helpers.StringToInt64(params.ToId),
		Status:      im_friend_records.WAITING_STATUS,
		CreatedAt:   date.NewDate(),
		Information: params.Information,
	}

	model.DB.Save(&records)

	var messageService services.ImMessageService

	var msg message.CreateFriendMessage

	msg.MsgCode = enum.WS_CREATE
	msg.ID = records.Id
	msg.ToID = records.ToId
	msg.FormId = records.FormId
	msg.Information = records.Information
	msg.CreatedAt = records.CreatedAt
	msg.Status = records.Status
	msg.Users.ID = users.ID
	msg.Users.Avatar = users.Avatar
	msg.Users.Name = users.Name

	messageService.SendFriendActionMessage(msg)

	response.SuccessResponse().ToJson(cxt)
	return
}

// 同意好友请求
func (friend FriendRecordController) Update(cxt *gin.Context) {
	id := cxt.MustGet("id")
	params := requests.UpdateFriendRequest{
		Status: helpers.StringToInt(cxt.PostForm("status")),
		ID:     cxt.PostForm("id"),
	}

	errs := validator.New().Struct(params)
	if errs != nil {
		response.ErrorResponse(enum.PARAMS_ERROR, errs.Error()).ToJson(cxt)
		return
	}

	var records im_friend_records.ImFriendRecords

	if result := model.DB.Table("im_friend_records").
		Where("id=? and to_id=? and status=0", params.ID, id).First(&records); result.RowsAffected == 0 {
		response.ErrorResponse(http.StatusInternalServerError, "数据不存在").ToJson(cxt)
		return
	}

	var users user.ImUsers

	model.DB.Table("im_users").Where("id=?", records.ToId).Find(&users)

	records.Status = params.Status

	model.DB.Updates(&records)

	var messageService services.ImMessageService

	var msg message.CreateFriendMessage
	var msgCode int

	if params.Status == 1 {
		msgCode = enum.WS_FRIEND_OK
		var friendDao friend_dao.FriendDao
		//添加好友关系
		friendDao.AgreeFriendRequest(records.FormId, records.ToId)
		friendDao.AgreeFriendRequest(records.ToId, records.FormId)

		// 添加会话关系
		var sessionDao session_dao.SessionDao
		sessionDao.CreateSession(records.FormId, records.ToId, 1)
		sessionDao.CreateSession(records.ToId, records.FormId, 1)

	} else {
		msgCode = enum.WS_FRIEND_ERROR
	}

	msg.MsgCode = msgCode
	msg.ID = records.Id
	msg.ToID = records.ToId
	msg.FormId = records.FormId
	msg.Information = records.Information
	msg.CreatedAt = records.CreatedAt
	msg.Status = records.Status
	msg.Users.ID = users.ID
	msg.Users.Avatar = users.Avatar
	msg.Users.Name = users.Name

	messageService.SendFriendActionMessage(msg)

	response.SuccessResponse().ToJson(cxt)
	return

}
