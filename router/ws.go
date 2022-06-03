/**
  @author:panliang
  @data:2022/5/27
  @note
**/
package router

import (
	"Im-Push-Services/app/middleware"
	"Im-Push-Services/service/handler"
	"github.com/gin-gonic/gin"
)

func RegisterWsRouters(router *gin.Engine) {

	WsService := new(handler.WsService)

	ws := router.Group("/im").Use(middleware.Auth())
	{
		ws.GET("/connect", WsService.Connect)
	}

}
