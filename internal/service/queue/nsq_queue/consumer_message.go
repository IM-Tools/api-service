package nsq_queue

import (
	"fmt"
	"im-services/internal/config"
	"im-services/internal/service/queue"
	"im-services/pkg/nsq"
)

const (
	ChannelOfflineTopic      = "channel-offline-private"
	ChannelGroupOfflineTopic = "channel-offline-group"
	ChannelNodeTopic         = "channel-node"
)

func ConsumersPrivateMessageInit() {
	err := nsq.NewConsumers(queue.OfflinePrivateTopic, ChannelOfflineTopic, config.Conf.Nsq.LookupHost)
	if err != nil {
		fmt.Println("new nsq consumer failed", err)
		return
	}
	select {}
}

func ConsumersGroupMessageInit() {
	err := nsq.NewGroupConsumers(queue.OfflineGroupTopic, ChannelGroupOfflineTopic, config.Conf.Nsq.LookupHost)
	if err != nil {
		fmt.Println("new nsq consumer failed", err)
		return
	}
	select {}
}

func NodeInit() {
	err := nsq.NewConsumers(queue.OfflinePrivateTopic, ChannelNodeTopic, config.Conf.Nsq.LookupHost)
	if err != nil {
		fmt.Println("new nsq consumer failed", err)
		return
	}
	select {}
}
