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
	conf := config.GetConfig().Solarman

	solarmanCollector := handler.NewSolarmanCollectorHandler()
	solarmanAlarm := handler.NewSolarmanAlarmHandler()

	cron := gocron.NewScheduler(time.Local)
	cron.Cron(conf.CollectorCrontab).Do(solarmanCollector.Run)
	cron.Cron(conf.NightCollectorCrontab).Do(solarmanCollector.Run)
	cron.Cron(conf.AlarmCrontab).Do(solarmanAlarm.Run)
	cron.StartBlocking()
}
