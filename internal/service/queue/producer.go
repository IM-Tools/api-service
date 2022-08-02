package queue

var (
	OfflinePrivateTopic = "offline_private_message" //离线私人消息频道
	OfflineGroupTopic   = "offline_group_message"   //离线私人消息频道
)

type MessageProducerQueueInterface interface {
	SendMessage(msg []byte)
}
