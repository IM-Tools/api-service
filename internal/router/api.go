package router

import (
	"im-services/docs"
	"im-services/internal/api/handler/auth"
	"im-services/internal/api/handler/cloud"
	"im-services/internal/api/handler/friend"
	"im-services/internal/api/handler/group"
	"im-services/internal/api/handler/message"
	"im-services/internal/api/handler/session"
	"im-services/internal/api/handler/user"
	"im-services/internal/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RegisterApiRoutes 注册api路由
func RegisterApiRoutes(router *gin.Engine) {

	var api *gin.RouterGroup
	docs.SwaggerInfo.BasePath = "/api"
	router.Use(middleware.Cors())
	api = router.Group("/api")
	// 登录
	authGroup := api.Group("/auth")
	{
		login := new(auth.AuthHandler)
		oauth := new(auth.OAuthHandler)

		authGroup.POST("/login", login.Login)                 //登录
		authGroup.POST("/registered", login.Registered)       //注册
		authGroup.POST("/sendEmailCode", login.SendEmailCode) //发送注册邮件
		authGroup.GET("/githubLogin", oauth.GithubOAuth)      //github登录
	}

	// 用户
	userGroup := api.Group("/user")
	{
		users := new(user.UsersHandler)
		userGroup.GET("/:id", users.Info) //获取用户信息
	}

	// 会话
	sessionGroup := api.Group("/sessions").Use(middleware.Auth())
	{
		sessions := new(session.SessionHandler)

		sessionGroup.GET("/", sessions.Index)        // 获取会话列表
		sessionGroup.POST("/", sessions.Store)       // 添加会话
		sessionGroup.PUT("/:id", sessions.Update)    // 更新会话
		sessionGroup.DELETE("/:id", sessions.Delete) // 移除会话

	}
	// 好友
	friendGroup := api.Group("/friends").Use(middleware.Auth())
	{
		friends := new(friend.FriendHandler)
		friendRecords := new(friend.FriendRecordHandler)

		friendGroup.GET("/", friends.Index)        //获取好友列表
		friendGroup.GET("/:id", friends.Show)      //获取好友详情信息
		friendGroup.DELETE("/:id", friends.Delete) //删除好友
		friendGroup.GET("/status/:id", friends.GetUserStatus)
		friendGroup.POST("/record", friendRecords.Store)       //发送好友请求
		friendGroup.GET("/record", friendRecords.Index)        //获取好友申请记录列表
		friendGroup.PUT("/record", friendRecords.Update)       //同意好友请求
		friendGroup.GET("/userQuery", friendRecords.UserQuery) //非好友用户查询

	}

	// 消息
	messageGroup := api.Group("/messages").Use(middleware.Auth())
	{
		messages := new(message.MessageHandler)
		groupMessages := new(message.GroupMessageHandler)
		messageGroup.GET("/", messages.Index)            //获取私聊消息列表
		messageGroup.GET("/groups", groupMessages.Index) //获取群聊消息列表

		messageGroup.POST("/private", messages.SendMessage)    // 发送私聊消息
		messageGroup.POST("/group", messages.SendMessage)      // 发送私聊消息
		messageGroup.POST("/video", messages.SendVideoMessage) // 发送视频请求
		messageGroup.POST("/recall", messages.RecallMessage)   // 消息撤回

	}

	// 群聊

	chatGroup := api.Group("/groups").Use(middleware.Auth())
	{
		groups := new(group.GroupHandler)
		chatGroup.POST("/store", groups.Store)             //创建群组
		chatGroup.POST("/applyJoin/:id", groups.ApplyJoin) //加入群组
		chatGroup.GET("/list", groups.Index)               //获取群组列表
		chatGroup.GET("/users/:id", groups.GetUsers)       //获取群成员信息
		chatGroup.DELETE("/:id", groups.Logout)            //退出群聊

	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	Clouds := new(cloud.QiNiuHandler)
	api.POST("/upload/file", Clouds.UploadFile).Use(middleware.Auth())
}
