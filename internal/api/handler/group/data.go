package group

import "im-services/internal/models/im_group_users"

type ImGroups struct {
	Id        int64  `gorm:"column:id" json:"id"`                 //群聊id
	UserId    int64  `gorm:"column:user_id" json:"user_id"`       //创建者
	Name      string `gorm:"column:name" json:"name"`             //群聊名称
	CreatedAt string `gorm:"column:created_at" json:"created_at"` //添加时间
	Info      string `gorm:"column:info" json:"info"`             //群聊描述
	Avatar    string `gorm:"column:avatar" json:"avatar"`         //群聊头像
	IsPwd     int8   `gorm:"column:is_pwd" json:"is_pwd"`         //是否加密 0 否 1 是
	Hot       int    `gorm:"column:hot" json:"hot"`               //热度
}

type GroupsDate struct {
	Groups ImGroups                      `json:"groups"`
	Users  []im_group_users.ImGroupUsers `json:"group_users"`
}

type SelectUser struct {
	SelectUser []string `form:"select_user[]"`
}
