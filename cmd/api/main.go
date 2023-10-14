package main

import (
	"github.com/HavvokLab/true-solar-monitoring/config"
	"github.com/HavvokLab/true-solar-monitoring/inverter/kstar"
	"github.com/HavvokLab/true-solar-monitoring/util"
)

func init() {
	config.InitConfig()
}

func init() {
}

func main() {
	username := "u2.kst"
	password := "Truec[8mugiup18"
	client, nil := kstar.NewKStarClient(&kstar.KStarCredential{Username: username, Password: password})
	res, err := client.GetRealtimeAlarmListOfDevice("I110261077A7021010003321")
	if err != nil {
		panic(err)
	}

	util.PrintJSON(map[string]interface{}{"x": res})
}
