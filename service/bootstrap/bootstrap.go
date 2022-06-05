/**
  @author:panliang
  @data:2022/5/27
  @note
**/
package bootstrap

import (
	"github.com/gin-gonic/gin"
	"im-services/config"
	"im-services/pkg/coroutine_poll"
	"im-services/pkg/logger"
	"im-services/pkg/model"
	"im-services/router"
	"im-services/service/client"
)

// 启动服务方法
func Start() {

	r := gin.Default()

	setUpLogger()

	model.InitDb()

	coroutine_poll.ConnectPool()

	go client.ImManager.Start()

	setRoute(r)

	gin.SetMode(config.Conf.Server.Mode)

	r.Run(config.Conf.Server.Listen)
}

// 初始化日志方法
func setUpLogger() {
	logger.InitLogger(
		config.Conf.Log.FileName,
		config.Conf.Log.MaxSize,
		config.Conf.Log.MaxBackup,
		config.Conf.Log.MaxAge,
		config.Conf.Log.Compress,
		config.Conf.Log.Type,
		config.Conf.Log.Level,
	)
}

// 注册路由方法
func setRoute(r *gin.Engine) {
	router.RegisterApiRoutes(r)
	router.RegisterWsRouters(r)
}
