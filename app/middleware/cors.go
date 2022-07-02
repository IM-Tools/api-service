/**
  @author:panliang
  @data:2022/6/13
  @note
**/
package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Cors 跨域
func Cors() gin.HandlerFunc {
	return func(cxt *gin.Context) {
		method := cxt.Request.Method

		cxt.Header("Access-Control-Allow-Origin", "*")
		cxt.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token")
		cxt.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		cxt.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		cxt.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			cxt.AbortWithStatus(http.StatusNoContent)
		}
	}
}
