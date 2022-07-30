package user

import (
	"github.com/gin-gonic/gin"
	"im-services/app/enum"
	"im-services/app/models/user"
	"im-services/pkg/model"
	"im-services/pkg/response"
)

type UsersController struct {
}

func (u *UsersController) Info(cxt *gin.Context) {
	var person Person
	if err := cxt.ShouldBindUri(&person); err != nil {
		response.FailResponse(enum.ParamError, err.Error()).ToJson(cxt)
		return
	}
	var users UserDetails

	if result := model.DB.Model(&user.ImUsers{}).Where("id=?", person.ID).First(&users); result.RowsAffected == 0 {
		response.ErrorResponse(enum.ParamError, "用户不存在").ToJson(cxt)
		return
	}
	response.SuccessResponse(users).ToJson(cxt)
	return

}
