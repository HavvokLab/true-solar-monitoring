package util

import (
	"fmt"
	"time"
)

func GetLogName(name string) string {
	now := time.Now()
	return fmt.Sprintf("logs/%v/%v.log", name, now.Format("20060102"))
}
