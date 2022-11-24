package session

import (
	"im-services/internal/api/handler"
	"im-services/internal/api/requests"
	"im-services/internal/dao/session_dao"
	"im-services/internal/enum"
	"im-services/internal/helpers"
	"im-services/internal/models/im_sessions"
	"im-services/pkg/model"
	"im-services/pkg/response"

	"github.com/gin-gonic/gin"
)

type SessionHandler struct {
}

// @BasePath /api

// PingExample godoc
// @Summary sessions/ 获取会话列表
// @Schemes
// @Description 获取会话列表
// @Tags 会话
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer "
// @Produce json
// @Success 200 {object} response.JsonResponse{data=[]im_sessions.ImSessions} "ok"
// @Router /sessions/ [get]
func (s SessionHandler) Index(cxt *gin.Context) {
	id := cxt.MustGet("id")

	var list []im_sessions.ImSessions

	if result := model.DB.Model(&im_sessions.ImSessions{}).
		Preload("Users").
		Preload("Groups").
		Where("form_id=? and status=0", id).
		Order("top_status desc").
		Find(&list); result.RowsAffected == 0 {
		response.SuccessResponse().ToJson(cxt)

		return
	}
	response.SuccessResponse(list).ToJson(cxt)

	return
}

// @BasePath /api

// PingExample godoc
// @Summary sessions/ 添加会话
// @Schemes
// @Description 添加会话
// @Tags 会话
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer "
// @Param type formData int true "会话类型"
// @Param id formData int true "聊天对象id"
// @Produce json
// @Success 200 {object} response.JsonResponse{data=[]im_sessions.ImSessions} "ok"
// @Router /sessions/ [post]
func (s SessionHandler) Store(cxt *gin.Context) {

	id := cxt.MustGet("id")
	params := requests.SessionStore{
		Id:   helpers.StringToInt64(cxt.PostForm("id")),
		Type: helpers.StringToInt(cxt.PostForm("type")),
	}

	errs := requests.Validate(params)

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

	sessions := sessionDao.CreateSession(helpers.InterfaceToInt64(id), params.Id, params.Type)

	response.SuccessResponse(sessions).ToJson(cxt)

	return

}

type Person struct {
	ID string `uri:"id" binding:"required"`
}

// @BasePath /api

// PingExample godoc
// @Summary sessions/:id 更新会话
// @Schemes
// @Description 更新会话
// @Tags 会话
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer "
// @Param id path int true "ID"
// @Param top_status formData int true "置顶 0 取消 1置顶"
// @Param note formData int true "会话备注"
// @Produce json
// @Success 200 {object} response.JsonResponse{} "ok"
// @Router /sessions/:id [put]
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
	errs := requests.Validate(params)
	if errs != nil {
		response.FailResponse(enum.ParamError, errs.Error()).ToJson(cxt)
		return
	}

	model.DB.Model(&im_sessions.ImSessions{}).Where("id", person.ID).Updates(&params)

	response.SuccessResponse().ToJson(cxt)
	return

}

// @BasePath /api

// PingExample godoc
// @Summary sessions/:id 删除会话
// @Schemes
// @Description 删除会话
// @Tags 会话
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer "
// @Param id path int true "ID"
// @Produce json
// @Success 200 {object} response.JsonResponse{} "ok"
// @Router /sessions/:id [delete]
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
