package router

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
)

var (
	login         auth.AuthHandler
	oauth         auth.OAuthHandler
	sessions      session.SessionHandler
	users         user.UsersHandler
	friends       friend.FriendHandler
	friendRecords friend.FriendRecordHandler
	messages      message.MessageHandler
	groupMessages message.GroupMessageHandler
	groups        group.GroupHandler
	clouds        cloud.QiNiuHandler
)

// RegisterApiRoutes 注册api路由
func RegisterApiRoutes(router *gin.Engine) {

	var api *gin.RouterGroup
	docs.SwaggerInfo.BasePath = "/api"
	router.Use(middleware.Cors())
	api = router.Group("/api")
	{
		// 登录
		authGroup := api.Group("/auth")
		{
			authGroup.POST("/login", login.Login)                 //登录
			authGroup.POST("/registered", login.Registered)       //注册
			authGroup.POST("/sendEmailCode", login.SendEmailCode) //发送注册邮件
			authGroup.GET("/githubLogin", oauth.GithubOAuth)      //github登录
		}

		api.Use(middleware.Auth())
		{
			// 用户
			api.GET("/user/:id", users.Info) //获取用户信息
			// 会话
			api.GET("/sessions", sessions.Index)         // 获取会话列表
			api.POST("/sessions", sessions.Store)        // 添加会话
			api.PUT("/sessions/:id", sessions.Update)    // 更新会话
			api.DELETE("/sessions/:id", sessions.Delete) // 移除会话

			// 好友

			api.Any("/friends", friends.Index)         //获取好友列表
			api.GET("/friends/:id", friends.Show)      //获取好友详情信息
			api.DELETE("/friends/:id", friends.Delete) //删除好友
			api.GET("/friends/status/:id", friends.GetUserStatus)
			api.POST("/friends/record", friendRecords.Store)       //发送好友请求
			api.GET("/friends/record", friendRecords.Index)        //获取好友申请记录列表
			api.PUT("/friends/record", friendRecords.Update)       //同意好友请求
			api.GET("/friends/userQuery", friendRecords.UserQuery) //非好友用户查询

			// 消息

			api.GET("/messages", messages.Index)             //获取私聊消息列表
			api.GET("/messages/groups", groupMessages.Index) //获取群聊消息列表

			api.POST("/messages/private", messages.SendMessage)    // 发送私聊消息
			api.POST("/messages/group", messages.SendMessage)      // 发送私聊消息
			api.POST("/messages/video", messages.SendVideoMessage) // 发送视频请求
			api.POST("/messages/recall", messages.RecallMessage)   // 消息撤回

			// 群聊

			api.POST("/groups/store", groups.Store)             //创建群组
			api.POST("/groups/applyJoin/:id", groups.ApplyJoin) //加入群组
			api.GET("/groups/list", groups.Index)               //获取群组列表
			api.GET("/groups/users/:id", groups.GetUsers)       //获取群成员信息
			api.DELETE("/groups/:id", groups.Logout)            //退出群聊

			api.POST("/upload/file", clouds.UploadFile).Use(middleware.Auth())
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
