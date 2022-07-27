package session

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"im-services/app/api/requests"
	"im-services/app/dao/session_dao"
	"im-services/app/enum"
	"im-services/app/helpers"
	"im-services/app/models/im_sessions"
	"im-services/pkg/model"
	"im-services/pkg/response"
	"net/http"
)

type SessionController struct {
}

func (s SessionController) Index(cxt *gin.Context) {
	id := cxt.MustGet("id")

	var list []im_sessions.ImSessions

	if result := model.DB.Model(&im_sessions.ImSessions{}).
		Preload("Users").
		Where("form_id=? and status=0", id).
		Order("top_status desc").
		Find(&list); result.RowsAffected == 0 {
		response.SuccessResponse().ToJson(cxt)

		return
	}

	response.SuccessResponse(list).ToJson(cxt)

	return
}

func (s SessionController) Store(cxt *gin.Context) {

	id := cxt.MustGet("id")
	params := requests.SessionStore{
		Id:   helpers.StringToInt64(cxt.PostForm("id")),
		Type: helpers.StringToInt(cxt.PostForm("type")),
	}

	errs := validator.New().Struct(params)

	if errs != nil {
		response.FailResponse(enum.ParamError, errs.Error()).ToJson(cxt)
		return
	}

	var session im_sessions.ImSessions
	if result := model.DB.Table("im_sessions").Where("form_id=? and to_id=?", id, params.Id).First(&session); result.RowsAffected > 0 {
		response.FailResponse(http.StatusInternalServerError, "已存在会话").ToJson(cxt)
		return
	}

	var sessionDao session_dao.SessionDao

	sessionDao.CreateSession(helpers.InterfaceToInt64(id), params.Id, params.Type)

	response.SuccessResponse().ToJson(cxt)

	return

}

type Person struct {
	ID string `uri:"id" binding:"required"`
}

func (s SessionController) Update(cxt *gin.Context) {
	var person Person
	if err := cxt.ShouldBindUri(&person); err != nil {
		response.FailResponse(enum.ParamError, err.Error()).ToJson(cxt)
		return
	}

	response.SuccessResponse().ToJson(cxt)
	return

}

func (s SessionController) Delete(cxt *gin.Context) {
	var person Person
	if err := cxt.ShouldBindUri(&person); err != nil {
		response.FailResponse(enum.ParamError, err.Error()).ToJson(cxt)
		return
	}

	model.DB.Delete(&im_sessions.ImSessions{}, person.ID)

	response.SuccessResponse().ToJson(cxt)

	return
}
