package middleware

import (
	"github.com/gin-gonic/gin"
	"im-services/pkg/jwt"
	"im-services/pkg/response"
	"net/http"
	"strings"
)

func Auth() gin.HandlerFunc {

	return func(cxt *gin.Context) {

		token := cxt.DefaultQuery("token", cxt.GetHeader("authorization"))

		err, token := ValidatedToken(token)
		if err != nil {
			response.ErrorResponse(http.StatusUnauthorized, err.Error()).WriteTo(cxt)
			cxt.Abort()
			return
		}

		claims, err := jwt.NewJWT().ParseToken(token)
		if err != nil {
			response.ErrorResponse(http.StatusUnauthorized, err.Error()).WriteTo(cxt)
			cxt.Abort()
			return
		}

		cxt.Set("id", claims.ID)
		cxt.Set("uid", claims.UID)

		cxt.Next()
	}

}

func ValidatedToken(token string) (error, string) {

	var err error
	if len(token) == 0 {
		return err, "Token 不能为空"
	}

	t := strings.Split(token, "Bearer ")
	if len(t) > 1 {
		return nil, t[1]
	}
	return nil, token
}
