package group

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"im-services/internal/api/requests"
	"im-services/internal/enum"
	"im-services/internal/helpers"
	"im-services/internal/models/im_groups"
	"im-services/internal/service/group"
	"im-services/pkg/model"
	"im-services/pkg/response"
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

// 创建一个聊天组
func (*GroupHandler) Store(cxt *gin.Context) {
	id := cxt.MustGet("id")
	params := requests.CreateGroupRequest{
		OwnerId: helpers.InterfaceToInt64(id),
		Name:    cxt.PostForm("name"),
		Info:    cxt.PostForm("info"),
		Avatar:  cxt.PostForm("avatar"),
	}
	errs := validator.New().Struct(params)
	if errs != nil {
		response.ErrorResponse(enum.ParamError, errs.Error()).ToJson(cxt)
		return
	}
	var imGroups im_groups.ImGroups
	imGroups.OwnerId = params.OwnerId
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
