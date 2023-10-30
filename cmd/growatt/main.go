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
	conf := config.GetConfig().Growatt
	collector := handler.NewGrowattCollectorHandler()

	cron := gocron.NewScheduler(time.Local)
	cron.Cron(conf.CollectorCrontab).Do(collector.Run)
	cron.Cron(conf.NightCollectorCrontab).Do(collector.Run)
	cron.StartBlocking()
}
