package handler

import (
	"time"

	"github.com/HavvokLab/true-solar-monitoring/constant"
	"github.com/HavvokLab/true-solar-monitoring/infra"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/HavvokLab/true-solar-monitoring/service"
	"github.com/gammazero/workerpool"
)

type MonthlyProductionHandler struct {
	logger logger.Logger
}

func NewMonthlyProductionHandler() *MonthlyProductionHandler {
	return &MonthlyProductionHandler{}
}

func (h *MonthlyProductionHandler) Run() {
	h.logger = logger.NewLogger(
		&logger.LoggerOption{
			LogName:     constant.MONTHLY_PRODUCTION_LOG_NAME,
			LogSize:     1024,
			LogAge:      90,
			LogBackup:   1,
			LogCompress: false,
			LogLevel:    logger.LOG_LEVEL_DEBUG,
			SkipCaller:  1,
		},
	)
	defer h.logger.Close()

	pool := workerpool.New(5)
	currentMonth := time.January
	endDate := time.Now()

	for {
		start := time.Date(2023, currentMonth, 1, 0, 0, 0, 0, time.Local)
		end := time.Date(2023, currentMonth+1, 1, 0, 0, 0, 0, time.Local)

		pool.Submit(h.run(&start, &end))
		if endDate.Month() == currentMonth {
			break
		} else {
			currentMonth += 1
		}
	}

	pool.StopWait()
}

func (h *MonthlyProductionHandler) run(start, end *time.Time) func() {
	return func() {
		elastic, err := infra.NewElasticsearch()
		if err != nil {
			h.logger.Errorf("[%v]Failed to connect to elasticsearch", start.Format(constant.YEAR_MONTH))
			return
		}

		masterSiteRepo, err := repo.NewMasterSiteRepo()
		if err != nil {
			h.logger.Error(err)
			return
		}

		solarRepo := repo.NewSolarRepo(elastic)
		serv := service.NewMonthlyProductionService(solarRepo, masterSiteRepo, h.logger)
		if err := serv.Run(start, end); err != nil {
			h.logger.Error(err)
		}
	}
}
