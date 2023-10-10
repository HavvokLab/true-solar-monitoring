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
	dailyPerformanceAlarmConfig := config.GetConfig().DailyPerformanceAlarm
	lowPerformanceAlarmConfig := config.GetConfig().LowPerformanceAlarm
	dailyPerformanceAlarm := handler.NewDailyPerformanceAlarmHandler()
	lowPerformanceAlarm := handler.NewLowPerformanceAlarmHandler()

	cron := gocron.NewScheduler(time.Local)
	cron.Cron(dailyPerformanceAlarmConfig.Crontab).Do(dailyPerformanceAlarm.Run)
	cron.Cron(lowPerformanceAlarmConfig.Crontab).Do(lowPerformanceAlarm.Run)
	cron.StartBlocking()
}
