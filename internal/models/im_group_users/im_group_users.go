package im_group_users

type ImGroupUsers struct {
	Id        int    `gorm:"column:id" json:"id"`
	UserId    int    `gorm:"column:user_id" json:"user_id"`
	CreatedAt int64  `gorm:"column:created_at" json:"created_at"`
	GroupId   int    `gorm:"column:group_id" json:"group_id"`
	Remark    string `gorm:"column:remark" json:"remark"`
	Avatar    string `gorm:"column:avatar" json:"avatar"`
	Name      string `gorm:"column:name" json:"name"`
}
