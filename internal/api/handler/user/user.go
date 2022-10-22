package user

import (
	"github.com/gin-gonic/gin"
	"im-services/internal/dao/friend_dao"
	"im-services/internal/dao/group_dao"
	"im-services/internal/enum"
	"im-services/internal/models/user"
	"im-services/pkg/model"
	"im-services/pkg/response"
)

type UsersHandler struct {
}

var (
	friendDao friend_dao.FriendDao
	groupDao  group_dao.GroupDao
)

// @BasePath /api

// PingExample godoc
// @Summary user/:id 获取用户信息
// @Schemes
// @Description 获取用户信息
// @Tags 用户
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer "
// @Produce json
// @Success 200 {object} response.JsonResponse{data=UserDetails} "ok"
// @Router /user/:id [get]
func (u *UsersHandler) Info(cxt *gin.Context) {
	var person Person
	if err := cxt.ShouldBindUri(&person); err != nil {
		response.FailResponse(enum.ParamError, err.Error()).ToJson(cxt)
		return
	}
	var users UserDetails

	if result := model.DB.Model(&user.ImUsers{}).
		Where("id=?", person.ID).
		First(&users); result.RowsAffected == 0 {
		response.ErrorResponse(enum.ParamError, "用户不存在").ToJson(cxt)
		return
	}
	response.SuccessResponse(users).ToJson(cxt)
	return

}

// @BasePath /api

// PingExample godoc
// @Summary user/:id 获取通讯录列表
// @Schemes
// @Description 获取通讯录列表
// @Tags 通讯录
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer "
// @Produce json
// @Success 200 {object} response.JsonResponse{data=UserDetails} "ok"
// @Router /user/:id [get]
func (*UsersHandler) AddressList(cxt *gin.Context) {

	userId := cxt.MustGet("id")

	err, lists := friendDao.GetFriendLists(userId)

	if err != nil {
		response.FailResponse(enum.ParamError, "获取用户列表失败").ToJson(cxt)
		return
	}
	groups := groupDao.GetGroupList(userId)

	response.SuccessResponse(lists).ToJson(cxt)

	response.SuccessResponse(gin.H{
		"friends": lists,
		"groups":  groups,
	}).ToJson(cxt)
}
