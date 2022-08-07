package tests

import (
	"fmt"
	"gotest.tools/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"testing"
)

const (
	BASE_API = "http://localhost:8000"
)

func Login() (response *http.Response, errs error, string2 string) {
	var (
		resp *http.Response
		err  error
	)

	resp, err = http.PostForm(BASE_API+"/api/auth/login", url.Values{
		"email":    {"2540463097@qq.com"},
		"password": {"123456"},
	})

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
	return resp, err, string(body)
}

func TestLogin(t *testing.T) {

	resp, err, body := Login()

	fmt.Println(body)
	if err != nil {
		assert.Error(t, err, "有错误发生，err 不为空")
		return
	}
	assert.Equal(t, 200, resp.StatusCode)
}
