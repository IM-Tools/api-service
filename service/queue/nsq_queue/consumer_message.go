/**
  @author:panliang
  @data:2022/6/7
  @note
**/
package nsq_queue

import (
	"fmt"
	"im-services/config"
	"im-services/pkg/nsq"
	"im-services/service/queue"
)

func ConsumersInit() {
	err := nsq.NewConsumers(queue.OfflinePrivateTopic, "channel-aa", config.Conf.Nsq.ConsumptionHost)
	if err != nil {
		fmt.Println("new nsq consumer failed", err)
		return
	}

	select {}
}
