package main

import (
	"Im-Push-Services/cmd"
	"Im-Push-Services/config"
	"Im-Push-Services/pkg/logger"
	"Im-Push-Services/pkg/model"
	"Im-Push-Services/router"
	"github.com/gin-gonic/gin"
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
