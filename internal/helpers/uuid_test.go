package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/magiconair/properties/assert"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestIp(t *testing.T) {

	var ip = "113.92.72.58"
	var info IpInfo
	url := "http://whois.pconline.com.cn/ipJson.jsp?ip=" + ip + "&json=true"
	resp, err := http.Get(url)
	if err != nil {
		log.Println("err")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("err")
	}
	fmt.Println(string(body))
	err = json.Unmarshal(body, &info)
	if err != nil {
		assert.Equal(t, err, "有错误发生，err 不为空")
		return
	}
	assert.Equal(t, 200, resp.StatusCode)

}
