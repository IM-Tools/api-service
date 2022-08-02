package nsq_queue

import (
	"fmt"
	"im-services/internal/config"
	"im-services/internal/service/queue"
	"im-services/pkg/nsq"
)

func ConsumersInit() {
	err := nsq.NewConsumers(queue.OfflinePrivateTopic, "channel-aa", config.Conf.Nsq.LookupHost)
	if err != nil {
		fmt.Println("new nsq consumer failed", err)
		return
	}

	select {}
}
