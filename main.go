package main

import (
	"github.com/gin-gonic/gin"
	"go-admin/cmd"
	"go-admin/config"
	"go-admin/pkg/model"
	"go-admin/router"
)

func init() {

	config.InitConfig("config.yaml")

}

func main() {
	cmd.Execute()

	gin.SetMode(config.Conf.Server.Mode)

	r := gin.Default()

	model.InitDb()

	router.RegisterApiRoutes(gin.Default())
	r.Run(config.Conf.Server.Listen)
}
