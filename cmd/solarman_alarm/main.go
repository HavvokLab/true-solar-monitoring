package main

import (
	"github.com/HavvokLab/true-solar-monitoring/config"
	"github.com/HavvokLab/true-solar-monitoring/handler"
	"github.com/HavvokLab/true-solar-monitoring/util"
)

func init() {
	config.InitConfig()
}

func init() {
	util.SetTimezone()
}

func main() {
	solarmanAlarm := handler.NewSolarmanAlarmHandler()
	solarmanAlarm.Run()
}
