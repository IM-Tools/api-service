package router

import (
	"github.com/gin-gonic/gin"
	"im-services/app/api/controllers/auth"
	"im-services/app/api/controllers/friend"
	"im-services/app/api/controllers/message"
	"im-services/app/api/controllers/session"
	"im-services/app/api/controllers/user"
	"im-services/app/middleware"
)

// RegisterApiRoutes 注册api路由
func RegisterApiRoutes(router *gin.Engine) {

	var api *gin.RouterGroup

	api = router.Group("/api")
	api.Use(middleware.Cors())
	// 登录
	authGroup := api.Group("/auth")
	{
		login := new(auth.AuthController)

		authGroup.POST("/login", login.Login)                 //登录
		authGroup.POST("/registered", login.Registered)       //注册
		authGroup.POST("/sendEmailCode", login.SendEmailCode) //发送注册邮件
	}

	// 登录
	userGroup := api.Group("/user")
	{
		users := new(user.UsersController)

		userGroup.GET("/:id", users.Info) //获取用户信息

	}

	// 会话
	sessionGroup := api.Group("/sessions").Use(middleware.Auth())
	{
		sessions := new(session.SessionController)

		sessionGroup.GET("/", sessions.Index)        // 获取会话列表
		sessionGroup.POST("/", sessions.Store)       // 添加会话
		sessionGroup.PUT("/:id", sessions.Update)    // 更新会话
		sessionGroup.DELETE("/:id", sessions.Delete) // 移除会话

	}

	// 好友
	friendGroup := api.Group("/friends").Use(middleware.Auth())
	{
		friends := new(friend.FriendController)
		friendRecords := new(friend.FriendRecordController)

		friendGroup.GET("/", friends.Index)   //获取好友列表
		friendGroup.GET("/:id", friends.Show) //获取好友详情信息
		friendGroup.GET("/status/:id", friends.GetUserStatus)
		friendGroup.POST("/record", friendRecords.Store)       //发送好友请求
		friendGroup.GET("/record", friendRecords.Index)        //获取好友申请记录列表
		friendGroup.PUT("/record", friendRecords.Update)       //同意好友请求
		friendGroup.GET("/userQuery", friendRecords.UserQuery) //非好友用户查询

	}

	// 消息
	messageGroup := api.Group("/messages").Use(middleware.Auth())
	{
		messages := new(message.MessageController)
		messageGroup.GET("/", messages.Index)                      //获取私聊消息列表
		messageGroup.POST("/private", messages.SendPrivateMessage) // 发送私聊消息
		messageGroup.POST("/recall", messages.RecallMessage)       // 消息撤回

	}
}
