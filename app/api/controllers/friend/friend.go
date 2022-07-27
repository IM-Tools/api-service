package friend

import (
	"github.com/gin-gonic/gin"
	"im-services/app/enum"
	"im-services/app/helpers"
	"im-services/app/models/im_friends"
	"im-services/app/service/dispatch"
	"im-services/pkg/model"
	"im-services/pkg/response"
)

type FriendController struct {
}

func (friend FriendController) Index(cxt *gin.Context) {
	id := cxt.MustGet("id")

	var list []im_friends.ImFriends

	result := model.DB.Model(&im_friends.ImFriends{}).Preload("Users").
		Where("form_id=?", id).
		Order("status desc").
		Order("top_time desc").
		Find(&list)
	if result.RowsAffected == 0 {
		response.SuccessResponse().ToJson(cxt)
		return
	}

	response.SuccessResponse(list).ToJson(cxt)
	return

}

func (friend FriendController) GetUserStatus(cxt *gin.Context) {
	var person Person
	if err := cxt.ShouldBindUri(&person); err != nil {
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
