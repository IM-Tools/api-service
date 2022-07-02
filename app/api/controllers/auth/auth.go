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
	enum "im-services/app/api/ enum"
	requests2 "im-services/app/api/requests"
	"im-services/app/helpers"
	"im-services/app/models/user"
	"im-services/app/services"
	"im-services/config"
	"im-services/pkg/date"
	"im-services/pkg/hash"
	"im-services/pkg/jwt"
	"im-services/pkg/model"
	"im-services/pkg/redis"
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

// 登录
func (*AuthController) Login(cxt *gin.Context) {

	params := requests2.LoginForm{
		Email:    cxt.PostForm("email"),
		Password: cxt.PostForm("password"),
	}

	validate := requests2.ValidateTransInit()
	err := validate.Struct(params)

	if err != nil {
		response.FailResponse(http.StatusInternalServerError, requests2.GetError(err)).WriteTo(cxt)
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

// 注册
func (*AuthController) Registered(cxt *gin.Context) {
	params := requests2.RegisteredForm{
		Email:          cxt.PostForm("email"),
		Name:           cxt.PostForm("name"),
		Password:       cxt.PostForm("password"),
		Code:           cxt.PostForm("code"),
		PasswordRepeat: cxt.PostForm("password_repeat"),
	}

	err := validator.New().Struct(params)

	if err != nil {
		response.FailResponse(enum.PARAMS_ERROR, err.Error()).WriteTo(cxt)
		return
	}

	createdAt := date.NewDate()

	model.DB.Table("im_users").Create(&user.ImUsers{
		Email:     params.Email,
		Password:  hash.BcryptHash(params.Password),
		Name:      params.Name,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	})

	response.SuccessResponse().ToJson(cxt)
	return
}

// 发送邮件
func (*AuthController) SendRegisteredMail(cxt *gin.Context) {

	email := cxt.Query("email")

	ok, message := requests2.IsEmailExits(email, "im_users")
	if !ok {
		response.FailResponse(enum.PARAMS_ERROR, message).ToJson(cxt)
		return
	}

	code := helpers.CreateEmailCode()

	var emailService services.EmailService

	html := fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Im-Services注册邮件</title>
</head>
<style>
    .mail{
        margin: 0 auto;
        border-radius: 45px;
        height: 400px;
        padding: 10px;
        background-color: #CC9933;
        background: url("https://img-blog.csdnimg.cn/c32f12dfd48241babd35b15189dc5c78.png") no-repeat;
    }
    .code {
        color: #f6512b;
        font-weight: bold;
        font-size: 30px;
        padding: 2px;
    }
</style>
<body>
<div class="mail">
    <h3>您好:您正在注册im-services应用账号!</h3>
    <p>下面是您的验证码:</p>
        <p class="code">%s</p>
        <p>请注意查收!谢谢</p>
</div>
<h3>如果可以请给项目点个star～<a target="_blank" href="https://github.com/IM-Tools/Im-Services">项目地址</a> </h3>
</body>
</html>`, code)

	err := emailService.SendEmail(email, "欢迎👏注册Im Services账号,这是一封邮箱验证码的邮件!🎉🎉🎉", html)
	if err != nil {
		response.FailResponse(enum.API_ERROR, "邮件发送失败,请检查是否是可用邮箱").ToJson(cxt)
		return
	}

	redis.RedisDB.Set(email, code, time.Minute*5)

	response.SuccessResponse().ToJson(cxt)
	return

}
