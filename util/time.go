package util

import "time"

func SetTimezone() {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	time.Local = loc
}
