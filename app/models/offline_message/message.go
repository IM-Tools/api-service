/**
  @author:panliang
  @data:2022/6/6
  @note
**/
package offline_message

type ImOfflineMessages struct {
	Id        int64  `gorm:"column:id;primaryKey" json:"id"`
	ReceiveId int64  `gorm:"column:receive_id" json:"receive_id"` //读取消息用户id
	Message   string `gorm:"column:message" json:"message"`       //消息体
	SendTime  int    `gorm:"column:send_time" json:"send_time"`   //消息接收时间
	Status    int    `gorm:"column:status" json:"status"`         //消息状态 0.未推送 1.已推送
}
