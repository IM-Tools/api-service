package group_invites

type ImGroupInvites struct {
	Id        int    `json:"id" form:"id"`
	UserId    int    `json:"user_id" form:"user_id"`       //分享人
	GroupId   int    `json:"group_id" form:"group_id"`     //群id
	Token     string `json:"token" form:"token"`           //入群加密口令
	CreatedAt int64  `json:"created_at" form:"created_at"` //分享时间
	EndTime   int    `json:"end_time" form:"end_time"`     //入群截止时间
	Avatar    string `json:"avatar" form:"avatar"`         //群头像
	Name      string `json:"name" form:"name"`             //群名称
}
