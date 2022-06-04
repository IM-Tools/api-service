package main

import (
	"github.com/gin-gonic/gin"
	"im-services/cmd"
	"im-services/config"
	"im-services/pkg/coroutine"
	"im-services/pkg/logger"
	"im-services/pkg/model"
	"im-services/router"
)

func init() {

	config.InitConfig("config.yaml")
	SetUpLogger()

}

func main() {
	cmd.Execute()

	gin.SetMode(config.Conf.Server.Mode)

	r := gin.Default()

	model.InitDb()
	
	coroutine.ConnectPool()

	router.RegisterApiRoutes(r)
	router.RegisterWsRouters(r)

	r.Run(config.Conf.Server.Listen)
}

func SetUpLogger() {
	logger.InitLogger(
		config.Conf.Log.Level,
		config.Conf.Log.MaxSize,
		config.Conf.Log.MaxBackup,
		config.Conf.Log.MaxAge,
		config.Conf.Log.Compress,
		config.Conf.Log.Type,
		config.Conf.Log.Level,
	)
}
