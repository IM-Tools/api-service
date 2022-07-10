/**
  @author:panliang
  @data:2022/7/2
  @note
**/
package cmd

import (
	"github.com/spf13/cobra"
	"im-services/app/api/service/bootstrap"
)

var CmdServe = &cobra.Command{
	Use:   "run",
	Short: "启动im服务",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	Run: Run,
}

func Run(cmd *cobra.Command, args []string) {
	bootstrap.Start()
}
