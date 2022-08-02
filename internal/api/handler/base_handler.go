package handler

import (
	"github.com/gin-gonic/gin"
)

type Person struct {
	ID string `uri:"id" binding:"required"`
}

func GetPersonId(cxt *gin.Context) (error, Person) {
	var person Person
	if err := cxt.ShouldBindUri(&person); err != nil {
		return err, person
	}
	return nil, person
}
