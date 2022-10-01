package requests

type CreateGroupRequest struct {
	OwnerId   int64  `json:"owner_id"`
	Name      string `json:"name" validate:"required,min=6,max=20"`  //群名称
	Info      string `json:"info" validate:"required,min=2,max=255"` //群介绍
	Avatar    string `json:"avatar" validate:"required"`             //群头像
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
