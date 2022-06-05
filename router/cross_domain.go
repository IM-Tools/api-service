/**
  @author:panliang
  @data:2022/5/16
  @note
**/
package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 设置服务端支持跨域
func startCors(router *gin.Engine) {

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	config.AllowHeaders = []string{
		"tus-resumable",
		"upload-length",
		"upload-metadata",
		"cache-control",
		"x-requested-with",
		"*",
	}
	router.Use(cors.New(config))

}
