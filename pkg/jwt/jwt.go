/**
  @author:panliang
  @data:2022/5/17
  @note
**/
package jwt

import (
	"errors"
	goJwt "github.com/golang-jwt/jwt/v4"
	"im-services/config"
	"im-services/pkg/logger"
	"time"
)

type JWT struct {
	SigningKey []byte
	MaxRefresh time.Duration
}

type CustomClaims struct {
	ID         int64  `json:"id"`
	UID        string `json:"uid"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	ExpireTime int64  `json:"expire_time"`

	// StandardClaims 结构体实现了 Claims 接口继承了  Valid() 方法
	// JWT 规定了7个官方字段，提供使用:
	// - iss (issuer)：发布者
	// - sub (subject)：主题
	// - iat (Issued At)：生成签名的时间
	// - exp (expiration time)：签名过期时间
	// - aud (audience)：观众，相当于接受者
	// - nbf (Not Before)：生效时间
	// - jti (JWT ID)：编号

	goJwt.StandardClaims
}

var (
	TokenInvalid error = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		SigningKey: []byte(config.Conf.JWT.Secret),
		MaxRefresh: time.Duration(config.Conf.JWT.Ttl) * time.Minute,
	}
}

// 创建token
func (j *JWT) createToken(claims CustomClaims) (string, error) {
	token := goJwt.NewWithClaims(goJwt.SigningMethodHS256, claims)
	res, err := token.SignedString(j.SigningKey)
	return res, err

}

// 解析token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := goJwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *goJwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, TokenInvalid
}

// 刷新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	goJwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	token, err := goJwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *goJwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		goJwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.createToken(*claims)
	}

	return "", TokenInvalid
}

func (jwt *JWT) IssueToken(ID int64, UID string, Name string, Email string, expireAtTime int64) string {

	claims := CustomClaims{
		ID,
		UID,
		Name,
		Email,
		expireAtTime,
		goJwt.StandardClaims{
			NotBefore: time.Now().Unix(),       // 签名生效时间
			IssuedAt:  time.Now().Unix(),       // 首次签名时间（后续刷新 Token 不会更新）
			ExpiresAt: expireAtTime,            // 签名过期时间
			Issuer:    config.Conf.Server.Name, // 签名颁发者
		},
	}

	// 2. 根据 claims 生成token对象
	token, err := jwt.createToken(claims)
	if err != nil {
		logger.Logger.DPanic(err.Error())
		return ""
	}
	return token
}
