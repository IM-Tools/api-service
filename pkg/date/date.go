package date

import "time"

func NewDate() string {
	return time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
}

func TimeUnixNano() int64 {
	return time.Now().UnixNano()
}

func TimeUnix() int64 {
	return time.Now().Unix()
}
