package handler

import (
	"time"

	"github.com/HavvokLab/true-solar-monitoring/constant"
	"github.com/HavvokLab/true-solar-monitoring/infra"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/HavvokLab/true-solar-monitoring/service"
	"github.com/HavvokLab/true-solar-monitoring/util"
)

type PlantAggregateHandler struct {
	logger logger.Logger
}

func NewPlantAggregateHandler() *PlantAggregateHandler {
	return &PlantAggregateHandler{}
}

func (h *PlantAggregateHandler) Run() {
	now := time.Now()
	h.logger = logger.NewLogger(&logger.LoggerOption{
		LogLevel:   logger.LOG_LEVEL_DEBUG,
		LogName:    util.GetLogName(constant.PLANT_AGGR_LOG_NAME),
		LogSize:    1024,
		SkipCaller: 1,
		LogBackup:  3,
	})
	defer h.logger.Close()

	db, err := infra.NewGormDB()
	if err != nil {
		h.logger.Errorf("failed to init gorm db: %v", err)
		return
	}

	elastic, err := infra.NewElasticsearch()
	if err != nil {
		h.logger.Errorf("failed to init elasticsearch: %v", err)
		return
	}

	plantRepo := repo.NewPlantRepo(db)
	solarRepo := repo.NewSolarRepo(elastic)
	serv := service.NewPlantAggregateService(plantRepo, solarRepo, h.logger)
	if err := serv.UpdatePlantByDateToSQLite(&now); err != nil {
		h.logger.Errorf("failed to update plant by date: %v", err)
		return
	}

	if err := serv.UpdatePlantByDateToElastic(&now); err != nil {
		h.logger.Errorf("failed to update plant by date: %v", err)
		return
	}

	h.logger.Infof("success update plant by date: %v", now.Format("2006.01.02"))
}
