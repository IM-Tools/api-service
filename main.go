package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"im-services/cmd"
	AppCmd "im-services/cmd/cmd"
	"im-services/internal/config"
	"im-services/internal/service/bootstrap"
	"im-services/pkg/console"
	"os"
)

func init() {
	config.InitConfig("config.yaml")
}

// @title im-services  接口文档
// @version 1.0
// @description

// @contact.name im-services
// @contact.url
// @contact.email pltrueover@gmail.com

// @license.name MIT
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8000
// @BasePath /api
func main() {
	var rootCmd = &cobra.Command{
		Use:   "im",
		Short: "Hugo is a very fast static site generator",
		Long: `A Fast and Flexible Static Site Generator built with
	           love by spf13 and friends in Go.
	           Complete documentation is available at https://hugo.spf13.com`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			bootstrap.LoadConfiguration()
		},
	}

	rootCmd.AddCommand(
		cmd.AppCmdServe,
		cmd.GroupConsumers,
		cmd.PrivateConsumers,
	)

	AppCmd.RegisterDefaultCmd(rootCmd, cmd.AppCmdServe)
	AppCmd.RegisterGlobalFlags(rootCmd)
	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("命令启动失败 %v: %s", os.Args, err.Error()))
	}

}
