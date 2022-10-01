package user

import (
	"github.com/gin-gonic/gin"
	"im-services/internal/enum"
	"im-services/internal/models/user"
	"im-services/pkg/model"
	"im-services/pkg/response"
)

type UsersHandler struct {
}

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
