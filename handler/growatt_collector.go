package handler

import (
	"time"

	"github.com/HavvokLab/true-solar-monitoring/constant"
	"github.com/HavvokLab/true-solar-monitoring/infra"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/HavvokLab/true-solar-monitoring/service"
	"github.com/HavvokLab/true-solar-monitoring/util"
	"github.com/gammazero/workerpool"
)

type GrowattCollectorHandler struct {
	logger logger.Logger
}

func NewGrowattCollectorHandler() *GrowattCollectorHandler {
	return &GrowattCollectorHandler{}
}

func (h *GrowattCollectorHandler) Run() {
	h.logger = logger.NewLogger(
		&logger.LoggerOption{
			LogName:     util.GetLogName(constant.GROWATT_COLLECTOR_LOG_NAME),
			LogSize:     1024,
			LogAge:      90,
			LogBackup:   1,
			LogCompress: false,
			LogLevel:    logger.LOG_LEVEL_DEBUG,
			SkipCaller:  1,
		},
	)
	defer h.logger.Close()

	db, err := infra.NewGormDB()
	if err != nil {
		h.logger.Error(err)
		return
	}

	credentialRepo := repo.NewGrowattCredentialRepo(db)
	credentials, err := credentialRepo.FindAll()
	if err != nil {
		h.logger.Error(err)
		return
	}

	pool := workerpool.New(len(credentials))
	for _, credential := range credentials {
		clone := credential
		pool.Submit(h.run(&clone))
	}
	pool.StopWait()
}

func (h *GrowattCollectorHandler) run(credential *model.GrowattCredential) func() {
	return func() {
		now := time.Now()
		elastic, err := infra.NewElasticsearch()
		if err != nil {
			h.logger.Errorf("[%v]Failed to connect to elasticsearch", credential.Username)
			return
		}
		solarRepo := repo.NewSolarRepo(elastic)

		db, err := infra.NewGormDB()
		if err != nil {
			h.logger.Errorf("[%v]Failed to connect to gorm", credential.Username)
			return
		}
		siteRegionRepo := repo.NewSiteRegionMappingRepo(db)

		serv := service.NewGrowattCollectorService(solarRepo, siteRegionRepo, h.logger)
		if err != nil {
			h.logger.Errorf("[%v]Failed to create service", credential.Username)
			return
		}

		if err := serv.Run(credential); err != nil {
			h.logger.Errorf("[%v]Failed to run service: %v", credential.Username, err)
			return
		}

		h.logger.Infof("[%v] Finished in %v", credential.Username, time.Since(now).String())

	}
}
