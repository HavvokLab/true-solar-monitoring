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

type KStarAlarmHandler struct {
	logger logger.Logger
}

func NewKStarAlarmHandler() *KStarAlarmHandler {
	return &KStarAlarmHandler{}
}

func (h *KStarAlarmHandler) Run() {
	h.logger = logger.NewLogger(
		&logger.LoggerOption{
			LogName:     util.GetLogName(constant.KSTAR_ALARM_LOG_NAME),
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

	credentialRepo := repo.NewKStarCredentialRepo(db)
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

func (h *KStarAlarmHandler) run(credential *model.KStarCredential) func() {
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

		serv := service.NewKStarAlarmService(snmpRepo, rdb, h.logger)
		if err := serv.Run(credential); err != nil {
			h.logger.Errorf("[%v]Failed to run service: %v", credential.Username, err)
			return
		}
		h.logger.Infof("[%v] Finished in %v", credential.Username, time.Since(now).String())
	}
}

func (h *KStarAlarmHandler) Mock() {
	h.logger = logger.NewLogger(
		&logger.LoggerOption{
			LogName:     util.GetLogName(constant.KSTAR_ALARM_LOG_NAME),
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

	credentialRepo := repo.NewKStarCredentialRepo(db)
	credentials, err := credentialRepo.FindAll()
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

func (h *KStarAlarmHandler) mock(credential *model.KStarCredential) func() {
	return func() {
		now := time.Now()
		snmpRepo := repo.NewMockSnmpRepo()
		defer snmpRepo.Close()

		rdb, err := infra.NewRedis()
		if err != nil {
			h.logger.Errorf("[%v]Failed to connect to redis", credential.Username)
			return
		}
		defer rdb.Close()

		serv := service.NewKStarAlarmService(snmpRepo, rdb, h.logger)
		if err := serv.Run(credential); err != nil {
			h.logger.Errorf("[%v]Failed to run service: %v", credential.Username, err)
			return
		}
		h.logger.Infof("[%v] Finished in %v", credential.Username, time.Since(now).String())
	}
}
