package handler

import (
	"math"
	"time"

	"github.com/HavvokLab/true-solar-monitoring/constant"
	"github.com/HavvokLab/true-solar-monitoring/infra"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/HavvokLab/true-solar-monitoring/service"
	"github.com/gammazero/workerpool"
)

type DailyProductionHandler struct {
	logger logger.Logger
}

func NewDailyProductionHandler() *DailyProductionHandler {
	return &DailyProductionHandler{}
}

func (h *DailyProductionHandler) RunAll() {
	h.logger = logger.NewLogger(
		&logger.LoggerOption{
			LogName:     constant.DAILY_PRODUCTION_LOG_NAME,
			LogSize:     1024,
			LogAge:      90,
			LogBackup:   1,
			LogCompress: false,
			LogLevel:    logger.LOG_LEVEL_DEBUG,
			SkipCaller:  1,
		},
	)
	defer h.logger.Close()

	pool := workerpool.New(10)
	startExecute := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.Local)
	endExecute := time.Now()
	duration := int(math.Ceil(endExecute.Sub(startExecute).Hours() / 24))

	initialDate := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.Local)
	for i := 0; i < duration; i++ {
		start := initialDate
		end := start.Add(24 * time.Hour)
		pool.Submit(h.run(&start, &end))
		initialDate = end
	}
	pool.StopWait()
}

func (h *DailyProductionHandler) Run() {
	h.logger = logger.NewLogger(
		&logger.LoggerOption{
			LogName:     constant.DAILY_PRODUCTION_LOG_NAME,
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
	end := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	start := end.AddDate(0, 0, -1)
	h.run(&start, &end)
}

func (h *DailyProductionHandler) run(start, end *time.Time) func() {
	return func() {
		elastic, err := infra.NewElasticsearch()
		if err != nil {
			h.logger.Errorf("[%v]Failed to connect to elasticsearch", start.Format(constant.YEAR_MONTH_DAY))
			return
		}

		masterSiteRepo, err := repo.NewMasterSiteRepo()
		if err != nil {
			h.logger.Error(err)
			return
		}

		solarRepo := repo.NewSolarRepo(elastic)
		serv := service.NewDailyProductionService(solarRepo, masterSiteRepo, h.logger)
		if err := serv.Run(start, end); err != nil {
			h.logger.Error(err)
		}
	}
}
