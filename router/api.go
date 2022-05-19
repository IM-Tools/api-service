/**
  @author:panliang
  @data:2022/5/16
  @note
**/
package router

import (
	"github.com/gin-gonic/gin"
)

// 注册api路由
func RegisterApiRoutes(router *gin.Engine) {

	startCors(router)

	api := router.Group("/api")
	{
		auth := api.Group("auth")
		{
			auth.POST("login")
		}
	}

}
