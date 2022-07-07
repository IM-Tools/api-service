/**
  @author:panliang
  @data:2022/7/7
  @note
**/
package im_messages

type ImMessages struct {
	Id        int     `gorm:"column:id" json:"id"`
	Msg       string  `gorm:"column:msg" json:"msg"`
	CreatedAt int64   `gorm:"column:created_at" json:"created_at"`
	FormId    int     `gorm:"column:form_id" json:"form_id"`
	ToId      int     `gorm:"column:to_id" json:"to_id"`
	IsRead    int8    `gorm:"column:is_read" json:"is_read"` //0 未读 1已读
	MsgType   int8    `gorm:"column:msg_type" json:"msg_type"`
	Status    int8    `gorm:"column:status" json:"status"`
	Data      string  `gorm:"column:data" json:"data"`
	Users     ImUsers `gorm:"foreignkey:ID;references:ToId"`
}

type ImUsers struct {
	ID     int64  `gorm:"column:id;primaryKey" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
	Email  string `gorm:"column:email" json:"email"`
	Avatar string `gorm:"column:avatar" json:"avatar"` //头像
}
