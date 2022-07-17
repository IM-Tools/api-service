/**
  @author:panliang
  @data:2022/6/3
  @note
**/
package requests

type LoginForm struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required" `
}
