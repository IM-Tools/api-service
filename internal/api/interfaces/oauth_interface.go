package interfaces

type OAuth interface {
	GetTokenAuthUrl(code string) string
	GetToken(url string) (*Token, error)
	GetUserInfo(token *Token) (map[string]interface{}, error)
}
type Token struct {
	AccessToken string `json:"access_token"`
}
