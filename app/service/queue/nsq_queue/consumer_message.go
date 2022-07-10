/**
  @author:panliang
  @data:2022/6/7
  @note
**/
package nsq_queue

import (
	"fmt"
	"im-services/app/service/queue"
	"im-services/pkg/nsq"
)

var (
	//addr = config.Conf.Nsqe.Host + ":" + config.Conf.Nsqe.LookupdPort
	addr = "127.0.0.1:4161"
)

func ConsumersInit() {
	fmt.Println(addr)
	err := nsq.NewConsumers(queue.OfflinePrivateTopic, "channel-aa", addr)
	if err != nil {
		fmt.Println("new nsq consumer failed", err)
		return
	}

	select {}
}
