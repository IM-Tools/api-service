package tests

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"im-services/internal/api/services"
	"im-services/internal/config"
	"testing"
)

func init() {
	config.InitConfig("../config.yaml")
}

//http://localhost:3000/login?login_type=gitee&code=ae6b2386b3adc212f2c552d791dfb27e524af8ecfdddf9166ff6bd255b360ed7
func TestGitee(t *testing.T) {
	var err error
	var code = "b5b86f6461655d1f8b94ffc451c16fd3012d0e26fed8e0191d13428e2fd629ac"
	oauth := new(services.GiteeOAuthService)
	var tokenAuthUrl = oauth.GetTokenAuthUrl(code)
	fmt.Println(tokenAuthUrl)
	var token *services.Token
	if token, err = oauth.GetToken(tokenAuthUrl); err != nil {
		assert.Equal(t, err, "有错误发生，err 不为空")
		return
	}
	var userInfo map[string]interface{}
	userInfo, err = oauth.GetUserInfo(token.AccessToken)

	fmt.Println(userInfo)
	if err != nil {
		assert.Equal(t, err, "有错误发生，err 不为空")
		return
	}
}
