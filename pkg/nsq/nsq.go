package nsq

import (
	"github.com/nsqio/go-nsq"
	"log"
)

var Producer *nsq.Producer

func NewProducer(addr string) (err error) {
	conf := nsq.NewConfig()
	Producer, err = nsq.NewProducer(addr, conf)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
