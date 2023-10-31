package main

import (
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
	repo := repo.NewPlantRepo(infra.GormDB)
	serv := service.NewPlantService(repo, logger.NewLoggerMock())
	if err := serv.ExportToCsv(); err != nil {
		panic(err)
	}

	println("done")
}
