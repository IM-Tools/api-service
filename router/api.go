/**
  @author:panliang
  @data:2022/5/16
  @note
**/
package router

import (
	"github.com/gin-gonic/gin"
	"im-services/app/api/controllers/auth"
	"im-services/app/api/controllers/session"
	"im-services/app/middleware"
)

// 注册api路由
func RegisterApiRoutes(router *gin.Engine) {

	var api *gin.RouterGroup
	api = router.Group("/api")

	authGroup := api.Group("/auth").Use(middleware.Cors())
	{
		login := new(auth.AuthController)

		authGroup.POST("/login", login.Login)                           //登录
		authGroup.POST("/registered", login.Registered)                 //注册
		authGroup.POST("/sendRegisteredMail", login.SendRegisteredMail) //发送注册邮件
	}

	sessionGroup := api.Group("/session").Use(middleware.Auth())
	{
		sessions := new(session.SessionController)

		sessionGroup.GET("/", sessions.Index)

	}

}
