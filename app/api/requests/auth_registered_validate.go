/**
  @author:panliang
  @data:2022/6/3
  @note
**/
package requests

type RegisteredForm struct {
	Email          string `validate:"required,email" json:"email"`
	Name           string `validate:"required,email" json:"name"`
	Password       string `validate:"required,string=6,20" json:"password"`
	PasswordRepeat string `validate:"required,string=6,20" json:"password_repeat"`
	Code           string `validate:"required,integer"`
}
