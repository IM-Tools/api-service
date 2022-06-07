/**
  @author:panliang
  @data:2022/6/7
  @note
**/
package nsq_queue

import (
	"im-services/pkg/logger"
	"im-services/pkg/nsq"
	"im-services/service/queue"
)

var (
	ProducerQueue MessageProducerQueue
)

type MessageProducerQueue struct {
}

func (messageQueue *MessageProducerQueue) SendMessage(msg []byte) {
	err := nsq.PublishMessage(queue.OfflinePrivateTopic, msg)
	if err != nil {
		logger.Logger.Info(err.Error())
	}
}
