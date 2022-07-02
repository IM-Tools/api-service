/**
  @author:panliang
  @data:2022/5/21
  @note
**/
package helpers

import (
	"fmt"
	"math/rand"
	"strconv"
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

func FirstElement(args []string) string {
	if len(args) > 0 {
		return args[0]
	}
	return ""
}
