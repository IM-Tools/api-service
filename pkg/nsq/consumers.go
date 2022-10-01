package nsq

import (
	"fmt"
	"im-services/internal/service/dao"

	"github.com/nsqio/go-nsq"
)

var (
	OfflineMessageSave = new(dao.OfflineMessageDao)
)

type Handler struct {
}

func (m *Handler) HandleMessage(msg *nsq.Message) (err error) {
	OfflineMessageSave.PrivateOfflineMessageSave(string(msg.Body))
	return

}

func (m *Handler) HandleGroupMessage(msg *nsq.Message) (err error) {
	OfflineMessageSave.PrivateOfflineMessageSave(string(msg.Body))
	return

}

func NewGroupConsumers(t string, c string, addr string) error {
	conf := nsq.NewConfig()
	nc, err := nsq.NewConsumer(t, c, conf)
	if err != nil {
		fmt.Println("create consumer failed err ", err)
		return err
	}
	consumer := &Handler{}
	nc.AddHandler(consumer)

	if err := nc.ConnectToNSQLookupd(addr); err != nil {
		return err
	}
	return nil
}

func NewConsumers(t string, c string, addr string) error {
	conf := nsq.NewConfig()
	nc, err := nsq.NewConsumer(t, c, conf)
	if err != nil {
		fmt.Println("create consumer failed err ", err)
		return err
	}
	consumer := &Handler{}
	nc.AddHandler(consumer)

	if err := nc.ConnectToNSQLookupd(addr); err != nil {
		return err
	}
	return nil
}
