package interfaces

import "im-services/internal/api/services"

type OAuth interface {
	GetTokenAuthUrl(code string) string
	GetToken(url string) (*services.Token, error)
	GetUserInfo(token string) (map[string]interface{}, error)
}
