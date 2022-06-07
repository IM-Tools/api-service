/**
  @author:panliang
  @data:2022/6/7
  @note
**/
package nsq_queue

import (
	"fmt"
	"im-services/pkg/nsq"
	"im-services/service/queue"
)

func ConsumersInit() {
	addr := "127.0.0.1:4161"
	err := nsq.NewConsumers(queue.OfflinePrivateTopic, "channel-aa", addr)
	if err != nil {
		fmt.Println("new nsq consumer failed", err)
		return
	}

	select {}
}
