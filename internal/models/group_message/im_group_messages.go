package group_message

type ImGroupMessages struct {
	Id              int64   `gorm:"column:id" form:"id"`
	Message         string  `gorm:"column:message" form:"message"`                     //消息实体
	CreatedAt       string  `gorm:"column:created_at" form:"created_at"`               //添加时间
	Data            string  `gorm:"column:data" form:"data"`                           //自定义内容
	SendTime        int64   `gorm:"column:send_time" form:"send_time"`                 //消息添加时间
	MsgType         int     `gorm:"column:msg_type" form:"msg_type"`                   //消息添加时间
	MessageId       int64   `gorm:"column:message_id" form:"message_id"`               //服务端消息id
	ClientMessageId int64   `gorm:"column:client_message_id" form:"client_message_id"` //客户端消息id
	FormId          int64   `gorm:"column:form_id" form:"form_id"`                     //消息发送者id
	GroupId         int64   `gorm:"column:group_id" form:"group_id"`                   //群聊id
	Users           ImUsers `gorm:"foreignkey:ID;references:FormId"`                   //关联用户信息
}

type ImUsers struct {
	ID     int64  `gorm:"column:id;primaryKey" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
	Email  string `gorm:"column:email" json:"email"`
	Avatar string `gorm:"column:avatar" json:"avatar"` //头像
}
