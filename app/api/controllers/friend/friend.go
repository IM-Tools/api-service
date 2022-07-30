package friend

import (
	"github.com/gin-gonic/gin"
	"im-services/app/api/controllers"
	"im-services/app/dao/friend_dao"
	"im-services/app/enum"
	"im-services/app/helpers"
	"im-services/app/service/dispatch"
	"im-services/pkg/response"
)

type FriendController struct {
}

func (friend FriendController) Index(cxt *gin.Context) {
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

func (friend FriendController) Show(cxt *gin.Context) {

	err, person := controllers.GetPersonId(cxt)
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

func (friend FriendController) GetUserStatus(cxt *gin.Context) {

	err, person := controllers.GetPersonId(cxt)
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
