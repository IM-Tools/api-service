package session

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"im-services/internal/api/handler"
	"im-services/internal/api/requests"
	"im-services/internal/dao/session_dao"
	"im-services/internal/enum"
	"im-services/internal/helpers"
	"im-services/internal/models/im_sessions"
	"im-services/pkg/model"
	"im-services/pkg/response"
)

type SessionHandler struct {
}

func (s SessionHandler) Index(cxt *gin.Context) {
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

func (s SessionHandler) Store(cxt *gin.Context) {

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
		response.SuccessResponse(session).ToJson(cxt)
		return
	}

	var sessionDao session_dao.SessionDao

	sessionDao.CreateSession(helpers.InterfaceToInt64(id), params.Id, params.Type)

	response.SuccessResponse(session).ToJson(cxt)

	return

}

type Person struct {
	ID string `uri:"id" binding:"required"`
}

func (s SessionHandler) Update(cxt *gin.Context) {
	err, person := handler.GetPersonId(cxt)
	if err != nil {
		response.FailResponse(enum.ParamError, err.Error()).ToJson(cxt)
		return
	}
	params := requests.SessionUpdate{
		TopStatus: helpers.StringToInt(cxt.PostForm("top_status")),
		Note:      cxt.PostForm("note"),
	}
	errs := validator.New().Struct(params)
	if errs != nil {
		response.FailResponse(enum.ParamError, errs.Error()).ToJson(cxt)
		return
	}

	model.DB.Model(&im_sessions.ImSessions{}).Where("id", person.ID).Updates(&params)

	response.SuccessResponse().ToJson(cxt)
	return

}

func (s SessionHandler) Delete(cxt *gin.Context) {

	err, person := handler.GetPersonId(cxt)
	if err != nil {
		response.FailResponse(enum.ParamError, err.Error()).ToJson(cxt)
		return
	}

	model.DB.Delete(&im_sessions.ImSessions{}, person.ID)

	response.SuccessResponse().ToJson(cxt)

	return
}
