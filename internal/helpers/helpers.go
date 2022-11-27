package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/axgle/mahonia"
	uuid "github.com/satori/go.uuid"
	"im-services/pkg/logger"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetNowFormatTodayTime() string {
	now := time.Now()
	dateStr := fmt.Sprintf("%02d-%02d-%02d", now.Year(), int(now.Month()),
		now.Day())

	return dateStr
}

func CreateEmailCode() string {
	return fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
}

func GetDayTime(days int) int64 {
	nowTimeStr := time.Now().Format("2006-01-02 15:04:05")
	//使用Parse 默认获取为UTC时区 需要获取本地时区 所以使用ParseInLocation
	timeS, _ := time.ParseInLocation("2006-01-02", nowTimeStr, time.Local)
	timeStamp := timeS.AddDate(0, 0, days).Unix()
	return timeStamp
}

func Int64ToString(int64_ int64) string {
	return strconv.Itoa(int(int64_))
}

func Float64ToString(float64_ float64) string {
	return strconv.Itoa(int(float64_))
}

func StringToInt(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}
func StringToInt64(str string) int64 {
	num, _ := strconv.Atoi(str)
	return int64(num)
}

func FirstElement(args []string) string {
	if len(args) > 0 {
		return args[0]
	}
	return ""
}

func Explode(delimiter, text string) []string {
	if len(delimiter) > len(text) {
		return strings.Split(delimiter, text)
	} else {
		return strings.Split(text, delimiter)
	}
}

func GetUuid() string {

	u1 := uuid.NewV4()

	return fmt.Sprintf("%s", u1)
}

func InterfaceToInt64(inter interface{}) int64 {

	return inter.(int64)
}

func InterfaceToInt64String(inter interface{}) string {
	int64Val := inter.(int64)
	return Int64ToString(int64Val)
}
func InterfaceToString(inter interface{}) string {

	return inter.(string)
}
func ErrorHandler(err error) {
	if err != nil {
		logger.Logger.Error(err.Error())
		return
	}
	return
}

type IpInfo struct {
	IP          string `json:"ip"`
	Pro         string `json:"pro"`
	ProCode     string `json:"proCode"`
	City        string `json:"city"`
	CityCode    string `json:"cityCode"`
	Region      string `json:"region"`
	RegionCode  string `json:"regionCode"`
	Addr        string `json:"addr"`
	RegionNames string `json:"regionNames"`
	Err         string `json:"err"`
}

func GetIpInfo(ip string) (error, IpInfo) {
	var info IpInfo
	url := "http://whois.pconline.com.cn/ipJson.jsp?ip=" + ip + "&json=true"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	req.Header.Set("accept", "application/json")
	if err != nil {
		return err, info
	}
	res, err := client.Do(req)
	if err != nil {
		return err, info
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	bodyStr := mahonia.NewDecoder("gbk").ConvertString(string(body))
	err = json.Unmarshal([]byte(bodyStr), &info)
	if err != nil {
		return err, info
	}
	return nil, info

}
