/**
  @author:panliang
  @data:2022/6/3
  @note
**/
package hash

import (
	"fmt"
	"testing"
)

func TestCreatePw(t *testing.T) {
	password := BcryptHash("123456")
	fmt.Println(password)
}
