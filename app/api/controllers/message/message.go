/**
  @author:panliang
  @data:2022/7/7
  @note
**/
package message

import (
	"github.com/gin-gonic/gin"
	"im-services/app/helpers"
	"im-services/app/models/im_messages"
	"im-services/app/models/user"
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
