package cmd

import (
	"github.com/spf13/cobra"
	"im-services/internal/service/queue/nsq_queue"
	"im-services/pkg/console"
)

var PrivateConsumers = &cobra.Command{
	Use:   "private_consumer",
	Short: "启动群离线消费者",
	Run:   runPrivateConsumers,
	Args:  cobra.NoArgs,
}

func runPrivateConsumers(cmd *cobra.Command, args []string) {

	//nsq_queue.ConsumersPrivateMessageInit()
	for {
		console.Success("开始私聊消费")
		nsq_queue.ConsumersPrivateMessageInit()
	}
}
