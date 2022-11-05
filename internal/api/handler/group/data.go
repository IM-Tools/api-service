package group

import "im-services/internal/models/im_group_users"

type GroupsDate struct {
	Groups ImGroups                      `json:"groups"`
	Users  []im_group_users.ImGroupUsers `json:"group_users"`
}

type SelectUser struct {
	SelectUser []string `form:"select_user[]"`
}

type GroupUsers struct {
	UserId string `json:"user_id"`
}

type ImGroups struct {
	Id        int64  `gorm:"column:id" json:"id"`                 //群聊id
	UserId    int64  `gorm:"column:user_id" json:"user_id"`       //创建者
	Name      string `gorm:"column:name" json:"name"`             //群聊名称
	CreatedAt string `gorm:"column:created_at" json:"created_at"` //添加时间
	Info      string `gorm:"column:info" json:"info"`             //群聊描述
	Avatar    string `gorm:"column:avatar" json:"avatar"`         //群聊头像
}
