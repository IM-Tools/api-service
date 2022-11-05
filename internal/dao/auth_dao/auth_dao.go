package auth_dao

import (
	"encoding/json"
	"fmt"
	"im-services/internal/dao/session_dao"
	"im-services/internal/helpers"
	"im-services/internal/models/user"
	"im-services/pkg/date"
	"im-services/pkg/hash"
	"im-services/pkg/model"
)

type AuthDao struct {
}

func (*AuthDao) CreateUser(email string, password string, name string) int64 {
	createdAt := date.NewDate()
	users := user.ImUsers{
		Email:         email,
		Password:      hash.BcryptHash(password),
		Name:          name,
		CreatedAt:     createdAt,
		UpdatedAt:     createdAt,
		Avatar:        fmt.Sprintf("https://api.multiavatar.com/Binx %s.png", name),
		LastLoginTime: createdAt,
		Uid:           helpers.GetUuid(),
		UserJson:      "{}",
		UserType:      1,
	}
	model.DB.Table("im_users").Create(&users)
	var sessionDao session_dao.SessionDao
	sessionDao.CreateSession(users.ID, 1, 1)
	sessionDao.CreateSession(1, users.ID, 1)
	return users.ID

}

func (*AuthDao) isOAuthExists(oauthId string) bool {
	var count int64
	model.DB.Table("im_users").Where("oauth_id=?", oauthId).Count(&count)
	if count > 0 {
		return true
	}
	return false
}

// 获取或创建第三方登录信息
func (auth *AuthDao) CreateOauthUser(userInfo map[string]interface{}, oAuth string) (err error, info user.ImUsers, isNew bool) {
	id := helpers.Float64ToString(userInfo["id"].(float64))
	var users user.ImUsers

	if len(id) > 0 {
		if result := model.DB.Table("im_users").Where(oAuth+"_id=?", id).First(&users); result.RowsAffected > 0 {
			return nil, users, false
		}
	}
	userByte, err := json.Marshal(userInfo)
	if err != nil {
		return err, users, true
	}

	switch oAuth {
	case "github":
		name := userInfo["login"].(string)
		email := userInfo["email"].(string)
		password := id + "password"
		createdAt := date.NewDate()
		users = user.ImUsers{
			Email:         email,
			Password:      hash.BcryptHash(password),
			Name:          name,
			CreatedAt:     createdAt,
			UpdatedAt:     createdAt,
			Avatar:        userInfo["avatar_url"].(string),
			LastLoginTime: createdAt,
			Bio:           userInfo["bio"].(string),
			Uid:           helpers.GetUuid(),
			UserJson:      string(userByte),
			GithubId:      id,
			Github:        1,
		}
	case "gitee":
		name := userInfo["name"].(string)
		email := userInfo["login"].(string)
		password := id + "password"
		createdAt := date.NewDate()
		users = user.ImUsers{
			Email:         email,
			Password:      hash.BcryptHash(password),
			Name:          name,
			CreatedAt:     createdAt,
			UpdatedAt:     createdAt,
			Avatar:        userInfo["avatar_url"].(string),
			LastLoginTime: createdAt,
			Bio:           userInfo["bio"].(string),
			Uid:           helpers.GetUuid(),
			UserJson:      string(userByte),
			GiteeId:       id,
			Gitee:         1,
			GiteeUrl:      userInfo["html_url"].(string),
		}
	}
	model.DB.Table("im_users").Create(&users)

	return nil, users, true

}
