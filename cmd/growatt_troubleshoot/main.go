package main

import (
	"github.com/HavvokLab/true-solar-monitoring/handler"
	"github.com/HavvokLab/true-solar-monitoring/util"
)

func init() {
	util.SetTimezone()
}

func main() {
	hdl := handler.NewGrowattTroubleShootHandler()
	hdl.Run()
}
