/**
  @author:panliang
  @data:2022/7/14
  @note
**/
package grpcAuth

import (
	"context"
	"errors"
	"im-services/pkg/jwt"
	"strings"
)

type AuthServer struct {
}

// token校验
func (s *AuthServer) CheckAuth(ctx context.Context, req *CheckAuthRequest) (*CheckAuthResponse, error) {
	err, token := ValidatedToken(req.Token)
	if err != nil {
		return &CheckAuthResponse{}, err
	}

	claims, err := jwt.NewJWT().ParseToken(token)
	if err != nil {
		return &CheckAuthResponse{}, err
	}

	return &CheckAuthResponse{
		Id:    claims.ID,
		Email: claims.Email,
		Uid:   claims.UID,
		Name:  claims.Name,
	}, nil
}

func ValidatedToken(token string) (error, string) {

	if len(token) == 0 {
		return errors.New("Token 不能为空"), ""
	}

	t := strings.Split(token, "Bearer ")
	if len(t) > 1 {
		return nil, t[1]
	}
	return nil, token
}
