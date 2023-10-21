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

type KStarCollectorHandler struct {
	logger logger.Logger
}

func NewKStarCollectorHandler() *KStarCollectorHandler {
	return &KStarCollectorHandler{}
}

func (h *KStarCollectorHandler) Run() {
	h.logger = logger.NewLogger(
		&logger.LoggerOption{
			LogName:     util.GetLogName(constant.KSTAR_COLLECTOR_LOG_NAME),
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
		h.logger.Errorf("Failed to connect to database: %v", err)
		return
	}

	credentialRepo := repo.NewKStarCredentialRepo(db)
	credentials, err := credentialRepo.GetCredentials()
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

func (h *KStarCollectorHandler) run(credential *model.KStarCredential) func() {
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

		serv, err := service.NewKStarCollectorService(solarRepo, siteRegionRepo, h.logger)
		if err != nil {
			h.logger.Errorf("[%v]Failed to create service", credential.Username)
			return
		}

		if err := serv.Run(credential); err != nil {
			h.logger.Error(err)
		}
		h.logger.Infof("[%v] Finished in %v", credential.Username, time.Since(now))
	}
}

func (h *KStarCollectorHandler) Mock() {
	h.logger = logger.NewLogger(
		&logger.LoggerOption{
			LogName:     util.GetLogName(constant.KSTAR_COLLECTOR_LOG_NAME),
			LogSize:     1024,
			LogAge:      90,
			LogBackup:   1,
			LogCompress: false,
			LogLevel:    logger.LOG_LEVEL_DEBUG,
			SkipCaller:  1,
		},
	)
	defer h.logger.Close()
	credentialRepo := repo.NewMockKStarCredentialRepo()
	credentials, err := credentialRepo.GetCredentials()
	if err != nil {
		h.logger.Error(err)
		return
	}

	pool := workerpool.New(len(credentials))
	for _, credential := range credentials {
		clone := credential
		pool.Submit(h.mock(&clone))
	}
	pool.StopWait()
}

func (h *KStarCollectorHandler) mock(credential *model.KStarCredential) func() {
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

		serv, err := service.NewKStarCollectorService(solarRepo, siteRegionRepo, h.logger)
		if err != nil {
			h.logger.Errorf("[%v]Failed to create service", credential.Username)
			return
		}

		if err := serv.Run(credential); err != nil {
			h.logger.Error(err)
		}
		h.logger.Infof("[%v] Finished in %v", credential.Username, time.Since(now))
	}
}
