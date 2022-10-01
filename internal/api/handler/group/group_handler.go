package group

import (
	"im-services/internal/api/requests"
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
// @Success 200 {object} response.JsonResponse{data=requests.CreateGroupRequest} "ok"
// @Router /groups/add [post]
func (*GroupHandler) Store(cxt *gin.Context) {
	id := cxt.MustGet("id")
	params := requests.CreateGroupRequest{
		UserId:   helpers.InterfaceToInt64(id),
		Name:     cxt.PostForm("name"),
		Info:     cxt.PostForm("info"),
		Avatar:   cxt.PostForm("avatar"),
		Password: cxt.PostForm("password"),
		IsPwd:    helpers.StringToInt(cxt.PostForm("is_pwd")),
		Theme:    cxt.PostForm("theme"),
	}
	errs := validator.New().Struct(params)
	if errs != nil {
		response.ErrorResponse(enum.ParamError, errs.Error()).ToJson(cxt)
		return
	}

	if params.IsPwd == im_groups.IS_PWD_YES {
		params.Password = hash.BcryptHash(params.Password)
	}

	var imGroups im_groups.ImGroups
	imGroups.UserId = params.UserId
	imGroups.Name = params.Name
	imGroups.Info = params.Info
	imGroups.Avatar = params.Avatar
	if model.DB.Model(&im_groups.ImGroups{}).Create(&imGroups).Error != nil {
		response.ErrorResponse(enum.ParamError, "创建群聊失败").ToJson(cxt)
		return
	}
	groups := group.NewGroup(imGroups)
	group.ImAppGroupGathers.SetGroups(groups)

	response.SuccessResponse(groups).WriteTo(cxt)
	return
}
