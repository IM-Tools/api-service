package auth

import (
	"fmt"
	"im-services/internal/api/requests"
	"im-services/internal/api/services"
	"im-services/internal/config"
	"im-services/internal/dao/auth_dao"
	"im-services/internal/enum"
	"im-services/internal/helpers"
	"im-services/internal/models/user"
	"im-services/pkg/hash"
	"im-services/pkg/jwt"
	"im-services/pkg/logger"
	"im-services/pkg/model"
	"im-services/pkg/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
}

type AuthHandlerInterface interface {

	// Login 登录

	Login(cxt *gin.Context)

	// Registered 注册

	Registered(cxt *gin.Context)

	// SendEmailCode 发送邮件

	SendEmailCode(cxt *gin.Context)
}

type loginResponse struct {
	ID         int64  `json:"id"`          //用户id
	UID        string `json:"uid"`         // uid
	Name       string `json:"name"`        //名称
	Avatar     string `json:"avatar"`      //头像
	Email      string `json:"email"`       //邮箱账号
	Token      string `json:"token"`       // token
	ExpireTime int64  `json:"expire_time"` // token过期时间
	Ttl        int64  `json:"ttl"`         // token有效期
}

var (
	auth auth_dao.AuthDao
)

// Login 登录
// @BasePath /api

// PingExample godoc
// @Summary Login 登录
// @Schemes
// @Description 登录接口
// @Tags 登录相关
// @Accept multipart/form-data
// @Produce json
// @Param email formData string true "账号"
// @Param password formData string true "密码"
// @Success 200 {object} response.JsonResponse{data=loginResponse} "ok"
// @Router /auth/login [post]
func (*AuthHandler) Login(cxt *gin.Context) {

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

// Registered 注册
// @BasePath /api

// PingExample godoc
// @Summary Registered 注册
// @Schemes
// @Description 注册接口
// @Tags 登录相关
// @Accept multipart/form-data
// @Produce json
// @Param email formData string true "账号"
// @Param name formData string true "名称"
// @Param email_type formData int false "邮件类型 1.注册 2.找回密码"
// @Param password formData string true "密码"
// @Param password_repeat formData string true "确认密码"
// @Param code formData string true "验证码"
// @Success 200 {object} response.JsonResponse{} "ok"
// @Router /auth/registered [post]
func (*AuthHandler) Registered(cxt *gin.Context) {

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
		response.FailResponse(enum.ParamError, err.Error()).WriteTo(cxt)
		return
	}

	ok, filed := user.IsUserExits(params.Email, params.Name)

	if ok {
		response.FailResponse(enum.ParamError, fmt.Sprintf("%s已经存在了", filed)).WriteTo(cxt)
		return
	}

	// var emailService services.EmailService

	// if !emailService.CheckCode(params.Email, params.Code, params.EmailType) {
	// 	response.FailResponse(enum.ParamError, "邮件验证码不正确").WriteTo(cxt)
	// 	return
	// }

	id := auth.CreateUser(params.Email, params.Password, params.Name)

	// 投递消息
	services.InitChatBotMessage(1, id)

	response.SuccessResponse().ToJson(cxt)
	return
}

// Registered 发送邮件
// @BasePath /api

// PingExample godoc
// @Summary Registered 发送邮件
// @Schemes
// @Description 发送邮件接口
// @Tags 登录相关
// @Accept multipart/form-data
// @Produce json
// @Param email formData string true "邮箱"
// @Param email_type formData int true "邮件类型 1.注册 2.找回密码"
// @Success 200 {object} response.JsonResponse{} "ok"
// @Router /auth/sendEmailCode [post]
func (*AuthHandler) SendEmailCode(cxt *gin.Context) {

	params := requests.SendEmailRequest{
		Email:     cxt.PostForm("email"),
		EmailType: helpers.StringToInt(cxt.PostForm("email_type")),
	}

	err := validator.New().Struct(params)

	if err != nil {
		response.FailResponse(enum.ParamError, err.Error()).WriteTo(cxt)
		return
	}

	ok := requests.IsTableFliedExits("email", params.Email, "im_users")

	switch params.EmailType {

	case services.REGISTERED_CODE:
		if ok {
			response.FailResponse(enum.ParamError, "邮箱已经被注册了").WriteTo(cxt)
			return
		}

	case services.RESET_PS_CODE:
		if !ok {
			response.FailResponse(enum.ParamError, "邮箱未注册了").WriteTo(cxt)
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
		response.FailResponse(enum.ApiError, "邮件发送失败,请检查是否是可用邮箱").ToJson(cxt)
		return
	}

	response.SuccessResponse().ToJson(cxt)
	return

}
