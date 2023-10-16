package main

import (
	"github.com/HavvokLab/true-solar-monitoring/config"
	"github.com/HavvokLab/true-solar-monitoring/handler"
)

func init() {
	config.InitConfig()
}

func init() {
}

func main() {
	hdl := handler.NewKStarCollectorHandler()
	hdl.Run()
}
