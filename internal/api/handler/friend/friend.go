package friend

import (
	"github.com/gin-gonic/gin"
	"im-services/internal/api/handler"
	"im-services/internal/dao/friend_dao"
	"im-services/internal/enum"
	"im-services/internal/helpers"
	"im-services/internal/service/dispatch"
	"im-services/pkg/response"
)

type FriendHandler struct {
}

// @BasePath /api

// PingExample godoc
// @Summary friends 获取好友列表
// @Schemes
// @Description 获取好友列表
// @Tags 好友
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer "
// @Produce json
// @Success 200 {object} response.JsonResponse{data=[]im_friends.ImFriends} "ok"
// @Router /friends/ [get]
func (friend FriendHandler) Index(cxt *gin.Context) {
	id := cxt.MustGet("id")

	var friendDao friend_dao.FriendDao

	err, lists := friendDao.GetFriendLists(id)

	if err != nil {
		response.SuccessResponse().ToJson(cxt)
		return
	}
	response.SuccessResponse(lists).ToJson(cxt)
	return

}

// @BasePath /api

// PingExample godoc
// @Summary friends/:id 获取好友详情
// @Schemes
// @Description 获取好友详情
// @Tags 好友
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer "
// @Param id path int true "ID"
// @Produce json
// @Success 200 {object} response.JsonResponse{data=im_friends.ImFriends} "ok"
// @Router /friends/:id [get]
func (friend FriendHandler) Show(cxt *gin.Context) {

	err, person := handler.GetPersonId(cxt)
	if err != nil {
		response.FailResponse(enum.ParamError, err.Error()).ToJson(cxt)
		return
	}

	var friendDao friend_dao.FriendDao

	err, lists := friendDao.GetFriends(person.ID)

	if err != nil {
		response.SuccessResponse().ToJson(cxt)
		return
	}
	response.SuccessResponse(&lists).ToJson(cxt)
	return
}

// @BasePath /api

// PingExample godoc
// @Summary friends/:id 删除好友
// @Schemes
// @Description 删除好友
// @Tags 好友
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer "
// @Param id path int true "好友ID"
// @Produce json
// @Success 200 {object} response.JsonResponse{} "ok"
// @Router /friends/:id [delete]
func (friend FriendHandler) Delete(cxt *gin.Context) {
	err, person := handler.GetPersonId(cxt)
	if err != nil {
		response.FailResponse(enum.ParamError, err.Error()).ToJson(cxt)
		return
	}
	var friendDao friend_dao.FriendDao

	errs := friendDao.DelFriends(person.ID, cxt.MustGet("id"))
	if errs != nil {
		response.FailResponse(enum.ParamError, errs.Error()).ToJson(cxt)
		return
	}
	response.SuccessResponse().ToJson(cxt)
	return
}

// @BasePath /api

// PingExample godoc
// @Summary friends/:id 获取好友在线状态
// @Schemes
// @Description 获取好友在线状态
// @Tags 好友
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer "
// @Param id path int true "ID"
// @Produce json
// @Success 200 {object} response.JsonResponse{data=UserStatus} "ok"
// @Router /friends/status/:id [get]
func (friend FriendHandler) GetUserStatus(cxt *gin.Context) {

	err, person := handler.GetPersonId(cxt)
	if err != nil {
		response.FailResponse(enum.ParamError, err.Error()).ToJson(cxt)
		return
	}

	var _dispatch dispatch.DispatchService
	ok, _ := _dispatch.IsDispatchNode(person.ID)

	if ok {
		response.SuccessResponse(&UserStatus{
			Status: enum.WsUserOnline,
			Id:     helpers.StringToInt(person.ID),
		}).ToJson(cxt)
		return
	}

	response.SuccessResponse(&UserStatus{
		Status: enum.WsUserOffline,
		Id:     helpers.StringToInt(person.ID),
	}).ToJson(cxt)
	return

}
