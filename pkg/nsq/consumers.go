package nsq

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"im-services/app/service/dao"
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
