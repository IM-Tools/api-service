/**
  @author:panliang
  @data:2022/6/30
  @note
**/
package im_sessions

type ImSessions struct {
	Id          int64   `gorm:"column:id;primaryKey" json:"id"` //会话表
	FormId      int64   `gorm:"column:form_id" json:"form_id"`
	ToId        int64   `gorm:"column:to_id" json:"to_id"`
	CreatedAt   string  `gorm:"column:created_at" json:"created_at"`
	TopStatus   int     `gorm:"column:top_status" json:"top_status"` //0.否 1.是
	TopTime     string  `gorm:"column:top_time" json:"top_time"`
	Note        string  `gorm:"column:note" json:"note"`                 //备注
	ChannelType int     `gorm:"column:channel_type" json:"channel_type"` //0.单聊 1.群聊
	Name        string  `gorm:"column:name" json:"name"`                 //会话名称
	Avatar      string  `gorm:"column:avatar" json:"avatar"`             //会话头像
	Status      int     `gorm:"column:status" json:"status"`             //会话状态 0.正常 1.禁用
	Users       ImUsers `gorm:"foreignKey:ID;references:ToId"`
}

type ImUsers struct {
	//model.BaseModel
	ID            int64  `gorm:"column:id;foreignKey" json:"id"`
	Name          string `gorm:"column:name" json:"name"`
	Email         string `gorm:"column:email" json:"email"`
	Avatar        string `gorm:"column:avatar" json:"avatar"`           //头像
	Status        int8   `gorm:"column:status" json:"status"`           //0 离线 1 在线
	Bio           string `gorm:"column:bio" json:"bio"`                 //用户简介
	Sex           int8   `gorm:"column:sex" json:"sex"`                 //0 未知 1.男 2.女
	ClientType    int8   `gorm:"column:client_type" json:"client_type"` //1.web 2.pc 3.app
	Age           int    `gorm:"column:age" json:"age"`
	LastLoginTime string `gorm:"column:last_login_time" json:"last_login_time"` //最后登录时间
}

const (
	SessionStatusOk = 0
	SessionStatusNo = 1
	TopStatusOk     = 1
	TopStatus       = 0
	GroupType       = 2
	PrivateType     = 1
)
