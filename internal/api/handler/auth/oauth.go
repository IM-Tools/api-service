package auth

import (
	"github.com/gin-gonic/gin"
	"im-services/internal/api/interfaces"
	"im-services/internal/api/services"
	"im-services/internal/config"
	"im-services/internal/helpers"
	"im-services/pkg/jwt"
	"im-services/pkg/response"
	"net/http"
	"time"
)

type OAuthHandler struct {
}

type Token struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

// @BasePath /api

// PingExample godoc
// @Summary friends github登录
// @Schemes
// @Description github登录
// @Tags 登录相关
// @Produce json
// @Param code path int true "github授权码"
// @Success 200 {object} response.JsonResponse{data=loginResponse} "ok"
// @Router /auth/githubLogin [get]
func (*OAuthHandler) GithubOAuth(cxt *gin.Context) {
	var err error
	var code = cxt.Query("code")
	var loginType = cxt.Query("login_type")

	var oauth interfaces.OAuth

	if loginType == "gitee" {
		oauth = new(services.GiteeOAuthService)
	} else {
		oauth = new(services.GithubOAuthService)
	}

	var tokenAuthUrl = oauth.GetTokenAuthUrl(code)
	var token *services.Token

	if token, err = oauth.GetToken(tokenAuthUrl); err != nil {
		response.FailResponse(http.StatusInternalServerError, err.Error()).WriteTo(cxt)
		return
	}
	var userInfo map[string]interface{}
	userInfo, err = oauth.GetUserInfo(token.AccessToken)
	if err != nil {
		response.FailResponse(http.StatusInternalServerError, err.Error()).WriteTo(cxt)
		return
	}

	err, users, isNew := auth.CreateOauthUser(userInfo, loginType)
	ttl := config.Conf.JWT.Ttl
	expireAtTime := time.Now().Unix() + ttl
	tokens := jwt.NewJWT().IssueToken(
		users.ID,
		users.Uid,
		users.Name,
		users.Email,
		expireAtTime,
	)
	// 异地登录事件
	eventHandle.LogoutEvent(helpers.Int64ToString(users.ID), cxt.Request.Header.Get("X-Forward-For"))

	if isNew {
		// 注册事件
		eventHandle.RegisterEvent(users.ID, users.Name)
	}

	response.SuccessResponse(&loginResponse{
		ID:         users.ID,
		UID:        users.Uid,
		Name:       users.Name,
		Avatar:     users.Avatar,
		Email:      users.Email,
		ExpireTime: expireAtTime,
		Token:      tokens,
		Ttl:        ttl,
	}).WriteTo(cxt)

}
