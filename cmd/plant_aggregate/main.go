package main

import (
	"time"

	"github.com/HavvokLab/true-solar-monitoring/config"
	"github.com/HavvokLab/true-solar-monitoring/infra"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/HavvokLab/true-solar-monitoring/service"
)

func init() {
	config.InitConfig()
}

func init() {
	infra.InitGormDB()
}

func main() {
	elastic, err := infra.NewElasticsearch()
	if err != nil {
		panic(err)
	}

	solarRepo := repo.NewSolarRepo(elastic)
	plantRepo := repo.NewPlantRepo(infra.GormDB)
	serv := service.NewPlantAggregateService(plantRepo, solarRepo, logger.NewLoggerMock())

	for i := 1; i <= 10; i++ {
		date := time.Date(2023, time.Month(i), 1, 0, 0, 0, 0, time.Local)
		if err := serv.UpdatePlantByDate(&date); err != nil {
			panic(err)
		}
	}
}
