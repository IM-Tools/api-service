/**
  @author:panliang
  @data:2022/5/17
  @note
**/
package user

import "im-services/pkg/model"

type ImUsers struct {
	//model.BaseModel
	ID            int64  `gorm:"column:id;primaryKey" json:"id"`
	Name          string `gorm:"column:name" json:"name"`
	Email         string `gorm:"column:email" json:"email"`
	Password      string `gorm:"column:password" json:"password"`
	CreatedAt     string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     string `gorm:"column:updated_at" json:"updated_at"`
	Avatar        string `gorm:"column:avatar" json:"avatar"`           //头像
	OauthId       string `gorm:"column:oauth_id" json:"oauth_id"`       //第三方id
	BoundOauth    int8   `gorm:"column:bound_oauth" json:"bound_oauth"` //1.github 2.gitee
	OauthType     int8   `gorm:"column:oauth_type" json:"oauth_type"`   //1.微博 2.github
	Status        int8   `gorm:"column:status" json:"status"`           //0 离线 1 在线
	Bio           string `gorm:"column:bio" json:"bio"`                 //用户简介
	Sex           int8   `gorm:"column:sex" json:"sex"`                 //0 未知 1.男 2.女
	ClientType    int8   `gorm:"column:client_type" json:"client_type"` //1.web 2.pc 3.app
	Age           int    `gorm:"column:age" json:"age"`
	LastLoginTime string `gorm:"column:last_login_time" json:"last_login_time"` //最后登录时间
	Uid           string `gorm:"column:uid" json:"uid"`                         //uid 关联
}

func IsEmailExists(email string) bool {
	if model.DB.Model(&ImUsers{}).Where("email=?", email).RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

func IsUserExits(email string, name string) (bool, string) {

	var user ImUsers

	result := model.DB.Table("im_users").Where("email=? or name =?", email, name).First(&user)

	if result.RowsAffected > 0 {
		if user.Email == email {
			return true, "email"
		}
		return true, "name"
	}
	return false, ""
}
