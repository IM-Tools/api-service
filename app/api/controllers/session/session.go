/**
  @author:panliang
  @data:2022/6/30
  @note
**/
package session

import (
	"github.com/gin-gonic/gin"
	"im-services/app/models/im_sessions"
	"im-services/pkg/model"
	"im-services/pkg/response"
)

type SessionController struct {
}

// 获取会话列表
func (session *SessionController) Index(cxt *gin.Context) {
	id := cxt.MustGet("id")

	var list im_sessions.ImSessions

	model.DB.Table("im_sessions").
		Where("m_id=? and status=0", id).
		Order("top_status desc").
		Find(&list)

	response.SuccessResponse(list).ToJson(cxt)

	return
}
