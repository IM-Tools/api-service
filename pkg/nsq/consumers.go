package nsq

import (
	"fmt"
	"im-services/internal/service/dao"

	"github.com/nsqio/go-nsq"
)

var (
	OfflineMessageSave = new(dao.OfflineMessageDao)
)

type PrivateHandler struct {
}

type GroupHandler struct {
}

func (m *PrivateHandler) HandleMessage(msg *nsq.Message) (err error) {
	OfflineMessageSave.PrivateOfflineMessageSave(string(msg.Body))
	return

}

func (group *GroupHandler) HandleMessage(msg *nsq.Message) (err error) {
	OfflineMessageSave.GroupOfflineMessageSave(string(msg.Body))
	return

}

func NewConsumers(t string, c string, addr string) error {

	conf := nsq.NewConfig()
	nc, err := nsq.NewConsumer(t, c, conf)
	if err != nil {
		fmt.Println("create consumer failed err ", err)
		return err
	}
	consumer := &PrivateHandler{}
	nc.AddHandler(consumer)

	if err := nc.ConnectToNSQLookupd(addr); err != nil {
		return err
	}
	return nil
}

func NewGroupConsumers(t string, c string, addr string) error {

	conf := nsq.NewConfig()
	nc, err := nsq.NewConsumer(t, c, conf)
	if err != nil {
		fmt.Println("create consumer failed err ", err)
		return err
	}
	consumer := &GroupHandler{}
	nc.AddHandler(consumer)

	if err := nc.ConnectToNSQLookupd(addr); err != nil {
		return err
	}
	return nil
}
