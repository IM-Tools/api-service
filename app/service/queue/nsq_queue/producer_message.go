package nsq_queue

import (
	"im-services/app/service/queue"
	"im-services/pkg/logger"
	"im-services/pkg/nsq"
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
