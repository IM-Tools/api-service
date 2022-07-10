/**
  @author:panliang
  @data:2022/5/27
  @note
**/
package handler

import (
	"github.com/gin-gonic/gin"
	client3 "im-services/app/api/service/client"
	"im-services/app/api/service/dispatch"
	"im-services/app/helpers"
	"im-services/pkg/ws"
	"net/http"
)

type WsService struct {
}

const (
	tourists_role = 1 // 游客
	user_role     = 2 // 用户
)

func (*WsService) Connect(cxt *gin.Context) {

	conn, err := ws.App(cxt.Writer, cxt.Request)

	// 升级失败 返回服务器 500

	if err != nil {
		http.Error(cxt.Writer, cxt.Errors.String(), http.StatusInternalServerError)
		return
	}

	var dService dispatch.DispatchService

	// 用户id
	id := helpers.InterfaceToInt64(cxt.MustGet("id"))
	uid := helpers.InterfaceToString(cxt.MustGet("uid"))
	dService.SetDispatchNode(helpers.Int64ToString(id))
	// 创建客户端
	client := client3.NewClient(helpers.Int64ToString(id), uid, user_role, conn)
	// 注册客户端
	client3.ImManager.Register <- client
	// 监听读写

	go client.Read()

	go client.Write()
}