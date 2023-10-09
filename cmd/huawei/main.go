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
	conf := config.GetConfig().Huawei
	collector := handler.NewHuaweiCollectorHandler()
	alarm := handler.NewHuaweiAlarmHandler()

	cron := gocron.NewScheduler(time.Local)
	cron.Cron(conf.CollectorCrontab).Do(collector.Run)
	cron.Cron(conf.NightCollectorCrontab).Do(collector.Run)
	cron.Cron(conf.AlarmCrontab).Do(alarm.Run)
	cron.StartBlocking()
}
