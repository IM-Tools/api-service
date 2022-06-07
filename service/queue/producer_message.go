/**
  @author:panliang
  @data:2022/6/6
  @note
**/
package queue

import "im-services/pkg/kafka"

type MessageProducerQueueInterface interface {
	MessageProducer(msg []byte) //消息生产者方法
}

type MessageProducerQueue struct {
}

var (
	offlinePrivateTopic = "offline_private_message" //离线私人消息频道
	offlineGroupTopic   = "offline_group_message"   //离线私人消息频道
)

// 消息消费 mag消息 channelType：1.私人消息 2.群聊消息
func OfflineMessageSaveQueue(msg []byte, channelType int) {
	var topic string
	if channelType == 1 {
		topic = offlinePrivateTopic
	} else {
		topic = offlineGroupTopic
	}
	producer := new(kafka.Producer)
	producer.InitProducer(topic)
	producer.SendMessage(msg)
}
