/**
  @author:panliang
  @data:2022/6/3
  @note
**/
package middleware

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var (
	Uni      *ut.UniversalTranslator
	Validate *validator.Validate
)

func ValidateTranslations() gin.HandlerFunc {

	return func(cxt *gin.Context) {
		cxt.Next()
	}

}
