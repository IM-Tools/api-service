package auth

import (
	"github.com/gin-gonic/gin"
	"im-services/internal/api/services"
	"im-services/internal/config"
	"im-services/pkg/jwt"
	"im-services/pkg/response"
	"net/http"
	"time"
)

type OAuthHandler struct {
}

type Token struct {
	AccessToken string `json:"access_token"`
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
	oauth := new(services.GithubOAuthService)
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
	err, users := auth.CreateGithubUser(userInfo)
	ttl := config.Conf.JWT.Ttl
	expireAtTime := time.Now().Unix() + ttl
	tokens := jwt.NewJWT().IssueToken(
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
		Token:      tokens,
		Ttl:        ttl,
	}).WriteTo(cxt)

}
