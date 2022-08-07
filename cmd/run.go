package cmd

import (
	"github.com/spf13/cobra"
	"im-services/internal/service/bootstrap"
)

var AppCmdServe = &cobra.Command{
	Use:   "run",
	Short: "启动im服务",
	Long:  `启动im服务`,
	Run:   Run,
}

func Run(cmd *cobra.Command, args []string) {
	bootstrap.Start()
}
