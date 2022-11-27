/**
  @author:panliang
  @data:2022/5/27
  @note
**/
package handler

import (
	"github.com/gin-gonic/gin"
	"im-services/internal/api/event"
	"im-services/internal/helpers"
	client2 "im-services/internal/service/client"
	"im-services/internal/service/dispatch"
	"im-services/pkg/ws"
	"net/http"
)

type WsService struct {
}

const (
	touristsRole = 1 // 游客
	userRole     = 2 // 用户
)

var (
	eventHandle event.EventHandle
)

func (*WsService) Connect(cxt *gin.Context) {

	conn, err := ws.App(cxt.Writer, cxt.Request)

	// 升级失败 返回服务器 500

	if err != nil {
		http.Error(cxt.Writer, cxt.Errors.String(), http.StatusInternalServerError)
		return
	}
	id := helpers.InterfaceToInt64(cxt.MustGet("id"))
	uid := helpers.InterfaceToString(cxt.MustGet("uid"))
	// 异地登录事件
	eventHandle.LogoutEvent(helpers.Int64ToString(id), cxt.Request.Header.Get("X-Forward-For"))

	var dService dispatch.DispatchService

	// 用户id
	dService.SetDispatchNode(helpers.Int64ToString(id))

	// 创建客户端

	client := client2.NewClient(helpers.Int64ToString(id), uid, userRole, conn)
	// 注册客户端
	client2.ImManager.Register <- client
	// 监听读写

	go client.Read()

	go client.Write()
}

func (*WsService) TouristsConnect(cxt *gin.Context) {
	conn, err := ws.App(cxt.Writer, cxt.Request)
	// 升级失败 返回服务器 500
	if err != nil {
		http.Error(cxt.Writer, cxt.Errors.String(), http.StatusInternalServerError)
		return
	}
	id := cxt.Query("token_id") // 生成规则 - ip+时间

	// 创建客户端
	client := client2.NewClient(id, "", touristsRole, conn)

	// 注册客户端

	client2.ImManager.Register <- client

	// 监听读写

	go client.Read()

	go client.Write()
}
