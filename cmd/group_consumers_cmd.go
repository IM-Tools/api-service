package cmd

import (
	"github.com/spf13/cobra"
	"im-services/internal/service/queue/nsq_queue"
	"im-services/pkg/console"
)

var GroupConsumers = &cobra.Command{
	Use:   "group_consumer",
	Short: "启动群离线消费者",
	Run:   runGroupConsumers,
	Args:  cobra.NoArgs,
}

func runGroupConsumers(cmd *cobra.Command, args []string) {
	for {
		console.Success("开始群聊消费")
		nsq_queue.ConsumersGroupMessageInit()
	}
}
