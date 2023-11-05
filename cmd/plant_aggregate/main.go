package main

import (
	"time"

	"github.com/HavvokLab/true-solar-monitoring/config"
	"github.com/HavvokLab/true-solar-monitoring/handler"
	"github.com/HavvokLab/true-solar-monitoring/util"
	"github.com/go-co-op/gocron"
)

func init() {
	util.SetTimezone()
	config.InitConfig()
}

func main() {
	conf := config.GetConfig().PlantAggregate
	hdl := handler.NewPlantAggregateHandler()

	cron := gocron.NewScheduler(time.Local)
	cron.Cron(conf.Crontab).Do(hdl.Run)
	cron.StartBlocking()
}
