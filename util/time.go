package util

import "time"

func SetTimezone() {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	time.Local = loc
}

func GetLastDayOfMonth(year int, month time.Month) int {
	date := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	date = date.AddDate(0, 1, 0)
	date = date.AddDate(0, 0, -1)
	return date.Day()
}
