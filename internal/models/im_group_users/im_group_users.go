package im_group_users

type ImGroupUsers struct {
	Id        int     `gorm:"column:id" json:"id"`
	UserId    int     `gorm:"column:user_id" json:"user_id"`
	CreatedAt string  `gorm:"column:created_at" json:"created_at"`
	GroupId   int     `gorm:"column:group_id" json:"group_id"`
	Remark    string  `gorm:"column:remark" json:"remark"`
	Avatar    string  `gorm:"column:avatar" json:"avatar"`
	Name      string  `gorm:"column:name" json:"name"`
	Users     ImUsers `gorm:"foreignKey:UserId;references:ID" json:"users"`
}

type ImUsers struct {
	ID            int    `gorm:"column:id;primaryKey" json:"id"`
	Name          string `gorm:"column:name" json:"name"`
	Email         string `gorm:"column:email" json:"email"`
	Avatar        string `gorm:"column:avatar" json:"avatar"`
	Status        int8   `gorm:"column:status" json:"status"`
	Bio           string `gorm:"column:bio" json:"bio"`
	Sex           int8   `gorm:"column:sex" json:"sex"`
	ClientType    int8   `gorm:"column:client_type" json:"client_type"`
	Age           int    `gorm:"column:age" json:"age"`
	LastLoginTime string `gorm:"column:last_login_time" json:"last_login_time"`
	Uid           string `gorm:"column:uid" json:"uid"`
}
