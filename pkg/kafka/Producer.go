package kafka

import (
	"fmt"
	"gopkg.in/Shopify/sarama.v1"
	config2 "im-services/internal/config"
	"im-services/pkg/logger"
)

type Producer struct {
	Producer   sarama.SyncProducer
	Topic      string //主题
	ProducerID int    //生产者Id
	MessageId  int
}

var (
	ProducerId = 1
)

func (p *Producer) InitProducer(topic string) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 连接kafka%s
	client, err := sarama.NewSyncProducer([]string{fmt.Sprintf("%s:%s", config2.Conf.Kafka.Host, config2.Conf.Kafka.Port)}, config)
	if err != nil {
		logger.Logger.Error("producer closed, err:" + err.Error())
		return
	}

	p.Producer = client
	p.Topic = topic
	p.ProducerID = ProducerId
	p.MessageId = 1

	ProducerId++
}

func (p *Producer) SendMessage(txt []byte) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = p.Topic
	msg.Value = sarama.StringEncoder(txt)
	pid, offset, err := p.Producer.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	logger.Logger.Info(fmt.Sprintf("ProducerID:%d pid:%v offset:%v msg:%s",
		p.ProducerID, pid, offset, txt))

	p.MessageId++
}

func (p *Producer) Close() {
	err := p.Producer.Close()

	if err != nil {
		fmt.Println(err)
	}
}
