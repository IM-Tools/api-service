/**
  @author:panliang
  @data:2022/5/16
  @note
**/
package router

import (
	"Im-Push-Services/app/controllers/auth"
	"github.com/gin-gonic/gin"
)

// 注册api路由
func RegisterApiRoutes(router *gin.Engine) {

	startCors(router)
	var api *gin.RouterGroup
	api = router.Group("/api")

	authGroup := api.Group("/auth")
	{
		login := new(auth.AuthController)

		authGroup.POST("/login", login.Login)
	}

}
