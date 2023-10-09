package main

import (
	"fmt"

	"github.com/HavvokLab/true-solar-monitoring/config"
	"github.com/HavvokLab/true-solar-monitoring/infra"
	"github.com/HavvokLab/true-solar-monitoring/util"
)

func init() {
	config.InitConfig()
}

func init() {
	util.SetTimezone()
}

func main() {
	// hdl := handler.NewSolarmanAlarmHandler()
	// hdl.Mock()
	// hdl := handler.NewSolarmanCollectorHandler()
	// hdl.Mock()
	_, err := infra.NewRedis()
	if err != nil {
		panic(err)
	}

	fmt.Println("done")
}
