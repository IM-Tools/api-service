/**
  @author:panliang
  @data:2022/6/3
  @note
**/
package requests

type LoginForm struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,string=6,20" json:"password"`
	Code     string `validate:"required,integer"`
}
