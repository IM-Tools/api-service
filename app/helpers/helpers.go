/**
  @author:panliang
  @data:2022/5/21
  @note
**/
package helpers

import (
	"fmt"
	"time"
)

func GetNowFormatTodayTime() string {
	now := time.Now()
	dateStr := fmt.Sprintf("%02d-%02d-%02d", now.Year(), int(now.Month()),
		now.Day())

	return dateStr
}
