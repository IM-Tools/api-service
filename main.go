package main

import (
	"im-services/cmd"
	"im-services/config"
	"im-services/service/bootstrap"
)

func init() {
	config.InitConfig("config.yaml")
}

func main() {
	cmd.Execute()
	bootstrap.Start()
}
