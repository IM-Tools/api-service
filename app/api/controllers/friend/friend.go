/**
  @author:panliang
  @data:2022/7/3
  @note
**/
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

// 获取好友列表
func (friend FriendController) Index(cxt *gin.Context) {
	id := cxt.MustGet("id")

	var list im_friends.ImFriends

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

type UserStatus struct {
	Status int `json:"status"`
	Id     int `json:"id"`
}

type Person struct {
	ID string `uri:"id" binding:"required"`
}

func (friend FriendController) GetUserStatus(cxt *gin.Context) {
	var person Person
	if err := cxt.ShouldBindUri(&person); err != nil {
		response.FailResponse(enum.PARAMS_ERROR, err.Error()).ToJson(cxt)
		return
	}
	var _dispatch dispatch.DispatchService
	ok, _ := _dispatch.IsDispatchNode(person.ID)

	if ok {
		response.SuccessResponse(&UserStatus{
			Status: enum.WS_USER_ONLINE,
			Id:     helpers.StringToInt(person.ID),
		}).ToJson(cxt)
		return
	}

	response.SuccessResponse(&UserStatus{
		Status: enum.WS_USER_OFFLINE,
		Id:     helpers.StringToInt(person.ID),
	}).ToJson(cxt)
	return

}

type Params struct {
	ToId string `uri:"to_id" binding:"required"`
}
