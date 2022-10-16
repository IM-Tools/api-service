/**
  @author:panliang
  @data:2022/7/7
  @note
**/
package im_messages

type ImMessages struct {
	Id        int64  `gorm:"column:id" json:"id"`
	Msg       string `gorm:"column:msg" json:"msg"`
	CreatedAt string `gorm:"column:created_at" json:"created_at"`
	FormId    int64  `gorm:"column:form_id" json:"form_id"`
	ToId      int64  `gorm:"column:to_id" json:"to_id"`
	IsRead    int    `gorm:"column:is_read" json:"is_read"`
	MsgType   int    `gorm:"column:msg_type" json:"msg_type"`
	Status    int    `gorm:"column:status" json:"status"`

	Data  string  `gorm:"column:data" json:"data"`
	Users ImUsers `gorm:"foreignkey:ID;references:ToId"`
}

type ImUsers struct {
	ID     int64  `gorm:"column:id;primaryKey" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
	Email  string `gorm:"column:email" json:"email"`
	Avatar string `gorm:"column:avatar" json:"avatar"` //头像
}

var (
	TEXT         = 1
	VOICE        = 2
	FILE         = 3
	IMAGE        = 4
	LOGOUT_GROUP = 5
	JOIN_GROUP   = 6
)
