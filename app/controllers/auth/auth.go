/**
  @author:panliang
  @data:2022/5/16
  @note
**/
package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"im-services/app/models/user"
	"im-services/app/requests"
	"im-services/config"
	"im-services/pkg/hash"
	"im-services/pkg/jwt"
	"im-services/pkg/model"
	"im-services/pkg/response"
	"net/http"
	"time"
)

type AuthController struct {
}

type loginResponse struct {
	ID         int64  `json:"id"`
	UID        int64  `json:"uid"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Token      string `json:"token"`
	ExpireTime int64  `json:"expire_time"`
	Ttl        int64  `json:"ttl"`
}

func (*AuthController) Login(cxt *gin.Context) {

	params := requests.LoginForm{
		Email:    cxt.PostForm("email"),
		Password: cxt.PostForm("password"),
	}

	validate := requests.ValidateTransInit()
	err := validate.Struct(params)

	if err != nil {
		response.FailResponse(http.StatusInternalServerError, requests.GetError(err)).WriteTo(cxt)
		return
	}

	var users user.ImUsers

	result := model.DB.Table("im_users").Where("email=?", params.Email).First(&users)

	if result.RowsAffected == 0 {
		response.FailResponse(http.StatusInternalServerError, "邮箱未注册").ToJson(cxt)
		return
	}

	if !hash.BcryptCheck(params.Password, users.Password) {
		response.FailResponse(http.StatusInternalServerError, "密码错误").ToJson(cxt)
		return
	}

	ttl := config.Conf.JWT.Ttl
	expireAtTime := time.Now().Unix() + ttl
	token := jwt.NewJWT().IssueToken(
		users.ID,
		users.Uid,
		users.Name,
		users.Email,
		expireAtTime,
	)

	fmt.Println(users.ID)
	response.SuccessResponse(&loginResponse{
		ID:         users.ID,
		UID:        users.Uid,
		Name:       users.Name,
		Email:      users.Email,
		ExpireTime: expireAtTime,
		Token:      token,
		Ttl:        ttl,
	}).WriteTo(cxt)

	return

}

func (*AuthController) registered(cxt *gin.Context) {
	params := requests.RegisteredForm{
		Email:          cxt.PostForm("email"),
		Name:           cxt.PostForm("name"),
		Password:       cxt.PostForm("password"),
		Code:           cxt.PostForm("code"),
		PasswordRepeat: cxt.PostForm("password_repeat"),
	}

	err := validator.New().Struct(params)

	if err != nil {
		response.FailResponse(http.StatusInternalServerError, err.Error()).WriteTo(cxt)
		return
	}

	model.DB.Table("im_users").Create(&user.ImUsers{
		Email:     params.Email,
		Password:  hash.BcryptHash(params.Password),
		Name:      params.Name,
		CreatedAt: time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"),
	})

	response.SuccessResponse().ToJson(cxt)
	return
}
