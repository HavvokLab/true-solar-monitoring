package handler

import (
	"time"

	"github.com/HavvokLab/true-solar-monitoring/constant"
	"github.com/HavvokLab/true-solar-monitoring/infra"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/model"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/HavvokLab/true-solar-monitoring/service"
	"github.com/gammazero/workerpool"
)

type HuaweiCollectorHandler struct {
	logger logger.Logger
}

func NewHuaweiCollectorHandler() *HuaweiCollectorHandler {
	logger := logger.NewLogger(
		&logger.LoggerOption{
			LogName:     constant.GetLogName(constant.SOLARMAN_COLLECTOR_LOG_NAME),
			LogSize:     1024,
			LogAge:      90,
			LogBackup:   1,
			LogCompress: false,
			LogLevel:    logger.LOG_LEVEL_DEBUG,
			SkipCaller:  1,
		},
	)
	return &HuaweiCollectorHandler{logger: logger}
}

func (h *HuaweiCollectorHandler) Run() {
	h.logger = logger.NewLogger(
		&logger.LoggerOption{
			LogName:     constant.GetLogName(constant.HUAWEI_COLLECTOR_LOG_NAME),
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

	credentialRepo := repo.NewHuaweiCredentialRepo(db)
	credentials, err := credentialRepo.GetCredentialsByOwner(constant.TRUE_OWNER)
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

func (h *HuaweiCollectorHandler) run(credential *model.HuaweiCredential) func() {
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

		serv := service.NewHuaweiCollectorService(solarRepo, siteRegionRepo, h.logger)
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
