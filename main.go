package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"im-services/cmd"
	cmd2 "im-services/cmd/cmd"
	"im-services/internal/config"
	"im-services/internal/service/bootstrap"
	"im-services/pkg/console"
	"os"
)

func init() {
	config.InitConfig("config.yaml")
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "im",
		Short: "Hugo is a very fast static site generator",
		Long: `A Fast and Flexible Static Site Generator built with
	           love by spf13 and friends in Go.
	           Complete documentation is available at http://hugo.spf13.com`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			bootstrap.LoadConfiguration()
		},
	}

	rootCmd.AddCommand(
		cmd.CmdServe,
	)

	cmd2.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("命令启动失败 %v: %s", os.Args, err.Error()))
	}

}
