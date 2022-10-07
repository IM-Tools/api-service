package group

import (
	"fmt"
	"im-services/internal/api/handler"
	"im-services/internal/api/requests"
	"im-services/internal/dao/group_dao"
	"im-services/internal/enum"
	"im-services/internal/helpers"
	"im-services/internal/models/im_groups"
	"im-services/internal/service/group"
	"im-services/pkg/hash"
	"im-services/pkg/model"
	"im-services/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	groupDao group_dao.GroupDao
)

type GroupHandler struct {
}

// 获取群聊列表
func (*GroupHandler) Index(cxt *gin.Context) {
	var imGroupLists []im_groups.ImGroups
	if result := model.DB.Model(&im_groups.ImGroups{}).
		Order("hot desc").
		Find(&imGroupLists); result.RowsAffected > 0 {
		response.SuccessResponse(imGroupLists).ToJson(cxt)
		return
	}
	response.SuccessResponse().ToJson(cxt)
	return
}

// @BasePath /api

// PingExample godoc
// @Summary friends/record 創建一個聊天組
// @Schemes
// @Description 創建一個聊天組
// @Tags 群聊
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer "
// @Param name formData string true "群聊名稱"
// @Param info formData string true "群聊介紹"
// @Param name formData string true "群聊頭像"
// @Param name formData string true "群聊密碼"
// @Param name formData int true "是否需要密碼 0 否 1 是"
// @Param theme formData string true "群聊主題"
// @Produce json
// @Success 200 {object} response.JsonResponse{data=im_groups.ImGroups} "ok"
// @Router /groups/add [post]
func (*GroupHandler) Store(cxt *gin.Context) {
	id := cxt.MustGet("id")
	var selectUser SelectUser

	cxt.ShouldBind(&selectUser)
	selectUser.SelectUser = append(selectUser.SelectUser, helpers.InterfaceToInt64String(id))
	fmt.Println(selectUser.SelectUser)
	params := requests.CreateGroupRequest{
		UserId:     helpers.InterfaceToInt64(id),
		Name:       cxt.PostForm("name"),
		Info:       cxt.PostForm("info"),
		Avatar:     cxt.PostForm("avatar"),
		Password:   cxt.PostForm("password"),
		IsPwd:      helpers.StringToInt(cxt.PostForm("is_pwd")),
		Theme:      cxt.PostForm("theme"),
		SelectUser: selectUser.SelectUser,
	}

	errs := validator.New().Struct(params)
	if errs != nil {
		response.ErrorResponse(enum.ParamError, errs.Error()).ToJson(cxt)
		return
	}

	if params.IsPwd == im_groups.IS_PWD_YES {
		params.Password = hash.BcryptHash(params.Password)
	}

	err, imGroups := groupDao.CreateGroup(params)
	if err != nil {
		response.FailResponse(enum.ApiError, "创建群聊失败！").WriteTo(cxt)
		return
	}

	groupDao.CreateSelectGroupUser(selectUser.SelectUser, int(imGroups.Id), params.Avatar, params.Name)

	groups := group.NewGroup(imGroups)
	group.ImAppGroupGathers.SetGroups(groups)
	// 创建成功之后发送创建群聊消息 --

	response.SuccessResponse(groups).WriteTo(cxt)
	return
}

// @BasePath /api

// PingExample godoc
// @Summary groups/ApplyJoin/:id 用户申请入群
// @Schemes
// @Description 用户申请入群
// @Tags 群聊
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer "
// @Param name formData string false "群聊密码 判断该群是否需要密码 is_pwd 字段"
// @Produce json
// @Success 200 {object} response.JsonResponse{} "ok"
// @Router /groups/ApplyJoin/:id [post]
func (*GroupHandler) ApplyJoin(cxt *gin.Context) {
	id := cxt.MustGet("id")
	err, person := handler.GetPersonId(cxt)
	if err != nil {
		response.FailResponse(enum.ParamError, "参数错误！").WriteTo(cxt)
		return
	}
	var group im_groups.ImGroups
	if result := model.DB.Model(&im_groups.ImGroups{}).Where("id=?", person.ID).Find(&group); result.RowsAffected == 0 {
		response.FailResponse(enum.ParamError, "群聊不存在！").WriteTo(cxt)
		return
	}

	if groupDao.IsGroupsUser(id, person.ID) {
		response.FailResponse(enum.ParamError, "已经是群成员了~").WriteTo(cxt)
		return
	}
	if group.IsPwd == int8(im_groups.IS_PWD_YES) {
		if !hash.BcryptCheck(cxt.PostForm("password"), group.Password) {
			response.FailResponse(enum.ParamError, "入群密码错误~,请联系管理员邀请").WriteTo(cxt)
			return
		}
	}

	groupDao.CreateOneGroupUser(group, int(helpers.InterfaceToInt64(id)))

	response.SuccessResponse().WriteTo(cxt)
	return

	// todo 成功之后 发送入群消息

}

// @BasePath /api

// PingExample godoc
// @Summary groups/users/:id 获取群聊用户信息
// @Schemes
// @Description 获取群聊用户信息
// @Tags 群聊
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer "
// @Produce json
// @Success 200 {object} response.JsonResponse{data=GroupsDate} "ok"
// @Router /groups/users/:id [get]
func (*GroupHandler) GetUsers(cxt *gin.Context) {
	err, person := handler.GetPersonId(cxt)
	if err != nil {
		response.FailResponse(enum.ParamError, "参数错误！").WriteTo(cxt)
		return
	}
	var group ImGroups
	if result := model.DB.Model(&im_groups.ImGroups{}).Where("id=?", person.ID).Find(&group); result.RowsAffected == 0 {
		response.FailResponse(enum.ParamError, "群聊不存在！").WriteTo(cxt)
		return
	}
	response.SuccessResponse(&GroupsDate{
		Groups: group,
		Users:  groupDao.GetGroupUsers(person.ID),
	}).WriteTo(cxt)
	return
}
