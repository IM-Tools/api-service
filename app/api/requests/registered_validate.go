/**
  @author:panliang
  @data:2022/6/3
  @note
**/
package requests

type RegisteredForm struct {
	Email          string `validate:"required,email" `
	Name           string `validate:"required"`
	EmailType      int    `validate:"required,gte=1,lte=2"`
	Password       string `json:"password" validate:"required,min=6,max=10"`
	PasswordRepeat string `validate:"required,eqcsfield=Password"`
	Code           string `validate:"required,len=4"`
}
