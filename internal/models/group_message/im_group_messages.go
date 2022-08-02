package group_message

type ImGroupMessages struct {
	Id              int64  `gorm:"column:id" form:"id"`
	Message         string `gorm:"column:message" form:"message"`                     //消息实体
	SendTime        int    `gorm:"column:send_time" form:"send_time"`                 //消息添加时间
	MessageId       int    `gorm:"column:message_id" form:"message_id"`               //服务端消息id
	ClientMessageId int    `gorm:"column:client_message_id" form:"client_message_id"` //客户端消息id
	FormId          int64  `gorm:"column:form_id" form:"form_id"`                     //消息发送者id
	GroupId         int64  `gorm:"column:group_id" form:"group_id"`                   //群聊id
}
