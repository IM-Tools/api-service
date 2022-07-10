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
	"im-services/app/api/requests"
	"im-services/app/api/services"
	"im-services/app/enum"
	"im-services/app/helpers"
	"im-services/app/models/user"
	"im-services/config"
	"im-services/pkg/date"
	"im-services/pkg/hash"
	"im-services/pkg/jwt"
	"im-services/pkg/logger"
	"im-services/pkg/model"
	"im-services/pkg/response"
	"net/http"
	"time"
)

type AuthController struct {
}

type AuthControllerInterface interface {

	// 登录
	Login(cxt *gin.Context)

	// 注册
	Registered(cxt *gin.Context)

	// 发送邮件
	SendEmailCode(cxt *gin.Context)
}

type loginResponse struct {
	ID         int64  `json:"id"`
	UID        string `json:"uid"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
	Email      string `json:"email"`
	Token      string `json:"token"`
	ExpireTime int64  `json:"expire_time"`
	Ttl        int64  `json:"ttl"`
}

// 登录
func (*AuthController) Login(cxt *gin.Context) {

	params := requests.LoginForm{
		Email:    cxt.PostForm("email"),
		Password: cxt.PostForm("password"),
	}

	errs := validator.New().Struct(params)

	if errs != nil {
		response.FailResponse(http.StatusInternalServerError, errs.Error()).WriteTo(cxt)
		return
	}

	var users user.ImUsers

	result := model.DB.Table("im_users").Where("email=?", params.Email).First(&users)

	if result.RowsAffected == 0 {
		response.FailResponse(http.StatusInternalServerError, "邮箱未注册").ToJson(cxt)
		return
	}

	fmt.Println(users.Password)
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
		Avatar:     users.Avatar,
		Email:      users.Email,
		ExpireTime: expireAtTime,
		Token:      token,
		Ttl:        ttl,
	}).WriteTo(cxt)

	return

}

// 注册
func (*AuthController) Registered(cxt *gin.Context) {

	params := requests.RegisteredForm{
		Email:          cxt.PostForm("email"),
		Name:           cxt.PostForm("name"),
		EmailType:      helpers.StringToInt(cxt.DefaultPostForm("email_type", "1")),
		Password:       cxt.PostForm("password"),
		PasswordRepeat: cxt.PostForm("password_repeat"),
		Code:           cxt.PostForm("code"),
	}

	err := validator.New().Struct(params)

	if err != nil {
		response.FailResponse(enum.PARAMS_ERROR, err.Error()).WriteTo(cxt)
		return
	}

	ok, filed := user.IsUserExits(params.Email, params.Name)

	if ok {
		response.FailResponse(enum.PARAMS_ERROR, fmt.Sprintf("%s已经存在了", filed)).WriteTo(cxt)
		return
	}

	//var emailService services.EmailService
	//
	//if !emailService.CheckCode(params.Email, params.Code, params.EmailType) {
	//	response.FailResponse(enum.PARAMS_ERROR, "邮件验证码不正确").WriteTo(cxt)
	//	return
	//}

	createdAt := date.NewDate()

	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}

	model.DB.Table("im_users").Create(&user.ImUsers{
		Email:         params.Email,
		Password:      hash.BcryptHash(params.Password),
		Name:          params.Name,
		CreatedAt:     createdAt,
		UpdatedAt:     createdAt,
		Avatar:        fmt.Sprintf("https://api.multiavatar.com/Binx %s.png", params.Name),
		LastLoginTime: createdAt,
		Uid:           helpers.GetUuid(),
	})

	response.SuccessResponse().ToJson(cxt)
	return
}

// 发送邮件
func (*AuthController) SendEmailCode(cxt *gin.Context) {

	params := requests.SendEmailRequest{
		Email:     cxt.PostForm("email"),
		EmailType: helpers.StringToInt(cxt.PostForm("email_type")),
	}

	err := validator.New().Struct(params)

	if err != nil {
		response.FailResponse(enum.PARAMS_ERROR, err.Error()).WriteTo(cxt)
		return
	}

	ok := requests.IsTableFliedExits("email", params.Email, "im_users")

	switch params.EmailType {

	case services.REGISTERED_CODE:
		if ok {
			response.FailResponse(enum.PARAMS_ERROR, "邮箱已经被注册了").WriteTo(cxt)
			return
		}

	case services.RESET_PS_CODE:
		if !ok {
			response.FailResponse(enum.PARAMS_ERROR, "邮箱未注册了").WriteTo(cxt)
			return
		}

	}

	var emailService services.EmailService

	code := helpers.CreateEmailCode()

	html := fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Im-Services邮件验证码</title>
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
    <h3>您好 ~ im-services应用账号!</h3>
    <p>下面是您的验证码:</p>
        <p class="code">%s</p>
        <p>请注意查收!谢谢</p>
</div>
<h3>如果可以请给项目点个star～<a target="_blank" href="https://github.com/IM-Tools/Im-Services">项目地址</a> </h3>
</body>
</html>`, code)

	subject := "欢迎使用～👏Im Services,这是一封邮箱验证码的邮件!🎉🎉🎉"

	err = emailService.SendEmail(code, params.EmailType, params.Email, subject, html)
	if err != nil {
		logger.Logger.Error("发送失败邮箱:" + params.Email + "错误日志:" + err.Error())
		response.FailResponse(enum.API_ERROR, "邮件发送失败,请检查是否是可用邮箱").ToJson(cxt)
		return
	}

	response.SuccessResponse().ToJson(cxt)
	return

}
