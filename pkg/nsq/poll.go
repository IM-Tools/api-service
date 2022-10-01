package nsq

import (
	"errors"
	"im-services/internal/config"
	"time"

	"github.com/nsqio/go-nsq"
	"github.com/silenceper/pool"
)

var (
	NsqProducerPool pool.Pool
)

func InitNewProducerPoll() error {
	factory := func() (interface{}, error) {
		producer, err := nsq.NewProducer(config.Conf.Nsq.NsqHost, nsq.NewConfig())
		if err != nil {
			return nil, err
		}
		return producer, nil
	}

	closeError := func(v interface{}) error {
		v.(*nsq.Producer).Stop()
		return nil
	}

	poolConfig := &pool.Config{
		InitialCap:  20,
		MaxIdle:     40,
		MaxCap:      50,
		Factory:     factory,
		Close:       closeError,
		IdleTimeout: 0 * time.Second,
	}
	var err error
	NsqProducerPool, err = pool.NewChannelPool(poolConfig)
	if err != nil {
		return errors.New("NewChannelPool init failed")
	}
	return err

}

// 私聊消息推送入topic
func PublishMessage(topic string, content []byte) error {
	nsqProducer, err := NsqProducerPool.Get()
	if err != nil {
		return err
	}
	defer NsqProducerPool.Put(nsqProducer)

	err = nsqProducer.(*nsq.Producer).Publish(topic, content)
	if err != nil {
		return err
	}
	return nil

}

// 群聊消息推送入topic
func PublishGroupMessage(topic string, content []byte) error {
	nsqProducer, err := NsqProducerPool.Get()
	if err != nil {
		return err
	}
	defer NsqProducerPool.Put(nsqProducer)

	err = nsqProducer.(*nsq.Producer).Publish(topic, content)
	if err != nil {
		return err
	}
	return nil

}

func Exit() {
	NsqProducerPool.Release()
}
