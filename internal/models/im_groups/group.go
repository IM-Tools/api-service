package im_groups

type ImGroups struct {
	Id        int64  `gorm:"column:id" json:"id"`                 //群聊id
	UserId    int64  `gorm:"column:user_id" json:"user_id"`       //创建者
	Name      string `gorm:"column:name" json:"name"`             //群聊名称
	CreatedAt int64  `gorm:"column:created_at" json:"created_at"` //添加时间
	Info      string `gorm:"column:info" json:"info"`             //群聊描述
	Avatar    string `gorm:"column:avatar" json:"avatar"`         //群聊头像
	Password  string `gorm:"column:password" json:"password"`     //密码
	IsPwd     int8   `gorm:"column:is_pwd" json:"is_pwd"`         //是否加密 0 否 1 是
	Hot       int    `gorm:"column:hot" json:"hot"`               //热度
}

var (
	IS_PWD_NO  = 0
	IS_PWD_YES = 1
)
