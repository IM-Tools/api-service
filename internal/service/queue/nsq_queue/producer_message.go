package nsq_queue

import (
	"im-services/internal/service/queue"
	"im-services/pkg/logger"
	"im-services/pkg/nsq"
)

var (
	ProducerQueue MessageProducerQueue
)

type MessageProducerQueue struct {
}

// 发送私聊消息到消息中间件
func (messageQueue *MessageProducerQueue) SendMessage(msg []byte) {
	err := nsq.PublishMessage(queue.OfflinePrivateTopic, msg)
	if err != nil {
		logger.Logger.Info(err.Error())
	}
}

// 发送群聊消息到消息中间件
func (messageQueue *MessageProducerQueue) SendGroupMessage(msg []byte) {
	err := nsq.PublishMessage(queue.OfflineGroupTopic, msg)
	if err != nil {
		logger.Logger.Info(err.Error())
	}
}
