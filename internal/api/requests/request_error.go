package requests

import (
	"github.com/go-playground/validator/v10"
)

func GetError(err error) string {

	errs := err.(validator.ValidationErrors)

	var message string
	for _, e := range errs {
		message = e.Error()
		break
	}
	return message
}
