package helpers

import "time"

func GetTimeNow() string {
	return time.Now().Local().Format("02-01-2006 15:04:05")

}
