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
