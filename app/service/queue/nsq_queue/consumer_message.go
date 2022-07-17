/**
  @author:panliang
  @data:2022/6/7
  @note
**/
package nsq_queue

import (
	"fmt"
	"im-services/app/service/queue"
	"im-services/config"
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
