/**
  @author:panliang
  @data:2022/5/16
  @note
**/
package router

import (
	"github.com/gin-gonic/gin"
	"im-services/app/api/controllers/auth"
	"im-services/app/api/controllers/friend"
	"im-services/app/api/controllers/message"
	"im-services/app/api/controllers/session"
	"im-services/app/middleware"
)

// 注册api路由
func RegisterApiRoutes(router *gin.Engine) {

	var api *gin.RouterGroup
	api = router.Group("/api")

	// 登录
	authGroup := api.Group("/auth").Use(middleware.Cors())
	{
		login := new(auth.AuthController)

		authGroup.POST("/login", login.Login)                 //登录
		authGroup.POST("/registered", login.Registered)       //注册
		authGroup.POST("/sendEmailCode", login.SendEmailCode) //发送注册邮件
	}

	// 会话
	sessionGroup := api.Group("/sessions").Use(middleware.Auth())
	{
		sessions := new(session.SessionController)

		sessionGroup.GET("/", sessions.Index)
		sessionGroup.POST("/", sessions.Store)

	}

	// 好友
	friendGroup := api.Group("/friends").Use(middleware.Auth())
	{
		friends := new(friend.FriendController)
		friendRecords := new(friend.FriendRecordController)

		friendGroup.GET("/", friends.Index)
		friendGroup.POST("/record", friendRecords.Store) //发送好友请求
		friendGroup.GET("/record", friendRecords.Index)  //发送好友请求
		friendGroup.PUT("/record", friendRecords.Update) //同意好友请求

	}

	// 消息
	messageGroup := api.Group("/messages").Use(middleware.Auth())
	{
		messages := new(message.MessageController)

		messageGroup.GET("/", messages.Index)
		messageGroup.POST("/private", messages.SendPrivateMessage) //发送私聊消息

	}
}
