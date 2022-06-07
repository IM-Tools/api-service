/**
  @author:panliang
  @data:2022/6/6
  @note
**/
package kafak

type MessageProducerQueueInterface interface {
	MessageProducer(msg []byte) //消息生产者方法
}

type MessageProducerQueue struct {
}

// 消息消费 mag消息 channelType：1.私人消息 2.群聊消息
//func OfflineMessageSaveQueue(msg []byte, channelType int) {
//	var topic string
//	if channelType == 1 {
//		topic = queue.OfflinePrivateTopic
//	} else {
//		topic = queue.OfflinePrivateTopic
//	}
//	ProducerQueue := new(MessageProducerQueue)
//}
