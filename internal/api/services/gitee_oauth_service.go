package services

import (
	"encoding/json"
	"fmt"
	"im-services/internal/config"
	"net/http"
)

type GiteeOAuthService struct {
}

//获取地址
func (*GiteeOAuthService) GetTokenAuthUrl(code string) string {
	return fmt.Sprintf(
		"https://gitee.com/oauth/token?client_id=%s&redirect_uri=%s&code=%s&grant_type=authorization_code&client_secret=%s",
		config.Conf.Gitee.ClientId, config.Conf.Gitee.RedirectUrl, code, config.Conf.Gitee.ClientSecret,
	)
}

func (*GiteeOAuthService) GetToken(url string) (*Token, error) {

	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodPost, url, nil); err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")

	// 发送请求并获得响应
	var httpClient = http.Client{}
	var res *http.Response
	if res, err = httpClient.Do(req); err != nil {
		return nil, err
	}
	var token Token
	// 将响应体解析为 token，并返回

	if err = json.NewDecoder(res.Body).Decode(&token); err != nil {
		return nil, err
	}

	return &token, nil
}

func (*GiteeOAuthService) GetUserInfo(token string) (map[string]interface{}, error) {

	// 形成请求
	var userInfoUrl = "https://gitee.com/api/v5/user?access_token=" + token
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, userInfoUrl, nil); err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	// 发送请求并获取响应
	var client = http.Client{}
	var res *http.Response
	if res, err = client.Do(req); err != nil {
		return nil, err
	}

	// 将响应的数据写入 userInfo 中，并返回
	var userInfo = make(map[string]interface{})

	if err = json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
		return nil, err
	}
	return userInfo, nil
}
