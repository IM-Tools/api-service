/**
  @author:panliang
  @data:2022/5/17
  @note
**/
package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var hmacSampleSecret []byte

type CustomClaims struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func GenerateTokenStr(claims *CustomClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		panic(err)
	}
	return tokenString
}
