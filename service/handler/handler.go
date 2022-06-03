/**
  @author:panliang
  @data:2022/5/27
  @note
**/
package handler

import (
	"Im-Push-Services/pkg/ws"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WsService struct {
}

func (*WsService) Connect(cxt *gin.Context) {

	_, err := ws.App(cxt.Writer, cxt.Request)

	// 升级失败 返回服务器 500

	if err != nil {
		http.Error(cxt.Writer, cxt.Errors.String(), http.StatusInternalServerError)
	}

	// 注册客户端

	// 监听读写

}
