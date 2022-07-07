/**
  @author:panliang
  @data:2022/7/3
  @note
**/
package friend

import (
	"github.com/gin-gonic/gin"
	"im-services/app/models/im_friends"
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

type Params struct {
	ToId string `uri:"to_id" binding:"required"`
}
