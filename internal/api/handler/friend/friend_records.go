package friend

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"im-services/internal/api/requests"
	"im-services/internal/api/services"
	"im-services/internal/dao/friend_dao"
	"im-services/internal/dao/session_dao"
	"im-services/internal/enum"
	"im-services/internal/helpers"
	"im-services/internal/models/im_friend_records"
	"im-services/internal/models/im_friends"
	"im-services/internal/models/user"
	"im-services/internal/service/client"
	"im-services/pkg/date"
	"im-services/pkg/model"
	"im-services/pkg/response"
	"net/http"
)

type FriendRecordHandler struct {
}

// @BasePath /api

// PingExample godoc
// @Summary friends/record 获取好友申请记录
// @Schemes
// @Description 获取好友申请记录
// @Tags 好友申请
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer "
// @Produce json
// @Success 200 {object} response.JsonResponse{data=[]im_friend_records.ImFriendRecords} "ok"
// @Router /friends/record [get]
func (friend *FriendRecordHandler) Index(cxt *gin.Context) {
	var list []im_friend_records.ImFriendRecords
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

// @BasePath /api

// PingExample godoc
// @Summary friends/record 发起好友申请
// @Schemes
// @Description 发起好友申请
// @Tags 好友申请
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer "
// @Param to_id formData string true "添加人id"
// @Param information formData array true "添加描述"
// @Produce json
// @Success 200 {object} response.JsonResponse{data=[]im_friend_records.ImFriendRecords} "ok"
// @Router /friends/record [post]
func (friend *FriendRecordHandler) Store(cxt *gin.Context) {
	id := cxt.MustGet("id")

	params := requests.CreateFriendRequest{
		ToId:        cxt.PostForm("to_id"),
		Information: cxt.PostForm("information"),
	}

	errs := validator.New().Struct(params)

	if errs != nil {
		response.ErrorResponse(enum.ParamError, errs.Error()).ToJson(cxt)
		return
	}
	var users user.ImUsers

	if result := model.DB.Table("im_users").Where("id=?", params.ToId).First(&users); result.RowsAffected == 0 {
		response.ErrorResponse(enum.ParamError, "用户不存在").ToJson(cxt)
		return
	}

	var record im_friend_records.ImFriendRecords
	if result := model.DB.Table("im_friend_records").Where("form_id=? and to_id=? and status=0", id, params.ToId).First(&record); result.RowsAffected > 0 {
		response.ErrorResponse(enum.ParamError, "请勿重复添加...").ToJson(cxt)
		return
	}

	var friends im_friends.ImFriends

	if result := model.DB.Table("im_friends").Where("form_id=? and to_id=?", id, params.ToId).First(&friends); result.RowsAffected > 0 {
		response.ErrorResponse(enum.ParamError, "用户已经是好友关系了...").ToJson(cxt)
		return
	}

	records := im_friend_records.ImFriendRecords{
		FormId:      helpers.InterfaceToInt64(id),
		ToId:        helpers.StringToInt64(params.ToId),
		Status:      im_friend_records.WaitingStatus,
		CreatedAt:   date.NewDate(),
		Information: params.Information,
	}

	model.DB.Save(&records)

	var messageService services.ImMessageService

	var msg client.CreateFriendMessage

	msg.MsgCode = enum.WsCreate
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

	records.Users.Name = users.Name
	records.Users.Id = users.ID
	records.Users.Avatar = users.Avatar
	response.SuccessResponse(records).ToJson(cxt)
	return
}

// @BasePath /api

// PingExample godoc
// @Summary record 同意或者拒绝请求
// @Schemes
// @Description 同意或者拒绝请求
// @Tags 好友申请
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer "
// @Param id formData string true "记录id"
// @Param status formData int true "状态 0等待 1同意 2拒绝"
// @Produce json
// @Success 200 {object} response.JsonResponse{} "ok"
// @Router /friends/record [put]
func (friend *FriendRecordHandler) Update(cxt *gin.Context) {
	id := cxt.MustGet("id")
	params := requests.UpdateFriendRequest{
		Status: helpers.StringToInt(cxt.PostForm("status")),
		ID:     cxt.PostForm("id"),
	}
	errs := validator.New().Struct(params)
	if errs != nil {
		response.ErrorResponse(enum.ParamError, errs.Error()).ToJson(cxt)
		return
	}
	var records im_friend_records.ImFriendRecords

	if result := model.DB.Table("im_friend_records").
		Where("id=? and status=0", params.ID).First(&records); result.RowsAffected == 0 {
		response.ErrorResponse(http.StatusInternalServerError, "数据不存在").ToJson(cxt)
		return
	}

	var friends im_friends.ImFriends

	if result := model.DB.Table("im_friends").Where("form_id=? and to_id=?", records.ToId, id).First(&friends); result.RowsAffected > 0 {
		response.ErrorResponse(enum.ParamError, "用户已经是好友关系了...").ToJson(cxt)
		return
	}

	var users user.ImUsers

	model.DB.Table("im_users").Where("id=?", id).Find(&users)

	records.Status = params.Status

	model.DB.Updates(&records)

	var messageService services.ImMessageService

	var msg client.CreateFriendMessage
	var msgCode int

	if params.Status == 1 {
		msgCode = enum.WsFriendOk
		var friendDao friend_dao.FriendDao
		//添加好友关系
		friendDao.AgreeFriendRequest(records.FormId, records.ToId)
		friendDao.AgreeFriendRequest(records.ToId, records.FormId)

		// 添加会话关系
		var sessionDao session_dao.SessionDao
		sessionDao.CreateSession(records.FormId, records.ToId, 1)
		sessionDao.CreateSession(records.ToId, records.FormId, 1)

	} else {
		msgCode = enum.WsFriendError
	}

	msg.MsgCode = msgCode
	msg.ID = records.Id
	msg.ToID = records.FormId
	msg.FormId = records.ToId
	msg.Information = records.Information
	msg.CreatedAt = records.CreatedAt
	msg.Status = records.Status
	msg.Users.ID = users.ID
	msg.Users.Avatar = users.Avatar
	msg.Users.Name = users.Name

	messageService.SendFriendActionMessage(msg)
	friends.Status = params.Status
	response.SuccessResponse(friends).WriteTo(cxt)
	return

}

// QueryUser 查询非好友用户
func (friend *FriendRecordHandler) UserQuery(cxt *gin.Context) {
	id := cxt.MustGet("id")
	params := requests.QueryUserRequest{
		Email: cxt.Query("email"),
	}
	errs := validator.New().Struct(params)
	if errs != nil {
		response.ErrorResponse(enum.ParamError, errs.Error()).ToJson(cxt)
		return
	}
	var friendDao friend_dao.FriendDao
	users := friendDao.GetNotFriendList(id, params.Email)
	response.SuccessResponse(users).ToJson(cxt)
	return
}
