/**
  @author:panliang
  @data:2022/5/27
  @note
**/
package handler

import (
	"github.com/gin-gonic/gin"
	"im-services/pkg/helpers"
	"im-services/pkg/ws"
	wsClient "im-services/service/client"
	"net/http"
)

type WsService struct {
}

func (*WsService) Connect(cxt *gin.Context) {

	conn, err := ws.App(cxt.Writer, cxt.Request)

	// 升级失败 返回服务器 500

	if err != nil {
		http.Error(cxt.Writer, cxt.Errors.String(), http.StatusInternalServerError)
		return
	}

	// 用户id
	id := helpers.InterfaceToInt64(cxt.MustGet("id"))

	// 创建客户端
	client := wsClient.NewClient(id, conn)
	// 注册客户端
	wsClient.ImManager.Register <- client
	// 监听读写

	go client.Read()

	go client.Write()
}
