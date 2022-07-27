package router

import (
	"github.com/gin-gonic/gin"
	"im-services/app/middleware"
	"im-services/app/service/handler"
)

// RegisterWsRouters 注册websocket路由
func RegisterWsRouters(router *gin.Engine) {

	WsService := new(handler.WsService)

	ws := router.Group("/im").Use(middleware.Auth()).Use(middleware.Cors())
	{
		ws.GET("/connect", WsService.Connect)
	}

}
