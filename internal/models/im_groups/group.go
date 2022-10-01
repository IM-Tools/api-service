package im_groups

type ImGroups struct {
	Id        int64  `gorm:"column:id" json:"id"`
	OwnerId   int64  `gorm:"column:owner_id" json:"owner_id"`
	Name      string `gorm:"column:name" json:"name"`
	CreatedAt int64  `gorm:"column:created_at" json:"created_at"`
	Info      string `gorm:"column:info" json:"info"`
	Avatar    string `gorm:"column:avatar" json:"avatar"`
	Number    int    `gorm:"column:number" json:"number"`
	Hot       int    `gorm:"column:hot" json:"hot"`
	UpdatedAt int64  `gorm:"column:updated_at" json:"updated_at"`
}
