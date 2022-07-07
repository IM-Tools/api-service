/**
  @author:panliang
  @data:2022/6/8
  @note
**/
package im_friends

import "im-services/app/models"

type ImFriends struct {
	models.BaseModel
	Id        int64   `gorm:"column:id" json:"id"`
	FormId    int64   `gorm:"column:form_id" json:"form_id"`
	ToId      int64   `gorm:"column:to_id" json:"to_id"`
	CreatedAt string  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt string  `gorm:"column:created_at" json:"updated_at"`
	Note      string  `gorm:"column:note" json:"note"`
	TopTime   string  `gorm:"column:top_time" json:"top_time"`
	Status    int     `gorm:"column:status" json:"status"` //0.未置顶 1.已置顶
	Uid       string  `gorm:"column:uid" json:"uid"`
	Users     ImUsers `gorm:"foreignkey:ID;references:ToId"`
}

type ImUsers struct {
	ID            int64  `gorm:"column:id;primaryKey" json:"id"`
	Name          string `gorm:"column:name" json:"name"`
	Email         string `gorm:"column:email" json:"email"`
	Avatar        string `gorm:"column:avatar" json:"avatar"`           //头像
	Status        int8   `gorm:"column:status" json:"status"`           //0 离线 1 在线
	Bio           string `gorm:"column:bio" json:"bio"`                 //用户简介
	Sex           int8   `gorm:"column:sex" json:"sex"`                 //0 未知 1.男 2.女
	ClientType    int8   `gorm:"column:client_type" json:"client_type"` //1.web 2.pc 3.app
	Age           int    `gorm:"column:age" json:"age"`
	LastLoginTime string `gorm:"column:last_login_time" json:"last_login_time"` //最后登录时间
	Uid           string `gorm:"column:uid" json:"uid"`                         //uid 关联
}
