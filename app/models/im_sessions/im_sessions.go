/**
  @author:panliang
  @data:2022/6/30
  @note
**/
package im_sessions

type ImSessions struct {
	Id          int64  `gorm:"column:id" json:"id"`
	MId         int    `gorm:"column:m_id" json:"m_id"`
	FId         int    `gorm:"column:f_id" json:"f_id"`
	CreatedAt   int64  `gorm:"column:created_at" json:"created_at"`
	TopStatus   int8   `gorm:"column:top_status" json:"top_status"` //0.否 1.是
	TopTime     int64  `gorm:"column:top_time" json:"top_time"`
	Note        string `gorm:"column:note" json:"note"`                 //备注
	ChannelType int8   `gorm:"column:channel_type" json:"channel_type"` //0.单聊 1.群聊
	Name        string `gorm:"column:name" json:"name"`                 //会话名称
	Avatar      string `gorm:"column:avatar" json:"avatar"`             //会话头像
	Status      int8   `gorm:"column:status" json:"status"`             //会话状态 0.正常 1.禁用
}
