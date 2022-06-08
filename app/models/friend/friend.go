/**
  @author:panliang
  @data:2022/6/8
  @note
**/
package friend

type ImFriends struct {
	Id        int64  `gorm:"column:id" json:"id"`
	MId       int64  `gorm:"column:m_id" json:"m_id"`
	FId       int64  `gorm:"column:f_id" json:"f_id"`
	CreatedAt string `gorm:"column:created_at" json:"created_at"`
	Note      string `gorm:"column:note" json:"note"`
	TopTime   string `gorm:"column:top_time" json:"top_time"`
	Status    int    `gorm:"column:status" json:"status"` //0.未置顶 1.已置顶
}
