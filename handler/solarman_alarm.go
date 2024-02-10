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

type SolarmanAlarmHandler struct {
	logger logger.Logger
}

func NewSolarmanAlarmHandler() *SolarmanAlarmHandler {
	return &SolarmanAlarmHandler{}
}

func (h *SolarmanAlarmHandler) Run() {
	h.logger = logger.NewLogger(
		&logger.LoggerOption{
			LogName:     util.GetLogName(constant.SOLARMAN_ALARM_LOG_NAME),
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

	credentialRepo := repo.NewSolarmanCredentialRepo(db)
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

func (h *SolarmanAlarmHandler) run(credential *model.SolarmanCredential) func() {
	return func() {
		now := time.Now()
		snmp, err := infra.NewSnmp()
		if err != nil {
			h.logger.Errorf("[%v]Failed to connect to snmp", credential.Username)
			return
		}

		snmpRepo := repo.NewSnmpRepo(snmp)
		defer snmpRepo.Close()

		rdb, err := infra.NewRedis()
		if err != nil {
			h.logger.Errorf("[%v]Failed to connect to redis", credential.Username)
			return
		}
		defer rdb.Close()

		elastic, err := infra.NewElasticsearch()
		if err != nil {
			h.logger.Errorf("[%v]Failed to connect to elasticsearch", credential.Username)
			return
		}
		solarRepo := repo.NewSolarRepo(elastic)

		serv := service.NewSolarmanAlarmService(solarRepo, snmpRepo, rdb, h.logger)
		if err := serv.Run(credential); err != nil {
			h.logger.Errorf("[%v]Failed to run service: %v", credential.Username, err)
			return
		}
		h.logger.Infof("[%v] Finished in %v", credential.Username, time.Since(now).String())
	}
}

func (h *SolarmanAlarmHandler) Mock() {
	h.logger = logger.NewLogger(
		&logger.LoggerOption{
			LogName:     util.GetLogName(constant.SOLARMAN_ALARM_LOG_NAME),
			LogSize:     1024,
			LogAge:      90,
			LogBackup:   1,
			LogCompress: false,
			LogLevel:    logger.LOG_LEVEL_DEBUG,
			SkipCaller:  1,
		},
	)
	defer h.logger.Close()

	credentialRepo := repo.NewMockSolarmanCredentialRepo()
	credentials, err := credentialRepo.FindAll()
	if err != nil {
		h.logger.Error(err)
		return
	}

	pool := workerpool.New(1)
	for _, credential := range credentials {
		clone := credential
		pool.Submit(h.mock(&clone))
	}
	pool.StopWait()
}

func (h *SolarmanAlarmHandler) mock(credential *model.SolarmanCredential) func() {
	return func() {
		snmpRepo := repo.NewMockSnmpRepo()
		defer snmpRepo.Close()

		rdb, err := infra.NewRedis()
		if err != nil {
			h.logger.Error(err)
			return
		}
		defer rdb.Close()

		elastic, err := infra.NewElasticsearch()
		if err != nil {
			h.logger.Errorf("[%v]Failed to connect to elasticsearch", credential.Username)
			return
		}
		solarRepo := repo.NewSolarRepo(elastic)

		serv := service.NewSolarmanAlarmService(solarRepo, snmpRepo, rdb, h.logger)
		if err := serv.Run(credential); err != nil {
			h.logger.Error(err)
		}
	}
}
