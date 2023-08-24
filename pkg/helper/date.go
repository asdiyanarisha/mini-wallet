package helper

import "time"

var LoadLocDate, _ = time.LoadLocation("Asia/Jakarta")

func InitDate() string {
	return time.Now().In(LoadLocDate).Format("2006-01-02T15:04:05-07:00")
}
