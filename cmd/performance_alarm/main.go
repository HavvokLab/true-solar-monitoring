package main

import (
	"time"

	"github.com/HavvokLab/true-solar-monitoring/config"
	"github.com/HavvokLab/true-solar-monitoring/handler"
	"github.com/HavvokLab/true-solar-monitoring/util"
	"github.com/go-co-op/gocron"
)

func init() {
	config.InitConfig()
}

func init() {
	util.SetTimezone()
}

func main() {
	conf := config.GetConfig().DailyPerformanceAlarm
	dailyPerformanceAlarm := handler.NewDailyPerformanceAlarmHandler()

	cron := gocron.NewScheduler(time.Local)
	cron.Cron(conf.Crontab).Do(dailyPerformanceAlarm.Run)
	cron.StartBlocking()
}
