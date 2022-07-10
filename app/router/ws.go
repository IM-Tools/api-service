/**
  @author:panliang
  @data:2022/5/27
  @note
**/
package router

import (
	"github.com/gin-gonic/gin"
	"im-services/app/api/service/handler"
	"im-services/app/middleware"
)

// 注册websocket长链接路由
func RegisterWsRouters(router *gin.Engine) {

	WsService := new(handler.WsService)

	ws := router.Group("/im").Use(middleware.Auth()).Use(middleware.Cors())
	{
		ws.GET("/connect", WsService.Connect)
	}

}
