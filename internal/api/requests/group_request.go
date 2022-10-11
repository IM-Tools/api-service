package requests

type CreateGroupRequest struct {
	UserId     int64    `json:"user_id"`
	Name       string   `json:"name" validate:"required,min=2,max=20"`  //群名称
	Info       string   `json:"info" validate:"required,min=2,max=255"` //群介绍
	Avatar     string   `json:"avatar" validate:"required"`             //群头像
	Password   string   `json:"password"`
	Theme      string   `json:"theme" validate:"required"`
	IsPwd      int      `json:"is_pwd"`
	SelectUser []string `form:"select_user[]"`
}
