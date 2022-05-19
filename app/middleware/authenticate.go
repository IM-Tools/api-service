/**
  @author:panliang
  @data:2022/5/16
  @note
**/
package middleware

import "github.com/gin-gonic/gin"

// 后台校验中间件
func AdminAuth() gin.HandlerFunc {

	return func(cxt *gin.Context) {
		cxt.Next()
	}

}