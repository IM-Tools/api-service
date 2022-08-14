package friend

type ImFriendList struct {
	Id          int64  `gorm:"column:id" json:"id"`
	FormId      int64  `gorm:"column:form_id" json:"form_id"`
	ToId        int64  `gorm:"column:to_id" json:"to_id"`
	Status      int    `gorm:"column:status" json:"status"` //0 等待通过 1 已通过 2 已拒绝
	CreatedAt   string `gorm:"column:created_at" json:"created_at"`
	Information string `gorm:"column:information" json:"information"` //请求信息

}
type UserStatus struct {
	Status int `json:"status"` // 0 未在线 1 在线
	Id     int `json:"id"`     // 用户id
}

type Person struct {
	ID string `uri:"id" binding:"required"`
}

type Params struct {
	ToId string `uri:"to_id" binding:"required"`
}
