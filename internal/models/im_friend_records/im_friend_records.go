package im_friend_records

type ImFriendRecords struct {
	Id          int64   `gorm:"column:id;primaryKey" json:"id"`
	FormId      int64   `gorm:"column:form_id" json:"form_id"`
	ToId        int64   `gorm:"column:to_id" json:"to_id"`
	Status      int     `gorm:"column:status" json:"status"` //0 等待通过 1 已通过 2 已拒绝
	CreatedAt   string  `gorm:"column:created_at" json:"created_at"`
	Information string  `gorm:"column:information" json:"information"` //请求信息
	Users       ImUsers `gorm:"foreignKey:FormId;references:Id" json:"users"`
}

type ImUsers struct {
	Id     int64  `gorm:"column:id;primaryKey" json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

const (
	WaitingStatus = 0
	ThroughStatus = 1
	DownStatus    = 2
)
