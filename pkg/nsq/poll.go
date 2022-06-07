/**
  @author:panliang
  @data:2022/6/7
  @note
**/
package nsq

import (
	"errors"
	"github.com/nsqio/go-nsq"
	"github.com/silenceper/pool"
	"time"
)

var NsqProducerPool pool.Pool

func InitNewProducerPoll() error {
	factory := func() (interface{}, error) {
		producer, err := nsq.NewProducer("0.0.0.0:4150", nsq.NewConfig())
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

func Exit() {
	NsqProducerPool.Release()
}
