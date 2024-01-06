package handler

import (
	"time"

	"github.com/HavvokLab/true-solar-monitoring/constant"
	"github.com/HavvokLab/true-solar-monitoring/infra"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/HavvokLab/true-solar-monitoring/service"
	"github.com/HavvokLab/true-solar-monitoring/util"
	"github.com/gammazero/workerpool"
)

type MonthlyProductionHandler struct {
	logger logger.Logger
}

func NewMonthlyProductionHandler() *MonthlyProductionHandler {
	return &MonthlyProductionHandler{}
}

func (h *MonthlyProductionHandler) RunAll() {
	h.logger = logger.NewLogger(
		&logger.LoggerOption{
			LogName:     util.GetLogName(constant.MONTHLY_PRODUCTION_LOG_NAME),
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
	now := time.Now()
	start := time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)
	end := start.AddDate(0, 1, 0).AddDate(0, 0, -1) // (1 january + 1 month) - 1 day = last day of current month

	for {
		startDate := start
		endDate := end
		pool.Submit(h.run(&startDate, &endDate))
		if now.Month() == start.Month() && now.Year() == start.Year() {
			break
		}

		start = start.AddDate(0, 1, 0)
		end = start.AddDate(0, 1, 0).AddDate(0, 0, -1)
	}

	pool.StopWait()
}

func (h *MonthlyProductionHandler) Run() {
	h.logger = logger.NewLogger(
		&logger.LoggerOption{
			LogName:     util.GetLogName(constant.MONTHLY_PRODUCTION_LOG_NAME),
			LogSize:     1024,
			LogAge:      90,
			LogBackup:   1,
			LogCompress: false,
			LogLevel:    logger.LOG_LEVEL_DEBUG,
			SkipCaller:  1,
		},
	)
	defer h.logger.Close()

	now := time.Now()
	if now.Day() == 1 {
		end := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
		start := end.AddDate(0, -1, 0)
		h.run(&start, &end)()
	}
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

func (h *MonthlyProductionHandler) DailyRun() {
	h.logger = logger.NewLogger(
		&logger.LoggerOption{
			LogName:     util.GetLogName(constant.MONTHLY_PRODUCTION_LOG_NAME),
			LogSize:     1024,
			LogAge:      90,
			LogBackup:   1,
			LogCompress: false,
			LogLevel:    logger.LOG_LEVEL_DEBUG,
			SkipCaller:  1,
		},
	)
	defer h.logger.Close()

	h.dailyRun()()
}

func (h *MonthlyProductionHandler) dailyRun() func() {
	return func() {
		elastic, err := infra.NewElasticsearch()
		now := time.Now()
		if err != nil {
			h.logger.Errorf("[%v]Failed to connect to elasticsearch", now.Format(constant.YEAR_MONTH))
			return
		}

		masterSiteRepo, err := repo.NewMasterSiteRepo()
		if err != nil {
			h.logger.Error(err)
			return
		}

		solarRepo := repo.NewSolarRepo(elastic)
		serv := service.NewMonthlyProductionService(solarRepo, masterSiteRepo, h.logger)
		if err := serv.DailyRun(); err != nil {
			h.logger.Error(err)
		}
	}
}
