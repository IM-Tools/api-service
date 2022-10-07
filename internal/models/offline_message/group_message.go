package offline_message

type ImGroupOfflineMessages struct {
	Id        int    `gorm:"column:id" json:"id"`
	Message   string `gorm:"column:message" json:"message"`     //消息体
	SendTime  int    `gorm:"column:send_time" json:"send_time"` //消息接收时间
	Status    int8   `gorm:"column:status" json:"status"`       //消息状态 0.未推送 1.已推送
	ReceiveId int    `gorm:"column:receive_id" json:"receive_id"`
}
